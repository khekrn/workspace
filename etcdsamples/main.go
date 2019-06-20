package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// 1. Install go client package via - go get -v go.etcd.io/etcd/clientv3
// 2. Download respective etcd package from github for windows or linux
func main() {
	cl, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})

	defer cl.Close()

	if err != nil {
		fmt.Println("connection failed = ", err)
		return
	}

	fmt.Println("connection success")

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Inserting 1000 elements....")
		for i := 0; i < 1000; i++ {
			index := strconv.Itoa(i + 1)
			key := "movie " + index
			value := "John Wick " + index
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			cl.Put(ctx, key, value)

			cancel()

			if err != nil {
				fmt.Println("put failed, err:", err)
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Get and removing elements....")
		for i := 0; i < 1000; i++ {
			key := "movie " + strconv.Itoa(i+1)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			resp, err := cl.Get(ctx, key)
			cancel()

			if err != nil {
				fmt.Println("get failed, err:", err)
				break
			}

			for _, ev := range resp.Kvs {
				fmt.Println(string(ev.Key), " - ", string(ev.Value))
			}
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("removing elements....")
		for i := 0; i < 1000; i++ {
			key := "movie " + strconv.Itoa(i+1)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			_, err := cl.Delete(ctx, key)
			cancel()

			if err != nil {
				fmt.Println("get failed, err:", err)
				break
			}

			fmt.Println("Deleted key = ", key)

			time.Sleep(1 * time.Second)

		}
	}()

	wg.Wait()

}

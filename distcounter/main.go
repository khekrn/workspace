package main

import (
	"distcounter/counter"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/memberlist"
	uuid "github.com/satori/go.uuid"
)

var (
	counterObj *counter.Counter
	m          *memberlist.Memberlist
	members    = flag.String("members", "", "comma seperated members")
	port       = flag.Int("port", 4041, "http port")
	rpcPort    = flag.Int("rpcport", 0, "memberlist port (0 = auto select)")
)

func start() error {
	fmt.Println("Hello Go")
	flag.Parse()

	counterObj = &counter.Counter{}
	hostname, _ := os.Hostname()
	fmt.Println(hostname)

	cfg := memberlist.DefaultLocalConfig()
	cfg.BindPort = *rpcPort
	uid, _ := uuid.NewV4()
	cfg.Name = hostname + "-" + uid.String()
	cfg.PushPullInterval = time.Second * 5
	cfg.ProbeInterval = time.Second * 1

	var err error
	m, err = memberlist.Create(cfg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if len(*members) > 0 {
		memList := strings.Split(*members, ",")
		_, err := m.Join(memList)
		if err != nil {
			log.Fatal("member list join error ", err)
			return err
		}
	}
	node := m.LocalNode()
	fmt.Printf("Local member %s:%d\n", node.Addr, node.Port)
	return nil
}

func defaultHandler() func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, counterObj)
	}
}

func incrementHandler() func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		amount := req.FormValue("increment")
		amountInt, err := strconv.Atoi(amount)
		if err != nil {
			log.Fatal("parse int error for string ", amount)
			return
		}

		if amountInt < 0 {
			log.Fatal("decrement is not yet supported....")
			return
		}

		counterObj.Increment(int32(amountInt))
		fmt.Printf("Incremented counter to %v\n", counterObj)
		fmt.Fprintln(resp, counterObj)
	}
}

func clusterHandler() func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		json.NewEncoder(resp).Encode(m.Members())
	}
}

func main() {
	if err := start(); err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", defaultHandler())
	http.HandleFunc("/increment", incrementHandler())
	http.HandleFunc("/cluster", clusterHandler())

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatal("cannot able to start server")
		return
	}

}

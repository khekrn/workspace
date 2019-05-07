package read

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

// AllContentsToMemory reads file and loads all content to memory
func AllContentsToMemory(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, fileInfo.Size())

	readBytes, err := file.Read(buffer)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bytes Read = ", readBytes)
	fmt.Println(string(buffer))
}

// AllContentsToMemoryV2 reads file and loads all content to memory
func AllContentsToMemoryV2(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func ReadInChunks(filePath string) {
	defaultBufferSize := 512
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, defaultBufferSize)

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println("Bytes Read = ", bytesRead)
		fmt.Println("Content = ", string(buffer[:defaultBufferSize]))
	}

}

func ReadChunksInParallel(filePath string) {

	type chunk struct {
		buffSize int
		offset   int64
	}

	bufferSize := 2048
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println(err)
		return
	}

	fileSize := int(fileInfo.Size())
	concurrency := int(fileSize / bufferSize)

	chunks := make([]chunk, concurrency)

	for i := 0; i < concurrency; i++ {
		chunks[i].buffSize = bufferSize
		chunks[i].offset = int64(bufferSize * i)
	}

	if remainder := fileSize % bufferSize; remainder != 0 {
		extraChunk := chunk{buffSize: bufferSize, offset: int64(concurrency * bufferSize)}
		concurrency++
		chunks = append(chunks, extraChunk)
	}

	var wg sync.WaitGroup

	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(ch []chunk, index int) {
			defer wg.Done()

			chunk := ch[index]
			buffer := make([]byte, chunk.buffSize)

			_, err := file.ReadAt(buffer, chunk.offset)
			if err != nil {
				fmt.Println("Error in index = ", index, " - ", err.Error())
				return
			}

			fmt.Println("Content = ", string(buffer))
		}(chunks, i)
	}

	wg.Wait()
}

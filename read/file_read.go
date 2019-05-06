package read

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
	defaultBufferSize := 100
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

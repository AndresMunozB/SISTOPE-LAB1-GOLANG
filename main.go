package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func read(namefile string, wg *sync.WaitGroup, position int, lines, lenline int) {
	defer wg.Done()

	file, err := os.Open(namefile)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)                           // creates a new reader
	_, err = reader.Discard(lines * position * (lenline + 1)) // discard the following 64 bytes
	check(err)
	// use isPrefix if is needed, this example doesn't use it
	// read line until a new line is found
	for i := 0; i < lines; i++ {
		line, _, err := reader.ReadLine()
		check(err)
		fmt.Printf("GoRutine: %d: %s\n", position, string(line))
	}

}

func main() {
	len := 60
	totalLines := 10
	gorutines := 5
	lines := totalLines / gorutines
	var wg sync.WaitGroup
	wg.Add(gorutines)
	for i := 0; i < gorutines; i++ {
		go read("test.txt", &wg, i, lines, len)
	}
	wg.Wait()
}

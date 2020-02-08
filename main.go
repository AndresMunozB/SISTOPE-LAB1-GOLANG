package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func function(str []byte) {

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
		function(line)
		//fmt.Printf("GoRutine: %d: %s\n", position, string(line))
	}

}

func main() {
	fmt.Printf("CPUs: %d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	len := 60
	totalLines := 50000000
	gorutines := 20
	lines := totalLines / gorutines
	var wg sync.WaitGroup
	wg.Add(gorutines)
	for i := 0; i < gorutines; i++ {
		go read("test2.txt", &wg, i, lines, len)
	}
	wg.Wait()
	fmt.Printf("Success\n")
}

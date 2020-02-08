package main

import (
	"fmt"
	"io"
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

	o3, err := file.Seek(int64(lines*position*(lenline+1)), 0)
	check(err)
	b3 := make([]byte, lenline)
	n3, err := io.ReadAtLeast(file, b3, lenline)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
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

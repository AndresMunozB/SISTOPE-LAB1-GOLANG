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

func read(file *os.File, wg *sync.WaitGroup, line int64, len int64) {
	defer wg.Done()
	// El paquete `io` tiene funciones que pueden ser
	// utiles para leer archivos. Por ejemplo, una
	// lectura como la anterior puede ser implementada
	// más robustamente con `ReadAtLeast`
	o3, err := file.Seek(line*(len+1), 0)
	check(err)
	b3 := make([]byte, 60)
	n3, err := io.ReadAtLeast(file, b3, 60)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	/*scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())*/
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	file, err := os.Open("test.txt")
	check(err)
	defer file.Close()
	for i := 0; i < 10; i++ {
		go read(file, &wg, int64(i), 60)
	}
	wg.Wait()
}

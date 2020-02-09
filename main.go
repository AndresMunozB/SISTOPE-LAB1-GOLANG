package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func function(str []byte) {

}

func succession(serie string) bool {
	if strings.Contains(serie, "AAAA") {
		return true
	} else if strings.Contains(serie, "CCCC") {
		return true
	} else if strings.Contains(serie, "GGGG") {
		return true
	} else if strings.Contains(serie, "TTTT") {
		return true
	}
	return false
}
func read(namefile string, wg *sync.WaitGroup, position int, lines, lenline int) {
	defer wg.Done()

	file, err := os.Open(namefile)
	check(err)
	defer file.Close()

	fileRP, err := os.Create("rp_" + strconv.Itoa(position) + ".txt")
	check(err)
	defer fileRP.Close()

	reader := bufio.NewReader(file)                           // creates a new reader
	_, err = reader.Discard(lines * position * (lenline + 1)) // discard the following 64 bytes
	check(err)
	// use isPrefix if is needed, this example doesn't use it
	// read line until a new line is found
	for i := 0; i < lines; i++ {
		line, _, err := reader.ReadLine()
		check(err)
		lineStr := string(line)
		if succession(lineStr) {
			lineStr += " SI\n"
		} else {
			lineStr += " NO\n"
		}
		_, err = fileRP.WriteString(lineStr)
		check(err)
		function(line)
		//fmt.Printf("GoRutine: %d: %s\n", position, string(line))
	}

}

func join(fileName string, gorutines int) {
	outputFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer outputFile.Close()
	for i := 0; i < gorutines; i++ {
		rpStr := "rp_" + strconv.Itoa(i) + ".txt"
		rpFile, err := os.Open(rpStr)
		check(err)

		n, err := io.Copy(outputFile, rpFile)
		check(err)
		fmt.Printf("wrote %d bytes of %s to %s\n", n, rpStr, fileName)
		rpFile.Close()

	}
}

func main() {
	defer duration(track("Execution Time:"))
	//SETTING THE THREADS
	//fmt.Printf("CPUs: %d\n", runtime.NumCPU())
	//runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("CPUs: %d\n", 1)
	runtime.GOMAXPROCS(1)

	//DECLARATING PARAMETERS
	len := 60
	totalLines := 50000000
	gorutines := 5
	lines := totalLines / gorutines
	nameInputFile := "test_50000000.txt"
	nameOutputFile := "output.txt"

	//RUNING GORUTINES
	var wg sync.WaitGroup
	wg.Add(gorutines)
	for i := 0; i < gorutines; i++ {
		go read(nameInputFile, &wg, i, lines, len)
	}

	//JOINING THE RESULTS
	wg.Wait()
	join(nameOutputFile, gorutines)
	fmt.Printf("Success\n")

}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

func generateLine(len int, r *rand.Rand) string {

	line := ""
	for i := 0; i < len; i++ {
		random := r.Intn(4)
		line += strconv.Itoa(random)
		//fmt.Printf("%d", random)
	}
	//fmt.Printf("\n")
	return line

}
func transformLine(str string) string {
	result := strings.ReplaceAll(str, "0", "A")
	result = strings.ReplaceAll(result, "1", "C")
	result = strings.ReplaceAll(result, "2", "G")
	result = strings.ReplaceAll(result, "3", "T")
	return result
}
func main() {
	defer duration(track("Execution Time:"))
	seed := int64(8)
	lines := 50000000
	f, err := os.Create("test_" + strconv.Itoa(lines) + ".txt")
	check(err)

	r := rand.New(rand.NewSource(seed))
	for i := 0; i < lines; i++ {
		str := transformLine(generateLine(60, r))
		//fmt.Printf("%s\n", str)
		fmt.Fprintln(f, str)
	}
	err = f.Close()
	fmt.Println("File written successfully.")

}

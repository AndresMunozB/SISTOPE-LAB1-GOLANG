package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

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
	seed := int64(4)
	lines := 50000000
	f, err := os.Create("test2.txt")
	check(err)

	r := rand.New(rand.NewSource(seed))
	for i := 0; i < lines; i++ {
		str := transformLine(generateLine(60, r))
		//fmt.Printf("%s\n", str)
		fmt.Fprintln(f, str)
	}
	err = f.Close()
	check(err)
	fmt.Println("File written successfully")

}

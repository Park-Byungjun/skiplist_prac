package main

import (
	"bufio"
	"fmt"
	"github.com/MauriceGit/skiplist"
	"io"
	"math/bits"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type IntElement int

func (e IntElement) ExtractKey() float64 {
	return float64(e)
}
func (e IntElement) String() string {
	return fmt.Sprintf("%03d", e)
}

type Element struct {
	key   float64
	value string
}

// Implement the interface used in skipList

func (e Element) ExtractKey() float64 {
	return e.key
}
func (e Element) String() string {
	return fmt.Sprintf("%s", e.value)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var x = rand.Uint64() & ((1 << uint(24)) - 1)
	zeroes := bits.TrailingZeros64(x)
	fmt.Println(x, zeroes)
	for i := 0; i < 5; i++ {
		myList := skiplist.New()
		fmt.Println(myList)

		reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
		listLen, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 32)
		if err != nil {
			fmt.Println("[parseInt] integer transform failed")
		}

		fmt.Println(myList.String())
		for i := 1; i < int(listLen)+1; i++ {
			//myList.Insert(Element{float64(i*100-1), "ELEM"+ strconv.Itoa(i*100-1)})
			//myList.Insert(Element{float64(i*100+1), "ELEM"+ strconv.Itoa(i*100+1)})
			myList.Insert(IntElement(i))
			//myList.Insert(Element{float64(i*100-1), "elem"+ strconv.Itoa(i*100-1)})
			//myList.Insert(Element{float64(i*100+1), "elem"+ strconv.Itoa(i*100+1)})
			fmt.Println(myList.String())
		}
		fmt.Println(myList.GetNodeCount())
		fmt.Println(myList.GetSmallestNode())
		fmt.Println(myList.GetLargestNode())
		fmt.Println(myList)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

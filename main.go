package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func timeit(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("# %s duration: %+v\n", name, elapsed)
}

func part1(scanner *bufio.Scanner) (result int) {
	defer timeit(time.Now(), "part1")
	return
}

func part2(scanner *bufio.Scanner) (result int) {
	defer timeit(time.Now(), "part2")
	return
}

func main() {
	defer timeit(time.Now(), "main")
	filename := "input"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	fd, err := os.Open(filename)
	defer fd.Close()
	check(err)

	scanner := bufio.NewScanner(fd)

	result1 := part1(scanner)
	fmt.Printf("part1 result: %+v\n", result1)

	fd.Seek(0, io.SeekStart)

	result2 := part2(scanner)
	fmt.Printf("part2 result: %+v\n", result2)
}

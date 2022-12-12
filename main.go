package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func logTime() {
	start := time.Now()
	duration := time.Since(start)
	res := duration.Nanoseconds()

	fmt.Printf("Время выполнения программы: %d наносекунд", res)
}

func file(filePath string) {
	f1, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Unable to open file [%s]", err.Error())
		os.Exit(1)
	}
	defer f1.Close()

	createFile, err := os.Create("out.txt")
	if err != nil {
		fmt.Printf("Unable to create file [%s]", err.Error())
		os.Exit(1)
	}
	defer createFile.Close()

	s := bufio.NewScanner(f1)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		writeFile, err := createFile.WriteString(s.Text())
		if err != nil {
			fmt.Printf("Unable to write file [%s]", err.Error())
		}
		fmt.Println("Сколько байт записано: ", writeFile)
	}

}

func main() {
	path := "in.txt"
	file(path)
	defer logTime()
}

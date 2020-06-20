package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

type DataLine struct {
	content string
	indentity int
}

func bai3() {
	buffReadData := make(chan *DataLine, 10)
	defer close(buffReadData)
	var wg sync.WaitGroup

	f, _ := os.Open("file.txt")
	defer f.Close()

	for i := 1 ; i <= 3 ; i++ {
		go printData(buffReadData, &wg)
	}

	scanner := bufio.NewScanner(f)

	count := 1

	for scanner.Scan() {
		dataLine := &DataLine {content: scanner.Text(), indentity: count}
		count ++
		buffReadData <- dataLine 
		wg.Add(1)
	}
	
	wg.Wait()
	
}

func printData(jobs chan *DataLine, wg *sync.WaitGroup) {
	
	for {
		select {
		case data := <- jobs :
			log.Printf("Hang %v : %v xong!\n", data.indentity, data.content)
			wg.Done()
		}
	}
}


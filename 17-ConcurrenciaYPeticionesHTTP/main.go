package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	// "sync"
	"time"
)

// var wg sync.WaitGroup

func hi(num int) {
	fmt.Println("Hola ", num)
	time.Sleep(1000 * time.Millisecond)
	// wg.Done()
}

func get(id int) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(id))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {
	// wg.Add(10)

	for i := 0; i < 100; i++ {
		// go hi(i)
		go get(i)
	}

	var str string
	fmt.Scan(&str)

	// wg.Wait()
}

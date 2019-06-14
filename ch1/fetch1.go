package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:]{
		go fetch(url, ch)
	}
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, cc chan string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		cc <- fmt.Sprint(err)
		return
	}
	bm, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil{
		cc <- fmt.Sprint(err)
	}
	tmCost := time.Since(start).Seconds()
	cc <- fmt.Sprintf("cost: %.2fs memory: %7dbytes url: %s", tmCost, bm, url)
}

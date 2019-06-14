package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}else{
		for _, arg := range files{
			fp, err := os.Open(arg)
			if err != nil{
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(fp, counts)
			_ = fp.Close()
		}
	}

	for line, lc := range counts{
		if lc >= 1{
			fmt.Printf("%d times \t %s\n", lc, line)
		}
	}
}


func countLines(fp *os.File, counts map[string]int){
	input := bufio.NewScanner(fp)
	for input.Scan(){
		counts[input.Text()] ++
	}
}
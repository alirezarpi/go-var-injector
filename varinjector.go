package main

import (
	"fmt"
    // "errors"
	// "os"
)

func printHelp(params map[string]string) {
	fmt.Printf("[ Var Injector ]\n")
	for k, v := range params {
		fmt.Printf("  %v\t%v\n", k, v)
	}
}

func main() {
	// args := os.Args[1:]

    params := map[string]string{
		"-e": "asda",
		"-s": "secret"}

	printHelp(params)
	
	for k := range params {
		fmt.Println(k)
		if (k == "-f")
	}
}

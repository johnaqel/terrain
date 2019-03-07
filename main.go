package main

import (
	"flag"
	"fmt"
)

func main() {
	resFlag := true
	flag.BoolVar(&resFlag, "r", resFlag, "List the Terraform resources declared")
	flag.Parse()

	fmt.Printf("resFlag: %v\n", resFlag)

}

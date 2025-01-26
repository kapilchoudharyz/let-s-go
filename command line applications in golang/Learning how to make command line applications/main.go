package main

import (
	"flag"
	"fmt"
)

func main() {
	//lengthOfOSArgs := len(os.Args)
	//fmt.Println(lengthOfOSArgs)
	//fmt.Printf("these are the arguments %v\n", os.Args)

	var (
		name string
		age  int
	)

	flag.StringVar(&name, "name", "Kapil", "you name")
	flag.IntVar(&age, "age", 22, "you age")
	flag.Parse()

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
}

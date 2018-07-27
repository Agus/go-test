package main

import "github.com/agus/go-test/subPackage"
import "fmt"

func main() {
	fmt.Println(subPackage.Double(5))
	fmt.Println(subPackage.Divide(6))
}

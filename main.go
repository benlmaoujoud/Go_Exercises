package main

import "fmt"

func vals() (string, string) {
    return "hello", "Go"
}

func main() {

    a, b := vals()
    fmt.Print(a," ",b)
    

   
	


}
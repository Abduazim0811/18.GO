// 1-masala
// package main

// import "fmt"

// func Fact(num int, ch chan int){
// 	son:=1
// 	for i:=1;i<=num;i++{
// 		son*=i
// 	}
// 	ch<-son
// }
// func main() {
// 	var num int
// 	ch:=make(chan int)
// 	fmt.Println("Son kiriting:")
// 	fmt.Scanln(&num)
// 	go Fact(num, ch)

// 	fmt.Println(<-ch)
// }
// 2-masala

package main

import (
	"fmt"
	"log"
	"os"
)

func ReadFileGo(file string, ch chan string) {
	str, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	ch <- string(str)
}

func main() {
	ch := make(chan string)
	go ReadFileGo("file1.txt", ch)
	go ReadFileGo("file2.txt", ch)
	go ReadFileGo("file3.txt", ch)
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}


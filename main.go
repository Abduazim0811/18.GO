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
// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// func main() {
//     var resault string
// 	files := []string{"file1.txt", "file2.txt", "file3.txt"}

// 	for _, file := range files {
//         str,err:=ReadFile(file)
//         if err!=nil{
//             log.Fatal(err)
//             continue
//         }
//         resault+=str
// 	}
//     fmt.Println("Birlashtirilgan content: ", resault)
// }

// func ReadFile(str string) (string, error) {
// 	file, err := os.Open(str)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	res, err := io.ReadAll(file)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(res), nil

// }

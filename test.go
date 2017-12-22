package main
import (
    "fmt"
    "strings"
)

func main() { book := "apple is better than kiwi"
 // Replace the "apple" with "banana" 
	result := strings.Replace(book, "apple", "banana", -1) 
	fmt.Println(result) 
}

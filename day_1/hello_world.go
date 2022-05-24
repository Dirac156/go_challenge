package main
import (
    "fmt"
    "os"
    "bufio"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 scanner := bufio.NewScanner(os.Stdin)
 scanner.Scan()
 s := scanner.Text()
 fmt.Println("Hello, World.")
 fmt.Println(s)
}
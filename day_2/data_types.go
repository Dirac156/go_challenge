package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    var i2 uint64
    
    var d2 float64
    
    var s2 string
    
    // Read and save an integer, double, and String to your variables.
    
    fmt.Scanf("%d", &i2)
    fmt.Scanf("%f", &d2)
    scanner.Scan()
    s2 = scanner.Text()
    
    // Print the sum of both integer variables on a new line.
    
    sumInt := i + i2
    
    fmt.Println(sumInt)
    
    // Print the sum of the double variables on a new line.
    
    sumDouble := d + d2
    
    fmt.Printf("%.1f\n", sumDouble)
    
    // Concatenate and print the String variables on a new line
    
    fmt.Println(s + "" +  s2)
    
    // The 's' variable above should be printed first.

}
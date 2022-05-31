package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 
    var num int
    fmt.Scan(&num)
    scanner := bufio.NewReader(os.Stdin)
    var lines []string
    var i int
    for i <= num {
        line, _ := scanner.ReadString('\n')
        lines = append(lines, line)
        i++
    }
    
    
//  loop over the input the text.

for _, v := range(lines) {
    if len(v) > 1 {
        var (
            strEven, strOdd string
        )
        
        for i := range(v) {
            if v[i] != '\r' && v[i] != '\n' {
                if (i % 2 == 0) {
                    strEven = strEven + string(v[i])
                } else {
                    strOdd = strOdd + string(v[i])
                }
            }
        }
        
        fmt.Printf("%s %s\n", strEven, strOdd)
    }
}


}
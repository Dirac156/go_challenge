package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    scanner := bufio.NewReader(os.Stdin)
    
    var numberEntries int 
    
    fmt.Scan(&numberEntries)
    
    m := make(map[string]string)
    
    for i := 0; i < numberEntries; i++ {
        // returns the user input and the error
        line, _ := scanner.ReadString('\n')
        // split the text where there is a white space
        words := strings.Fields(line)
        
        // add values inside the map
        m[words[0]] = words[1]
    }
    
    for true {
        line, err := scanner.ReadString('\n')
        
        if err != nil && line == "" {
            break
        }
        
		// remove line break
        line = strings.TrimSuffix(line, "\n")
        
        val, exist := m[line]

        
        if exist == true  {
            fmt.Printf("%s=%s\n", line, val)
        } else {
            fmt.Println("Not found")
        }
        
        if err != nil {
           break 
        }
    };
}
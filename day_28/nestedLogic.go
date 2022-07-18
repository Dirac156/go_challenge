package main
import (
    "fmt"
)

type date struct {
    day int
    month int
    year int
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    
    var D, M, Y int
    
    fmt.Scan(&D)
    fmt.Scan(&M)
    fmt.Scan(&Y)
    returnedDate := date{ day : D, month: M, year: Y }
    
    fmt.Scan(&D)
    fmt.Scan(&M)
    fmt.Scan(&Y)
    dueDate := date{ day : D, month: M, year: Y }
    
    // case covered:
    /*
        1. The book is returned in another calendar year.
        2. The book is returned the same year but another month
        3. The book is returned the same year and month but another calendar month.
        4. The book is returned before the dueDate or on time.
        
    */
    
    if (returnedDate.year > dueDate.year) {
        // different calendar year
        fmt.Println(10000)
    } else if ((returnedDate.month > dueDate.month) && (returnedDate.year == dueDate.year)) {
        // returned same year but different months
        fmt.Println( 500 * (returnedDate.month - dueDate.month))
    } else if ((returnedDate.day > dueDate.day) && (returnedDate.year == dueDate.year) && (returnedDate.month == dueDate.month)) {
        // returned same year and same month but after the deadline
        fmt.Println( 15 * (returnedDate.day - dueDate.day))   
    }else {
        // returned on time
        fmt.Println(0)
    }
 
 
 
}

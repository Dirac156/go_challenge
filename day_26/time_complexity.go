package main
import "fmt"


func isPrimeNumber(number int) bool {
    if number < 2 {
        return false
    }
    for i := 2; i*i <= number; i++ {
        if (number % i) == 0 {
            return false;
        }
    }
    return true;
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 // get number of inputs
 var nbr int
 
 _, err := fmt.Scanf("%d", &nbr);
 
  if err != nil {
     fmt.Println(err)
 }
 
 for i := 0; i < nbr; i++ {
     var j int
    _, err := fmt.Scanf("%d", &j);
 
    if err != nil {
        fmt.Println(err)
    }
    if (isPrimeNumber(j)) {
        fmt.Println("Prime")
    } else {
        fmt.Println("Not prime")
    }
 }


}

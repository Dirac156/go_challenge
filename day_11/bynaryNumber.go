package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)


func IntegerToBinary(n int32) string {
   return strconv.FormatInt(int64(n), 2)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)
    
    currentNumber := 0
    maxNumber := 0
    
    numberBinary := IntegerToBinary(n)
    
    for _, v := range(numberBinary) {
        byteToInt, _ := strconv.Atoi(string(v))
        
        if byteToInt == 0 {
            if maxNumber < currentNumber {
                maxNumber = currentNumber
            }
            
            currentNumber = 0
        } else {
            currentNumber++
        }
    }
    
    if maxNumber < currentNumber {
        maxNumber = currentNumber
    }
    
    fmt.Println(maxNumber)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != 6 {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }
    
    var maxHourGlass int32
    
    // loop over the each 4 rows of the arr
    for i := 0; i < 4; i++ {
        prime := i;
        second := i + 1;
        last := i + 2;
        
        for j := 0; j < 4; j++ {
            primeNumbe := j
            secondNumb := j + 1
            lastNumb := j + 2
            
            var line1, line2, line3 int32
            
            line1 = arr[prime][primeNumbe] + arr[prime][secondNumb] + arr[prime][lastNumb]
            
            line2 = arr[second][secondNumb]
            
            line3 = arr[last][primeNumbe] + arr[last][secondNumb] + arr[last][lastNumb]
            
            var sum int32
            
            sum = line1 + line2 + line3
            
            if i == 0 && j == 0 {
                maxHourGlass = sum
            }
            
            if ( sum > maxHourGlass) {
                maxHourGlass = sum
            }
            
        }
    }
    
    fmt.Printf("%d\n", maxHourGlass)
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

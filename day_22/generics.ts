'use strict';

process.stdin.resume();
process.stdin.setEncoding('utf-8');
let inputString: string = '';
let inputLines: string[] = [];
let currentLine: number = 0;
process.stdin.on('data', function(inputStdin: string): void {
    inputString += inputStdin;
});

process.stdin.on('end', function(): void {
    inputLines = inputString.split('\n');
    inputString = '';
    main();
});

function readLine(): string {
    return inputLines[currentLine++];
}

function printArray(a: Array<any>){
    let aList: number = Number(a[0])
    let len: number = Number(a[0]) + 1
    let i = 1
    while (true) {
        for (let j = 0; j < aList; i++, j++) {
            console.log(a[i])
        }
        
        if  (len < a.length) {
            len += Number(a[i]) + 1
            aList = Number(a[i])
            i++
        } else {
            break
        }
    }
}

function main() {
    // Enter your code here
    printArray(inputLines)
}

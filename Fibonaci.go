package main

import "fmt"

func isFibonacciNumber(A int) int {  
    a, b := 1,0         // Согласно формуле 1 и 2 число начинаются с 1 и 0 соответственно
    n := 0

    for b < A {
        a, b = b, a+b
        n++
    }

    if b == A {
        return n
    } else {
        return -1    // если число не является числом фибоначи, то ответ -1
    }
}

func main() {
    var A int
    fmt.Scan(&A)              // Считываем число
    
    result := isFibonacciNumber(A)
    
    fmt.Println(result)
}
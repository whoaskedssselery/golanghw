package main

import (
    "fmt"
)

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for _, str := range strs[1:] {
        for len(str) < len(prefix) || str[:len(prefix)] != prefix {
            prefix = prefix[:len(prefix)-1]
            if prefix == "" {
                return ""
            }
        }
    }

    return prefix
}

func main() {
    var strs []string
    var input string
    var n int

    fmt.Print("Введите количество строк: ")
    fmt.Scan(&n)

    fmt.Println("Введите строки:")
    for i := 0; i < n; i++ {
        fmt.Scan(&input)
        strs = append(strs, input)
    }

    result := longestCommonPrefix(strs)
    if result == "" {
        fmt.Println("Нет общего префикса")
    } else {
        fmt.Println("Longest common prefix:", result)
    }
}

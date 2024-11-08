package main
import "fmt"
var a int64
func main() {
    fmt.Scan(&a)
    if((a%4==0 && a%100!=0)||a%400==0){
        fmt.Print("Год високосный")
    }else{
        fmt.Print("Год не високосный")
    }
}

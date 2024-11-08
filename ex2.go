package main

import (
    "fmt"
    "strconv"
    "strings"
)
func calculateSimple(operand1, operand2 float64, operator string) float64 {
    switch operator {
    case "+":
        return operand1 + operand2
    case "-":
        return operand1 - operand2
    case "*":
        return operand1 * operand2
    case "/":
        return operand1 / operand2
    default:
        return 0
    }
}
func calculate(expression string) float64 {
    expression = strings.ReplaceAll(expression, " ", "")

    var numbers []float64
    var operators []string
    var i int
	
    for i < len(expression) {
        char := string(expression[i])
        if char == "(" {
            balance := 1
            j := i + 1
            for j < len(expression) && balance > 0 {
                if expression[j] == '(' {
                    balance++
                } else if expression[j] == ')' {
                    balance--
                }
                j++
            }
            num := calculate(expression[i+1 : j-1])
            numbers = append(numbers, num)
            i = j
            continue
        }
        if char >= "0" && char <= "9" || char == '.' {
            var numStr string
            for i < len(expression) && (expression[i] >= '0' && expression[i] <= '9' || expression[i] == '.') {
                numStr += string(expression[i])
                i++
            }
            num, _ := strconv.ParseFloat(numStr, 64)
            numbers = append(numbers, num)
            continue
        }
        if char == "+" || char == "-" || char == "*" || char == "/" {
            for len(operators) > 0 && precedence(char) <= precedence(operators[len(operators)-1]) {
                applyOperation(&numbers, &operators)
            }
            operators = append(operators, char)
        }
        i++
    }
    for len(operators) > 0 {
        applyOperation(&numbers, &operators)
    }
    return numbers[0]
}
func precedence(op string) int {
    switch op {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    default:
        return 0
    }
}


func applyOperation(numbers *[]float64, operators *[]string) {
    if len(*numbers) < 2 || len(*operators) == 0 {
        return
    }

    operand2 := (*numbers)[len(*numbers)-1]
    operand1 := (*numbers)[len(*numbers)-2]
    operator := (*operators)[len(*operators)-1]

    *numbers = (*numbers)[:len(*numbers)-2]
    *operators = (*operators)[:len(*operators)-1]

    result := calculateSimple(operand1, operand2, operator)
    *numbers = append(*numbers, result)
}

var expression string
func main() {
    fmt.Scan(&expression)
    result := calculate(expression)
    fmt.Println("Result:", result)
}

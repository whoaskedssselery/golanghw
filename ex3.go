import (
	"fmt"
   )
   func generatePascalsTriangle(n int) [][]int {
	triangle := make([][]int, n)
	for i := range triangle {
	 triangle[i] = make([]int, i+1)
	 triangle[i][0], triangle[i][i] = 1, 1
	 for j := 1; j < i; j++ {
	  triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
	 }
	}
	return triangle
   }

   func printPascalsTriangle(triangle [][]int) {
	for _, row := range triangle {
	 for _, val := range row {
	  fmt.Printf("%d ", val)
	 }
	 fmt.Println()
	}
   }
   func main() {
	var n int
	fmt.Print("Введите размер треугольника Паскаля: ","\n")
	fmt.Scan(&n)

	if n <= 0 {
	 fmt.Println("Размер должен быть положительным числом.")
	 return
	}
	triangle := generatePascalsTriangle(n)
	printPascalsTriangle(triangle)
   }


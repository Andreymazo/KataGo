
// package main

// import (
//     "fmt"
//     "strings"
// )

// func main() {
//     s := "This,is,a,delimited,string"
//     v := strings.Split(s, ",")
//     fmt.Println(v)     // [This is a delimited string]
// }

// package main

// import (
// 	"fmt"
// )

// const N = 1e6

// var f32 = make([]float32, 1024)
// var slice64 []float64

//	func FuncVar(f32 []float32) []float64 {
//		f64 := make([]float64, len(f32))
//		var f float32
//		var i int
//		for i, f = range f32 {
//			f64[i] = float64(f)
//		}
//		return f64
//	}
// func main() {
// 	// BenchmarkFuncVar()
// 	// BenchmarkRangeVar()
// 	slice32 := make([]float32, 10)
// 	slice64 := convertTo64(slice32)
// 	fmt.Println(slice64)
// }

// func convertTo64(ar []float32) []float64 {
// 	newar := make([]float64, len(ar))
// 	var v float32
// 	var i int
// 	for i, v = range ar {
// 		newar[i] = float64(v)
// 	}
// 	return newar
// }

// https://golangify.com/string-to-array

// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// const refString = "Mary*had,a%little_lamb"

// func main() {

// 	words := regexp.MustCompile("[*,%_]{1}").Split(refString, -1)
// 	for idx, word := range words {
// 		fmt.Printf("Word %d is: %s\n", idx, word)
// 	}

// }

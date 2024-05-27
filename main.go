package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const vadacha_panici string = "Выдача паники"

var action_lst = [4]string{"+", "-", "*", "/"}
var nums_arabsk = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var nums_rimsk = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var new_lst_arabsk [2]string
var new_lst_rimsk [2]string
var arab_flag [2]string
var rimsk_flag [2]string
var new_flag [2]string
var new_list [2]string
var res_action string
var num int
var num1 int
var num2 int
var new_list_int [2]int

// For i in ... Сделаем функцию, чтобы такое потом использовать
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func convertRimskArab(s string) string {
	dict_arab_rimsk := map[string]string{
		"1": "I", "2": "II", "3": "III", "4": "IV", "5": "V",
		"6": "VI", "7": "VII", "8": "VIII", "9": "IX", "10": "X",
	}
	for k, v := range dict_arab_rimsk {

		// fmt.Println("value", v)
		if s == v {
			s = k
		}
	}
	return s
}

// На входе строка на выходе две цифры, действие - строка и флаг - строка
func findAction(s string) ([2]string, [2]string, string, error) {

	for i := 0; i < len(s); i += 1 {
		if stringInSlice(string([]rune(s)[i]), action_lst[:]) == true { // Если есть дейстие
			res_action = string([]rune(s)[i])
			res := strings.Split(s, string([]rune(s)[i]))
			fmt.Println(res)
			for i, j := range res {
				if stringInSlice(j, nums_arabsk[:]) {
					new_lst_arabsk := append(new_lst_arabsk[:i], j)
					arab_flag := append(arab_flag[:i], j)
					fmt.Println("new_lst_arabsk===arab_flag=========", new_lst_arabsk, arab_flag)
				} else if stringInSlice(j, nums_rimsk[:]) {
					new_lst_rimsk := append(new_lst_rimsk[:i], j)
					rimsk_flag := append(rimsk_flag[:i], j)
					fmt.Println("new_lst_arabsk, rimsk_flag-----------", new_lst_rimsk, rimsk_flag)
				}
				fmt.Println(i, j)
			}
			// fmt.Println("WE here  string([]rune(s)[i])", string([]rune(s)[i])) // +
			// fmt.Println(new_lst_rimsk, rimsk_flag)
		} else if stringInSlice(string([]rune(s)[i]), action_lst[:]) == false {
			fmt.Println("=========================00000000000")
			return new_list, new_flag, res_action, errors.New("Выдача паники, действие должно быть")
		}
	}
	if stringInSlice(new_lst_rimsk[0], nums_rimsk[:]) {
		new_list = new_lst_rimsk
		new_flag = rimsk_flag
		// return new_lst_rimsk, rimsk_flag, nil //  если есть римск, то римские возвращаем
		//errors.New("Выдача паники, действие должно быть")
	} else if stringInSlice(new_lst_arabsk[0], nums_arabsk[:]) {
		new_list = new_lst_arabsk
		new_flag = arab_flag
		// return new_lst_arabsk, arab_flag, nil
	}
	return new_list, new_flag, res_action, nil
}

func convertListRimskArab(lst_in [2]string) [2]string {
	for i, value := range new_list {
		if stringInSlice(value, nums_rimsk[:]) {
			fmt.Println("new_list[:i]", new_list[i], reflect.TypeOf(new_list[i]))
			arab_value := convertRimskArab(new_list[i])
			new_list[i] = arab_value
		}
		fmt.Println("new_list", new_list)
	}
	return new_list
}
func convertStrNum(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Here will insert smthing")
		panic(err)
	}
	return i
}
func calc(new_list [2]string, new_flag [2]string, res_action string) [2]int {
	for i, value := range new_list {
		num = convertStrNum(value)
		new_list_int[i] = num
	}
	num1 = new_list_int[0]
	num2 = new_list_int[1]
	dic_operats := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	a, b := num1, num2

	for op, fv := range dic_operats {
		if op == res_action {
			// fmt.Printf("%d %s %d = %d\n", a, op, b, fv(a, b))
			fmt.Printf("%d\n", fv(a, b))
		}
	}
	return new_list_int
}

func main() {
	fmt.Println("Input")
	var strIn string
	fmt.Scanln(&strIn)
	new_list, new_flag, res_action, _ = findAction(strIn)
	convertListRimskArab(new_list)
	fmt.Println("new_list changed", new_list, res_action)
	fmt.Println("Output")
	calc(new_list, new_flag, res_action)
}

// func increment() func() int {
// 	count := 0
// 	return func() int {
// 		count++
// 		return count
// 	}
// }
//Цифры из строки
////////////////////////////////
// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	str := "1,3,5,6,8,42"

// 	strs := strings.Split(str, ",")
// 	var ints []int
// 	for _, s := range strs {
// 		num, err := strconv.Atoi(s)
// 		if err == nil {
// 			ints = append(ints, num)
// 		}
// 	}

// 	fmt.Println(ints)
// }
//////////////////////////////////////////
// import (
// 	"fmt"
// 	"regexp"
// )
//Находим цифры в строке
// func main() {
// 	str := "Hello, today is January 11. I am 28 now. Date of birth 11.01.1990."

// 	fmt.Println("Test string", str)

// 	re := regexp.MustCompile("[0-9]+")
// 	fmt.Println(re.FindAllString(str, -1))
// }
// Test string Hello, today is January 11. I am 28 now. Date of birth 11.01.1990.
// [11 28 11 01 1990]
//////////////////////////////////////

// Разделяем строку на элементы
// func main() {
// 	str := "ab£"
// 	chars := []rune(str)
// 	for i := 0; i < len(chars); i++ {
// 		char := string(chars[i])
// 		println(char)
// 	}
// }

// import "fmt"

// "reflect"

// "flag"
// "errors"

// "os"
// //////////////////////////
// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strings"
// )

// func main() {
// 	s := "This,is,a,delimited,string"
// 	v := strings.Split(s, ",")
// 	fmt.Println(reflect.TypeOf(v)) // [This is a delimited string]

// }

///////////////////////////////
// func main() {
// 	matrix := make([][]int, 10)
// 	for x := 0; x < 10; x++ {
// 		for y := 0; y < 10; y++ {
// 			matrix[y] = make([]int, 10)
// 			matrix[y][x] = x
// 		}
// 		fmt.Println((matrix[x]))
// 	}

// }

// 	inc := increment()
// 	fmt.Println(inc())
// 	fmt.Println(inc())
// 	fmt.Println(inc())
// 	fmt.Println(reflect.TypeOf(inc()))

// 	inc2 := increment2()
// 	fmt.Println(inc2)
// 	fmt.Println(inc2)
// 	fmt.Println(reflect.TypeOf(inc2))

// 	// fmt.Println(findMin(1, 2, 3, 45, 123, 21, -4))
// 	// func() {
// 	// 	fmt.Println("анонимная функц")
// 	// }()
// }
// func increment() func() int {
// 	count := 0
// 	return func() int {
// 		count++
// 		return count
// 	}
// }
// func increment2() int {
// 	count := 0
// 	count++
// 	return count
// }

// func findMin(numbers ...int) int {
// 	if len(numbers) == 0 {
// 		return 0
// 	}
// 	min := numbers[0]
// 	for _, i := range numbers {
// 		if i < min {
// 			min = i
// 		}
// 	}
// 	return min
// }

// message, err := enterTheClub((70))
// if err != nil {
// 	// log.Fatal(err)
// 	fmt.Println(err)
// 	return
// }

// fmt.Println(message)
// fmt.Println(enterTheClub((70)))
// fmt.Println(prediction("фавимаи"))

// func enterTheClub(age int) (string, error) {
// 	if age >= 18 && age <= 45 {
// 		return "Входи аккуратно", nil
// 	} else if age > 45 && age < 65 {
// 		return "Вам не понравится музыка", nil
// 	} else if age >= 65 {
// 		return "Нет входа", errors.New("too old")
// 	}
// 	return "Нет входа", errors.New("too young")
// }
// func prediction(dayOfWeek string) (string, error) {
// 	switch dayOfWeek {
// 	case "пн":
// 		return "Хорошего пн", nil
// 	case "вт":
// 		return "Хорошего вт", nil
// 	case "ср":
// 		return "Хорошего ср", nil
// 	case "чт":
// 		return "Хорошего чт", nil
// 	case "пт":
// 		return "Хорошего пт", nil
// 	default:
// 		return "Невалидный день недели", errors.New("Invalid day")
// 	}
// }

// package main

// import (
// 	// "flag"
// 	"fmt"
// 	// "os"
// )

// func main() {
// 	// message := ""
// 	// flag := true
// 	message, _ := enterTheClub((19))
// 	fmt.Println(message)
// 	fmt.Println(enterTheClub((19)))
// 	// var message string
// 	// message = sayHello("Вася", 35)
// 	// printMessage(message)

// 	// messagefor := sayHello("Maxim")
// 	// printMessage(messagefor)
// 	// programName := os.Args[0]
// 	// fmt.Println(programName)
// 	// var name string
// 	// flag.StringVar(&name, "n", "admin", "Specify name. Default is admin.")

// 	// flag.Usage = func() {
// 	// fmt.Printf("Usage of our Program: \n")
// 	// fmt.Printf("./go-project -n username\n")
// 	// // flag.PrintDefaults() // prints default usage
// 	// }
// 	// flag.Parse()
// 	// print()
// 	// printMessage("Smthing here")
// 	// const message0 = "Vidacha paniki"
// 	// message0 = "Выдача паники"
// 	// var message string // message := ” the same
// 	// message = "Ia skoro stanu ...."
// 	// message := []byte("Я долго буду гнать велосипед ...")
// 	// var a rune = 'a'
// 	// b := byte(0x61)
// 	// bs := []byte{b}
// 	// a := rune(62)
// 	// fmt.Println(message, message0)
// 	// fmt.Println(reflect.TypeOf(message), reflect.TypeOf(message0))
// 	// fmt.Printf("%c\n", a)
// 	// fmt.Println(a)
// 	// fmt.Printf("%c\n", b)
// }
// func printMessage(message string) {
// 	fmt.Println(message)
// }
// func sayHello(name string, age int) string {
// 	d := ""
// 	d = fmt.Sprintf("Ваше имяб %s! вам %d лет", name, age)
// 	// fmt.Println(d)
// 	return d
// }

// // func print() {
// // a := rune(62)
// // fmt.Printf("%c\n", a)
// // }
// // func printMessage(message string) {

// // 	fmt.Println(message)
// // }

// //	func sayHello(name string) string {
// //		return "Privet " + name
// //	}
// func enterTheClub(age int) (string, bool) {
// 	if age >= 18 && age <= 45{
// 		return "Входи аккуратно", true
// 	} else if age > 45 && age < 65 {
// 		return "Вам не понравится музыка", true
// 	} else if age >= 65 {
// 		return "Нет входить"
// 	}

// 	return "Нет входа", false

// }

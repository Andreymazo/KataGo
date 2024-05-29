package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

const vadacha_panici string = "Выдача паники"

var action_lst = [4]string{"+", "-", "*", "/"}
var nums_arabsk = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var nums_rimsk = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var new_lst_arabsk []string
var new_list_arab_from_rimsk []string
var new_lst_rimsk []string
var arab_flag []string
var rimsk_flag []string
var new_flag []string
var new_list []string
var res string
var res_action string
var res_action_lst []string
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

// На входе строка на выходе три списка: действие, арабск, римск
// 3 условия на входе: если действие, если не действие и есть в списке арабск и если не действие и есть в списке римск
func findAction(strIn string) ([]string, []string, []string, string, error) {

	for i := 0; i < len(strIn); i += 1 {
		res = string([]rune(strIn)[i])
		if stringInSlice(string([]rune(strIn)[i]), action_lst[:]) == true { // Если есть дейстие
			res_action = res
			res_action_lst = append(res_action_lst, res)
			fmt.Println("res_action_lst", res_action_lst)
		} else if stringInSlice(string([]rune(strIn)[i]), nums_arabsk[:]) == true { // Если цифра арабская
			new_lst_arabsk = append(new_lst_arabsk, res)
			fmt.Println("new_lst_arabsk", new_lst_arabsk)

			// rimsk_flag := append(rimsk_flag[:i], j)
		} else if stringInSlice(string([]rune(strIn)[i]), nums_rimsk[:]) == true { // Если цифра римск
			new_lst_rimsk = append(new_lst_rimsk, res)
			// fmt.Println("new_lst_rimsk", new_lst_rimsk)
		}

	}
	fmt.Println("new_list, new_flag, res_action===", new_lst_arabsk, new_lst_rimsk, res_action_lst)
	if len(res_action_lst) > 1 {
		return new_lst_arabsk, new_lst_rimsk, res_action_lst, res_action, errors.New("Действие должно быть одно")
	} else {
		return new_lst_arabsk, new_lst_rimsk, res_action_lst, res_action, nil
	}
}

// На входе список арабск, список римс, действие, список действий ,на выходе: список араб, список-флаг, действие
func change_lst_to_flag(new_lst_arabsk []string, new_lst_rimsk []string, res_action_lst []string, res_action string) ([]string, []string, []string, string) {
	fmt.Println("---------------;;;;;----------------", new_lst_rimsk, len(new_lst_rimsk), new_lst_arabsk, len(new_lst_arabsk))

	if len(new_lst_arabsk) > 1 && len(new_lst_rimsk) == 0 {
		new_list = new_lst_arabsk
		new_flag = new_lst_arabsk

	} else if len(new_lst_rimsk) > 1 && len(new_lst_arabsk) == 0 {
		fmt.Println("00000000000000000000000000000000")
		new_list = convertListRimskArab(new_lst_rimsk)
		new_flag = new_lst_rimsk
		fmt.Println("new_list_arab_from_rimsk", new_list_arab_from_rimsk)

	} else if len(new_lst_rimsk) == 0 && len(new_lst_arabsk) == 0 {
		errors.New("Нет чисел")
	}
	fmt.Println("new_list, new_flag, res_action_lst, res_action", new_list, new_flag, res_action_lst, res_action)
	return new_list, new_flag, res_action_lst, res_action
}

func convertListRimskArab(new_lst_rimsk []string) []string {
	for i, value := range new_lst_rimsk {
		if stringInSlice(value, nums_rimsk[:]) {
			fmt.Println("new_list[:i]", new_lst_rimsk[i], reflect.TypeOf(new_lst_rimsk[i]))
			arab_value := convertRimskArab(new_lst_rimsk[i])
			new_lst_rimsk[i] = arab_value
		}
		fmt.Println("new_lst_rimsk_converted", new_lst_rimsk)
	}
	return new_lst_rimsk
}
func convertStrNum(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Here will insert smthing")
		panic(err)
	}
	return i
}
func calc(new_list []string, new_flag []string, res_action string) [2]int {
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
	new_lst_arabsk, new_lst_rimsk, res_action_lst, res_action, _ = findAction(strIn)
	new_list, new_flag, res_action_lst, res_action = change_lst_to_flag(new_lst_arabsk, new_lst_rimsk, res_action_lst, res_action)
	fmt.Println("new_list before", new_list, res_action)
	fmt.Println("new_list changed", new_list, res_action)
	fmt.Println("Output")
	calc(new_list, new_flag, res_action)
}

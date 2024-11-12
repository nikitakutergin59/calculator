package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	tokens := tokenize(expression)
	postfix, err := infixToPostfix(tokens)
	if err != nil {
		return 0, err
	}
	return evaluatePostfix(postfix)
}

func tokenize(expr string) []string {
	var tokens []string
	var currentToken strings.Builder

	for _, char := range expr {
		switch char {
		case ' ':
			continue
		case '+', '-', '*', '/', '(', ')':
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, string(char))
		default:
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}

func infixToPostfix(tokens []string) ([]string, error) {
	var output []string
	var operators []string

	for _, token := range tokens {
		if isNumber(token) {
			output = append(output, token)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			operators = operators[:len(operators)-1]
		} else if isOperator(token) {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		} else {
			return nil, fmt.Errorf("invalid character")
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, errors.New("mismatched parentheses")
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

func evaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if isNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				stack = append(stack, a/b)
			default:
				return 0, fmt.Errorf("unknown operator: %s", token)
			}
		} else {
			return 0, fmt.Errorf("invalid token: %s", token)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}

func isNumber(token string) bool {
	if _, err := strconv.ParseFloat(token, 64); err == nil {
		return true
	}
	return false
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
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

// среднее арифметическое, дисперсия,
func cr_ar(array []int) (float64, float64) {
	sum := 0
	no_sum := len(array)
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	cRaR := float64(sum) / float64(no_sum)

	sum_sq_diff := 0.0
	for i := 0; i <= len(array)-1; i++ {
		diff := float64(array[i]) - cRaR
		sum_sq_diff += diff * diff
	}
	variance := sum_sq_diff / float64(no_sum)
	return cRaR, variance
}

// наименьшее наибольшее, размах
type Stats struct {
	Max    float64
	Min    float64
	Razmax float64
}

func maXmiN(array []int) Stats {
	Max := array[0]
	Min := array[0]
	for i := 0; i <= len(array)-1; i++ {
		if array[i] > Max {
			Max = array[i]
		} else if array[i] < Min {
			Min = array[i]
		}
	}
	razmax := float64(Max) - float64(Min)
	return Stats{float64(Max), float64(Min), razmax}
}

// частота
type FrequencyResult struct {
	Value          int
	ValueFrequency int
}

func frequency(array []int) FrequencyResult {
	frequencyMap := make(map[int]int) // Используем карту для подсчета частоты

	for _, value := range array {
		frequencyMap[value]++
	}
	// Находим самое часто встречающееся значение
	maxFrequency := 0
	maxValue := 0
	for value, frequency := range frequencyMap {
		if frequency > maxFrequency {
			maxFrequency = frequency
			maxValue = value
		}
	}
	return FrequencyResult{maxValue, maxFrequency}
}

// вывод и форматирование значенний
// калькулятор
func Calculator(expression string) {
	fmt.Println("Введите выражение: ")
	fmt.Scanln(&expression)
	result, err := Calc(expression)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		roundedResult := strconv.FormatFloat(result, 'f', 2, 64)
		fmt.Println(roundedResult)
	}
}

// среднее арифметическое
func CrArMaxMin(s string) {
	fmt.Println("Введите произвольный список(через запятую): ")
	fmt.Scanln(&s)
	value := 0.0
	for _, char := range s {
		_, err := strconv.ParseFloat(string(char), 64)
		if err == nil {
			value++
		}
	}
	fmt.Println("Количество чисел в списке:", int(value))
	// Разбиваем строку по запятым
	numbers := strings.Split(s, ",")
	// Преобразование элементов массива в числа
	var m []int
	for _, numStr := range numbers {
		num, err := strconv.Atoi(strings.TrimSpace(numStr)) // Удаляем пробелы и преобразуем в число
		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			return
		}
		m = append(m, num)
	}
	fmt.Println("Ваш список:", m) // Выводим полученный массив
	//среднее арифметическое и дисперсия
	average, variance := cr_ar(m)
	// проверяем если у среднего более 2 знаков после запятой
	if average-float64(int(average)) != 0 { // Проверяем, есть ли дробная часть
		roundedResult_1 := strconv.FormatFloat(average, 'f', 2, 64)
		fmt.Println("Среднее арифметическое:", roundedResult_1)
	} else {
		fmt.Println("Среднее арифметическое:", average)
	}
	// Проверяем, есть ли у дисперсии более 2 знаков после запятой
	if variance-float64(int(variance)) != 0 { // Проверяем, есть ли дробная часть
		roundedResult_2 := strconv.FormatFloat(variance, 'f', 2, 64)
		fmt.Println("Дисперсия:", roundedResult_2)
	} else {
		fmt.Println("Дисперсия:", variance)
	}
	// минимальное максимальное значение произвольного списка, размах произвольного списка
	result_2 := maXmiN(m)
	if result_2.Razmax > result_2.Min {
		str := fmt.Sprintf("Максимальное: %v; Минимальное: %v; Размах: %v, 👎значительный", result_2.Max, result_2.Min, result_2.Razmax)
		fmt.Println(str)
	} else if result_2.Razmax == 0 {
		str := fmt.Sprintf("Максимальное: %v; Минимальное: %v; Размах: %v", result_2.Max, result_2.Min, result_2.Razmax)
		fmt.Println(str)
	} else if result_2.Razmax == result_2.Min {
		str := fmt.Sprintf("Максимальное: %v; Минимальное: %v; Размах: %v, 👍номарльный", result_2.Max, result_2.Min, result_2.Razmax)
		fmt.Println(str)
	} else {
		str := fmt.Sprintf("Максимальное: %v; Минимальное: %v; Размах: %v, 👍незначительный", result_2.Max, result_2.Min, result_2.Razmax)
		fmt.Println(str)
	}
	// частота
	result_3 := frequency(m)
	str := fmt.Sprintf("Число: %d; повторяеться в списке %d раз", result_3.Value, result_3.ValueFrequency)
	fmt.Println(str)

}

// вход!!!!
func main() {
	fmt.Println("Калькулятор: 1; Теорема бeзу: 2; Cреднее, дисперсия, размах...: 3; Seting: 4")
	var num int
	fmt.Scanln(&num)
	if num == 1 {
		Calculator("")
		main()
	} else if num == 2 {
		Bezy()
		main()
	} else if num == 3 {
		CrArMaxMin("")
		main()
	} else if num == 4 {
		main()
	} else {
		fmt.Println("Неверный номер!!!!!")
		main()
	}
}

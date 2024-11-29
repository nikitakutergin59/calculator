package calculator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Эта функция является главным “входным пунктом” калькулятора. Она принимает строку expression, содержащую арифметическое выражение, и возвращает два значения:
// float64: Результат вычисления выражения (в случае успеха).
// error: Ошибка, если выражение некорректно или возникла ошибка во время вычисления (например, деление на ноль).
func Calc(expression string) (float64, error) {
	tokens, err := tokenize(expression) // <--- ИЗМЕНЕНО: Теперь получаем два значения
	if err != nil {
		return 0, err // Передаем ошибку дальше
	}
	postfix, err := infixToPostfix(tokens)
	if err != nil {
		return 0, err
	}
	result, err := evaluatePostfix(postfix)
	if err != nil {
		return 0, err // Передаем ошибку дальше
	}
	return result, nil
}

// Функция токенизации разбивает входную строку expr на отдельные токены. Она использует strings.Builder для эффективного построения токенов
func tokenize(expr string) ([]string, error) {
	var tokens []string
	var currentToken strings.Builder

	for _, char := range expr {
		switch {
		case char == ' ':
			continue
		case string(char) == "(":
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, string(char))
		case string(char) == ")":
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, string(char))
		case char == '+' || char == '-' || char == '*' || char == '/' || string(char) == "sqrt": //Обработка sqrt как токена
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
	// ВАЖНО: обрабатывает скобки вокруг аргумента sqrt
	i := 0
	for i < len(tokens) {
		if tokens[i] == "sqrt" && i+1 < len(tokens) && tokens[i+1] == "(" {
			j := i + 2
			arg := ""
			foundClosingParen := false
			for j < len(tokens) {
				if tokens[j] == ")" {
					foundClosingParen = true
					break
				}
				arg += tokens[j]
				j++
			}
			if !foundClosingParen {
				return nil, errors.New("missing closing parenthesis")
			}
			// ИЗМЕНЕНИЕ: Корректное удаление лишних токенов
			tokens = append(tokens[:i], append([]string{fmt.Sprintf("sqrt(%s)", arg)}, tokens[j+1:]...)...)
			i = 0 // Reset i
			continue
		}
		i++
	}
	return tokens, nil
}

// Эта функция — сердце алгоритма. Она преобразует токены из инфиксной нотации в постфиксную. Она использует стек (operators) для хранения операторов.
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
			return nil, fmt.Errorf("invalid character: %s", token)
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

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("cannot calculate square root of a negative number")
	}
	return math.Sqrt(x), nil
}

// Эта функция вычисляет значение арифметического выражения, представленного в постфиксной нотации (postfix).
// Она использует стек (stack) для хранения чисел.
func evaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if isNumber(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %w", err)
			}
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 && !strings.HasPrefix(token, "sqrt(") { // проверка на sqrt(x)
				return 0, errors.New("invalid expression")
			}
			if strings.HasPrefix(token, "sqrt(") {
				numStr := token[5 : len(token)-1]
				num, err := strconv.ParseFloat(numStr, 64)
				if err != nil {
					return 0, fmt.Errorf("invalid number in sqrt: %w", err)
				}
				result, err := sqrt(num)
				if err != nil {
					return 0, err
				}
				stack = append(stack, result)
				continue // Переходим к следующему токену после обработки sqrt
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

// определяет тип токена
func isNumber(token string) bool {
	if _, err := strconv.ParseFloat(token, 64); err == nil {
		return true
	}
	return false
}

// определяет тип токена
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || strings.HasPrefix(token, "sqrt(")
}

// определяет тип токена
func precedence(op string) int {
	if strings.HasPrefix(op, "sqrt(") {
		return 3
	}
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

// вывод и форматирование значенний
// калькулятор
func Calculator(message string) {
	var expression string // Переменная для хранения выражения
	for {                 // Бесконечный цикл для обработки нескольких выражений
		fmt.Println(message) // Выводим приветственное сообщение
		fmt.Print("Введите выражение (или 'menu' для выхода): ")
		fmt.Scanln(&expression) // Считываем выражение с клавиатуры

		// Если пользователь ввёл 'menu', выходим из цикла.
		if expression == "menu" {
			break
		}

		// Вычисляем выражение.  Предполагается, что функция Calc уже определена.
		result, err := Calc(expression)
		if err != nil {
			fmt.Println("Ошибка:", err) // Выводим сообщение об ошибке
		} else {
			fmt.Println(formatFloat(result)) // Выводим результат, отформатированный функцией formatFloat
		}
	}
}

// вычислений количества знаков после запятой
func formatFloat(num float64) string {
	s := strconv.FormatFloat(num, 'f', -1, 64)
	parts := strings.Split(s, ".")
	if len(parts) == 2 {
		decimalPart := parts[1]
		precision := 0
		for _, r := range decimalPart {
			if string(r) != "0" {
				precision++
			}
		}
		if precision > 0 {
			return fmt.Sprintf("%."+fmt.Sprint(precision)+"f", num)
		}
	}
	return s
}

package crar

import (
	"fmt"
	"strconv"
	"strings"
)

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

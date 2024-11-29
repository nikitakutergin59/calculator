package crar

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// среднее арифметическое, дисперсия,
func cr_ar(array []float64) (float64, float64, float64) {
	sum := 0.0
	no_sum := float64(len(array))
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	cRaR := sum / no_sum

	sum_sq_diff := 0.0
	for i := 0; i < len(array); i++ {
		diff := array[i] - cRaR
		sum_sq_diff += diff * diff
	}
	variance := sum_sq_diff / no_sum
	SqrtVariance := math.Sqrt(variance)
	return cRaR, variance, SqrtVariance
}

// наименьшее наибольшее, размах
type Stats struct {
	Max    float64
	Min    float64
	Razmax float64
}

func maXmiN(array []float64) Stats {
	Max := array[0]
	Min := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > Max {
			Max = array[i]
		} else if array[i] < Min {
			Min = array[i]
		}
	}
	razmax := Max - Min
	return Stats{Max, Min, razmax}
}

func Median(array []float64) float64 {
	NewArray := make([]float64, len(array))
	copy(NewArray, array) // Создаем копию, чтобы не менять исходный массив
	sort.Float64s(NewArray)
	if len(NewArray)%2 != 0 {
		return NewArray[len(NewArray)/2]
	} else {
		left_central_element := NewArray[len(NewArray)/2-1]
		right_central_element := NewArray[len(NewArray)/2]
		return (left_central_element + right_central_element) / 2
	}
}

// среднее арифметическое
func CrArMaxMinValue(s string) {
	// Разбиваем строку по запятым
	nambers := strings.Split(s, ",")
	// Преобразование элементов массива в числа
	var m []float64
	for _, numStr := range nambers {
		num, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64) // Удаляем пробелы и преобразуем в число
		if err != nil {
			fmt.Println("ошибка преобразования:", err)
			return
		}
		m = append(m, num)
	}
	fmt.Println("Ваш список:", m) // Выводим полученный массив
	// считаем сколько чисел в списке
	value := 0.0
	for _, char := range s {
		if char == ',' {
			value++
		}
	}
	fmt.Println("Кол-во чисел в списке:", int(value)+1)

	//среднее арифметическое и дисперсия
	average, variance, SqrtVariance := cr_ar(m)
	// проверяем если у среднего более 2 знаков после запятой
	if average-float64(int(average)) != 0 { // Проверяем, есть ли дробная часть
		roundedResult_1 := strconv.FormatFloat(average, 'f', 3, 64)
		fmt.Println("Среднее арифметическое:", roundedResult_1)
	} else {
		fmt.Println("Среднее арифметическое:", average)
	}
	// Проверяем, есть ли у дисперсии более 2 знаков после запятой
	if variance-float64(int(variance)) != 0 { // Проверяем, есть ли дробная часть
		roundedResult_2 := strconv.FormatFloat(variance, 'f', 3, 64)
		fmt.Println("Дисперсия:", roundedResult_2)
	} else {
		fmt.Println("Дисперсия:", variance)
	}
	// Проверяем есть ли у стандартного отклонения более 2 знаков после запятой
	if SqrtVariance-float64(int(SqrtVariance)) != 0 { // Проверяем есть ли дробная часть
		roundResult_3 := strconv.FormatFloat(SqrtVariance, 'f', 3, 64)
		fmt.Println("Стандартное отклонение:", roundResult_3)
	} else {
		fmt.Println("Стандартное отклонение:", SqrtVariance)
	}
	// Мода произвольного списка чисел
	result_1 := Median(m)
	if len(m)%2 != 0 {
		fmt.Println("Медиана(нечётного кол-во елементов):", result_1)
	} else {
		fmt.Println("Медиана(чётного кол-во елементов):", result_1)
	}

	// минимальное максимальное значение произвольного списка, размах произвольного списка
	result := maXmiN(m)
	if result.Razmax > result.Min {
		str := fmt.Sprintf("Максимальное: %v; Минимальное: %v; Размах: %v, 👎значительный", result.Max, result.Min, result.Razmax)
		fmt.Println(str)
	} else if result.Razmax == 0 {
		str := fmt.Sprintf("Максимальное: %v; Минимальное: %v; Размах: %v", result.Max, result.Min, result.Razmax)
		fmt.Println(str)
	} else if result.Razmax == result.Min {
		str := fmt.Sprintf("Максимальное: %v; Минимальное: %v; Размах: %v, 👍нормальный", result.Max, result.Min, result.Razmax)
		fmt.Println(str)
	} else {
		str := fmt.Sprintf("Максимальное: %v; Минимальное: %v; Размах: %v, 👍незначительный", result.Max, result.Min, result.Razmax)
		fmt.Println(str)
	}
}

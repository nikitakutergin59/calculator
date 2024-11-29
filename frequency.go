package frequency

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// частота
type FrequencyResult struct {
	Value         float64
	ValueQuantity int
}

// Подсчет частоты
func countFrequency(array []float64) []FrequencyResult {
	frequencyMap := make(map[float64]int)
	for _, value := range array {
		frequencyMap[value]++
	}
	var results []FrequencyResult
	for value, quantity := range frequencyMap {
		results = append(results, FrequencyResult{value, quantity})
	}
	return results
}

// Преобразование строки в массив чисел
func ParseNumbers(s string) ([]float64, error) {
	var numbers []float64
	valuesStr := strings.Split(s, ",")
	for _, valueStr := range valuesStr {
		value, err := strconv.ParseFloat(strings.TrimSpace(valueStr), 64)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования: %w", err)
		}
		numbers = append(numbers, value)
	}
	return numbers, nil
}

// Подсчет количества чисел в строке
func countNumbers(s string) int {
	value := len(strings.Split(s, ","))
	return value
}

// Нахождение моды (с учетом нескольких мод с максимальной частотой)
func FindModa(array []float64) []float64 {
	results := countFrequency(array)
	if len(results) == 0 {
		return []float64{} // Пустой массив, если входной массив пуст
	}

	maxQuantity := 0
	for _, result := range results {
		if result.ValueQuantity > maxQuantity {
			maxQuantity = result.ValueQuantity
		}
	}

	var moda []float64
	for _, result := range results {
		if result.ValueQuantity == maxQuantity {
			moda = append(moda, result.Value)
		}
	}

	// Только если несколько чисел имеют максимальную частоту, тогда возвращается весь список мод
	sort.Float64s(moda)

	return moda

}

// Форматирование вывода
func FormatFrequency(s string) {
	numbers, err := ParseNumbers(s)
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	numCount := countNumbers(s)
	fmt.Printf("Количество чисел в списке: %d\n", numCount)

	results := countFrequency(numbers)
	for _, result := range results {
		ratio := float64(result.ValueQuantity) / float64(numCount)
		if result.Value-float64(int64(result.Value)) != 0 {
			fmt.Printf("Число %.3f повторяется %d раз. Частота повторения: %.3f\n", result.Value, result.ValueQuantity, ratio)
		} else {
			fmt.Printf("Число %d повторяется %d раз. Частота повторения: %.3f\n", int64(result.Value), result.ValueQuantity, ratio)
		}
	}
	moda := FindModa(numbers)
	fmt.Println("Мода:", moda)
}

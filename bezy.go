package bezy

import (
	"fmt"
	"strings"
)

// теорема безу
// делители свобоного члена
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Chlen struct {
	Chlen_0 string
	Chlen_1 string
	Chlen_2 string
	Chlen_3 int
}

func (s Chlen) TheDivisorOfTheFreeMultiplier(Chlen_3 int) []int {
	all := []int{}
	for i := 1; i <= abs(s.Chlen_3); i++ {
		if s.Chlen_3%i == 0 {
			all = append(all, i)
			all = append(all, -i)
		}
	}
	return all
}

func Bezy() {
	expression_1 := Chlen{Chlen_0: "x^3", Chlen_1: "7x^2", Chlen_2: "-4x", Chlen_3: -28}
	divisors := expression_1.TheDivisorOfTheFreeMultiplier(expression_1.Chlen_3)
	divisorsString := strings.Join(strings.Split(fmt.Sprint(divisors), " "), "; ")
	fmt.Println("Делители свободного члена:", divisorsString)
}

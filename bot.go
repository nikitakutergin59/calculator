package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"githab.com/nikitakutergin59/calculator/calculator"
	"github.com/nikitakutergin59/calculator/bezy"
	"github.com/nikitakutergin59/calculator/cr_ar"
	"github.com/nikitakutergin59/calculator/diskriminant"
	"github.com/nikitakutergin59/calculator/frequency"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN") // Получаем токен из переменной окружения
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message != nil { // Проверяем, есть ли сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = processMessage(update.Message.Text) // Вызываем функцию обработки сообщения
			bot.Send(msg)
		}
	}
	var num string
	for {
		fmt.Println("\nМеню")
		fmt.Println("Калькулятор: 1")
		fmt.Println("Теорема Безу: 2")
		fmt.Println("Дискриминат: 3")
		fmt.Println("Среднее, дисперсия, размах: 4")
		fmt.Println("Частота: 5")
		fmt.Println("Выход: 0")

		fmt.Print("Выберите пункт меню: ")
		fmt.Scanln(&num)
		switch num {
		case "1":
			calculator.Calculator("")
		case "2":
			var Value_0, Value_1, Value_2, Value_3 float64
			fmt.Print("Введите коэффициент(x³): ")
			fmt.Scanln(&Value_0)
			fmt.Print("Введите коэффициент(x²): ")
			fmt.Scanln(&Value_1)
			fmt.Print("Ведите коэффициент(x): ")
			fmt.Scanln(&Value_2)
			fmt.Print("Введите свободный член: ")
			fmt.Scanln(&Value_3)
			expression := bezy.Chlen{Value_0: Value_0, Chlen_0: "x³", Value_1: Value_1, Chlen_1: "x²", Value_2: Value_2, Chlen_2: "x", Value_3: Value_3}
			bezy.Bezy(expression)
		case "3":
			var aStr, bStr, cStr string // Объявление переменных для ввода
			fmt.Print("Введите значение a: ")
			fmt.Scanln(&aStr)
			fmt.Print("Введите значение b: ")
			fmt.Scanln(&bStr)
			fmt.Print("Введите значение c: ")
			fmt.Scanln(&cStr)

			a, err := strconv.ParseFloat(aStr, 64)
			if err != nil {
				fmt.Println("Ошибка при вводе a:", err)
				continue // Или другой способ обработки ошибки
			}
			b, err := strconv.ParseFloat(bStr, 64)
			if err != nil {
				fmt.Println("Ошибка при вводе b:", err)
				continue // Или другой способ обработки ошибки
			}
			c, err := strconv.ParseFloat(cStr, 64)

			if err != nil {
				fmt.Println("Ошибка при вводе c:", err)
				continue // Или другой способ обработки ошибки
			}

			if a-float64(int64(a)) == 0 { // Проверяем, целое ли число a
				fmt.Printf("Ваше квадратное уравнение: %dx²", int64(a))
			} else {
				fmt.Printf("Ваше квадратное уравнение:%.2fx²", a)
			}
			if b-float64(int64(b)) == 0 { // Проверяем, целое ли число b
				if b >= 0 {
					fmt.Printf("+%dx", int64(b))
				} else {
					fmt.Printf("%dx", int64(b))
				}
			} else {
				if b >= 0 {
					fmt.Printf("+%.2fx", b)
				} else {
					fmt.Printf("%.2fx", b)
				}
			}
			if c-float64(int64(c)) == 0 { // Проверяем, целое ли число c
				if c >= 0 {
					fmt.Printf("+%d=0\n", int64(c))
				} else {
					fmt.Printf("%d=0\n", int64(c))
				}
			} else {
				if c >= 0 {
					fmt.Printf("+%.2f=0\n", c)
				} else {
					fmt.Printf("%.2f=0\n", c)
				}
			}

			discriminant := diskriminant.CalculateDiscriminant(a, b, c)
			roots, err := diskriminant.CalculateRoots(a, b, discriminant)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}

			fmt.Println("Корни уравнения:")
			for i, root := range roots {
				fmt.Printf("Корень %d: %v\n", i+1, root)
			}
		case "4":
			fmt.Print("Введите список чисел (через запятую): ")
			var s string
			fmt.Scanln(&s)
			crar.CrArMaxMinValue(s)
		case "5":
			fmt.Print("Введите список чисел (через запятую): ")
			var input string
			fmt.Scanln(&input)
			frequency.FormatFrequency(input)
		case "0":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный номер пункта меню.")
		}
	}
}

func processMessage(message string) string {
	parts := strings.Fields(message) // Разбиваем сообщение на слова

	if len(parts) == 0 {
		return "Введите команду"
	}

	command := parts[0] // Первое слово — команда

	switch command {
	case "/calc":
		//Обработка калькулятора
		return calculate(parts[1:]) //Вызов отдельной функции
	case "/bezu":
		//Обработка теоремы Безу
		return bezu(parts[1:]) //Вызов отдельной функции
	case "/discriminant":
		//Обработка дискриминанта
		return discriminant(parts[1:]) //Вызов отдельной функции
	case "/stats":
		return stats(parts[1:]) //Вызов отдельной функции
	case "/frequency":
		return frequencyCalc(parts[1:]) //Вызов отдельной функции
	case "/start":
		return "Привет! Я математический бот."
	default:
		return "Неизвестная команда."
	}
}

func calculate(args []string) string {
	//Ваш код калькулятора
	return "Рузультата калькулятора"
}

func bezu(args []string) string {
	//Ваш код Теоремы Безу, обрабатывающий args
	return "Результат Теоремы Безу"
}

func discriminant(args []string) string {
	//Ваш код дискриминанта, обрабатывающий args
	return "Результат дискриминанта"
}

func stats(args []string) string {
	//Ваш код статистики, обрабатывающий args
	return "Результат статистики"
}

func frequencyCalc(args []string) string {
	//Ваш код частоты, обрабатывающий args
	return "Результат частоты"
}

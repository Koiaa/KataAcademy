package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := calc()
	fmt.Println(result)
}

func calc() string {
	var aStr, op, bStr string
	var ans, counter int

	for {
		fmt.Println("Введите запрос")
		_, err := fmt.Scanf("%s %s %s\n", &aStr, &op, &bStr)
		if err != nil {
			fmt.Println("Ошибка ввода", err)
			continue
		}
		break
	}

	rim := map[string]string{"I": "1", "II": "2", "III": "3", "IV": "4", "V": "5", "VI": "6", "VII": "7", "VIII": "8", "IX": "9", "X": "10"}

	a, err := strconv.Atoi(aStr)
	if err != nil {
		counter++
		a, err = strconv.Atoi(rim[aStr])
		if err != nil {
			panic("Ошибка конвертации a")
		}
	} else if a < 1 || a > 10 {
		panic("Вводимое число меньше 1 или больше 10")

	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		counter++
		b, err = strconv.Atoi(rim[bStr])
		if err != nil {
			panic("Ошибка конвертации b")

		}
	} else if b < 1 || b > 10 {
		panic("Вводимое число меньше 1 или больше 10")

	}

	if counter == 1 {
		panic("Ошибка. Необходимо использовать числа одного алфавита")

	}

	switch op {
	case "+":
		ans = a + b
	case "-":
		ans = a - b
	case "*":
		ans = a * b
	case "/":
		ans = a / b
	default:
		panic("Введен неверный оператор:")

	}

	if counter == 2 && ans < 1 {
		panic("Ответ с использованием римских цифр не может быть отрицательным")

	}

	if counter == 2 {
		var ans int = ans
		rimStr := map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C"}
		rimList := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
		rimAns := ""
		for _, num := range rimList {
			for num <= ans {
				rimAns += rimStr[num]
				ans -= num
			}
		}
		return rimAns
	}

	return strconv.Itoa(ans)
}

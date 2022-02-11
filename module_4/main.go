package main

import "fmt"

func main() {
	task1()
	fmt.Println("====================")
	task2()
	fmt.Println("====================")
	task3()
	fmt.Println("====================")
	task4()
	fmt.Println("====================")
	task5()
}

//ЕГЭ
func task1() {
	fmt.Println("Баллы ЕГЭ.")
	fmt.Println("Введите результат первого экзамена:")
	var firstExamScore int
	fmt.Scan(&firstExamScore)
	fmt.Println("Введите результат второго экзамена:")
	var secondExamScore int
	fmt.Scan(&secondExamScore)
	fmt.Println("Введите результат третьего экзамена:")
	var thirdExamScore int
	fmt.Scan(&thirdExamScore)

	totalScore := firstExamScore + secondExamScore + thirdExamScore

	const passingScore = 275

	fmt.Printf("Сумма проходных баллов: %d\n", passingScore)
	fmt.Printf("Количество набранных баллов: %d\n", totalScore)

	if totalScore >= passingScore {
		fmt.Println("Вы поступили!")
	} else {
		fmt.Println("Вы не поступили!")
	}
}

//Три Числа
func task2() {
	fmt.Println("Три числа.")
	fmt.Println("Введите первое число:")
	var firstNum int
	fmt.Scan(&firstNum)
	fmt.Println("Введите второе число:")
	var secondNum int
	fmt.Scan(&secondNum)
	fmt.Println("Введите третье число:")
	var thirdNum int
	fmt.Scan(&thirdNum)

	const requiredNum = 5

	if firstNum > requiredNum || secondNum > requiredNum || thirdNum > requiredNum {
		fmt.Printf("Среди введённых чисел есть число больше %d\n", requiredNum)
	} else {
		fmt.Printf("Среди введённых чисел нет числа больше %d\n", requiredNum)
	}
}

//Банкомат
func task3() {
	fmt.Println("Банкомат.")
	fmt.Println("Введите сумму снятия со счёта:")
	var enteredSum int
	fmt.Scan(&enteredSum)

	const atmBanknoteDenomination = 100
	const maxPossibleWithdrawSum = 100_000

	if enteredSum < atmBanknoteDenomination || enteredSum > maxPossibleWithdrawSum {
		fmt.Printf("Сумма не может быть меньше %d рублей и больше %d рублей\n", atmBanknoteDenomination, maxPossibleWithdrawSum)
	} else {
		fmt.Println("Операция успешно выполнена.")
		fmt.Printf("Вы сняли %d рублей.\n", enteredSum)
	}
}

//Три числа: ещё попытка
func task4() {
	fmt.Println("Три числа.")
	fmt.Println("Введите первое число:")
	var firstNum int
	fmt.Scan(&firstNum)
	fmt.Println("Введите второе число:")
	var secondNum int
	fmt.Scan(&secondNum)
	fmt.Println("Введите третье число:")
	var thirdNum int
	fmt.Scan(&thirdNum)

	const requiredNum = 5

	var numArr []int
	numArr = append(numArr, firstNum, secondNum, thirdNum)

	numAmount := 0

	for _, v := range numArr {
		if v >= requiredNum {
			numAmount++
		}
	}

	if numAmount == 0 {
		fmt.Printf("Среди введённых чисел нет больших или равных %d.\n", requiredNum)
	} else {
		fmt.Printf("Среди введённых чисел %d больше или равны %d\n", numAmount, requiredNum)
	}
}

//Ресторан
func task5() {
	fmt.Println("Введите день недели:")
	var dayOfTheWeek string
	fmt.Scan(&dayOfTheWeek)
	fmt.Println("Введите число гостей:")
	var numberOfGuests float64
	fmt.Scan(&numberOfGuests)
	fmt.Println("Введите сумму чека:")
	var chequeTotal float64
	fmt.Scan(&chequeTotal)

	const (
		mondayDiscountPercentage    = 0.1
		fridayDiscountPercentage    = 0.05
		additionalServicePercentage = 0.1
		guestAmount                 = 5.
		sumToCheckOnFriday          = 10_000
	)

	mondayDiscount := chequeTotal * mondayDiscountPercentage
	fridayDiscount := chequeTotal * fridayDiscountPercentage
	additionalServicePrice := chequeTotal * additionalServicePercentage

	if dayOfTheWeek == "Понедельник" {
		chequeTotal = chequeTotal - mondayDiscount
	}

	if dayOfTheWeek == "Пятница" && chequeTotal > sumToCheckOnFriday {
		chequeTotal = chequeTotal - fridayDiscount
		fmt.Printf("Скидка по пятницам: %.1f\n", fridayDiscount)
	}

	if numberOfGuests > guestAmount {
		chequeTotal = chequeTotal + additionalServicePrice
		fmt.Printf("Надбавка на обслуживание: %.1f\n", additionalServicePrice)
	}

	fmt.Printf("Сумма к оплате: %.1f\n", chequeTotal)
}

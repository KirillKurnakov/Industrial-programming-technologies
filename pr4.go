package main

import (
	"fmt"
	"math"
)
	// 1.Задачи на линейное программирование (без условных операторов и циклов)
	//Сумма цифр числа
func sumDigits() {
	fmt.Println(" Сумма цифр числа ")
	var number int
	//fmt.Print("Введите 4-значное число: ")
	//fmt.Scan(&number)
	number=1234
	fmt.Println("Сумма цифр числа: ", number%10 + (number/10)%10 + 
	(number/100)%10 + (number/1000)%10)
}
	//Преобразование температуры
func temperature() {
	fmt.Println(" Преобразование температуры ")
	var tempC,tempF float32
	//fmt.Println("Введите температуру в Цельсиях")
	//fmt.Scan(&tempC)
	tempC=34
	fmt.Println("Температура в Цельсиях: ", tempC, " Температура в Фаренгейтах: ", ((tempC*9/5)+ 32))
	//fmt.Println("Введите температуру в Фаренгейтах")
	//fmt.Scan(&tempF)
	tempF=70
	fmt.Print("Температура в Фаренгейтах: ", tempF, " Температура в Цельсиях: ", ((tempF-32)*5/9))
}

	//Удвоение каждого элемента массива
func doubleMassives() {
	fmt.Println(" Удвоение каждого элемента массива ")
	var numbers [5] int = [5]int{1,2,3,4,5}
	/*fmt.Print("Введите 1-ое число: ")
	fmt.Scan(&numbers[0])
	fmt.Print("Введите 2-ое число: ")
	fmt.Scan(&numbers[1])
	fmt.Print("Введите 3-ое число: ")
	fmt.Scan(&numbers[2])
	fmt.Print("Введите 4-ое число: ")
	fmt.Scan(&numbers[3])
	fmt.Print("Введите 5-ое число: ")
	fmt.Scan(&numbers[4])*/
	fmt.Println("Исходный массив: ", numbers)
	numbers[0]=numbers[0]*2
	numbers[1]=numbers[1]*2
	numbers[2]=numbers[2]*2
	numbers[3]=numbers[3]*2
	numbers[4]=numbers[4]*2
	fmt.Println("Итоговый массив: ", numbers)
}

	//Объединение строк
func sumString() {
	fmt.Println(" Объединение строк ")
	var text,text1 string
	//fmt.Print("Введите первое слово: ")
	//fmt.Scan(&text)
	text="Привет"
	//fmt.Print("Введите второе слово: ")
	//fmt.Scan(&text1)
	text1="Мир!"
	fmt.Println(text + " " + text1)
}

	//Расчет расстояния между двумя точками
func distanceBtwDot() {
	fmt.Println(" Расчет расстояния между двумя точками ")
	var firstKoord [2] int
	var secondKoord [2] int
	//fmt.Print("Введите координаты первой точки (x,y): ")
	//fmt.Scan(&firstKoord[0])
	//fmt.Scan(&firstKoord[1])
	firstKoord[0]=1
	firstKoord[1]=1
	fmt.Println("Точка 1(x,y): ", firstKoord[0], " ", firstKoord[1])
	//fmt.Print("Введите координаты второй точки (x,y): ")
	//fmt.Scan(&secondKoord[0])
	//fmt.Scan(&secondKoord[1])
	secondKoord[0]=4
	secondKoord[1]=5
	fmt.Println("Точка 2(x,y): ", secondKoord[0], " ", secondKoord[1])
	fmt.Println("Результат: ", math.Sqrt((math.Pow(float64(secondKoord[0] - firstKoord[0]),2))+(math.Pow(float64(secondKoord[1] - firstKoord[1]),2))))
}

	// 2. Задачи с условным оператором
	//Проверка на четность/нечетность
func checkEvenOdd() {
	fmt.Println(" Проверка на четность/нечетность ")
	var number int
	//fmt.Print("Введите число: ")
	//fmt.Scan(&number)
	number=45
	fmt.Println("Число: ", number)
	if number%2==0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}

	//Проверка високосного года
func VisocosYear() {
	fmt.Println(" Проверка високосного года ")
	var number int
	//fmt.Print("Введите год: ")
	//fmt.Scan(&number)
	number=2024
	fmt.Println("Год: ", number)
	if number%4==0 {
		fmt.Println("Год високосный! ")
	} else {
		fmt.Println("Год невисокосный! ")
	}
}

	// Определение наибольшего из трех чисел
func maxNumber() {
	fmt.Println(" Определение наибольшего из трех чисел ")
	var numbers [3] int = [3]int{4,9,7}
	//fmt.Print("Введите 1-ое число: ")
	//fmt.Scan(&numbers[0])
	//fmt.Print("Введите 2-ое число: ")
	//fmt.Scan(&numbers[1])
	//fmt.Print("Введите 3-ое число: ")
	//fmt.Scan(&numbers[2])
	fmt.Println("Исходный массив: ", numbers)
	if numbers[0]>numbers[1] && numbers[0] > numbers[2] {
		fmt.Println("Наибольшее число: ", numbers[0])
	} else if (numbers[1]>numbers[0] && numbers[1] > numbers[2]) {
		fmt.Println("Наибольшее число: ", numbers[1])
	} else {
		fmt.Println("Наибольшее число: ", numbers[2])
	}
}

	//Категория возраста
func ageCheck() {
	fmt.Println(" Категория возраста ")
	var number int
	//fmt.Print("Введите возраст: ")
	//fmt.Scan(&number)
	//до 14 - ребенок
	//от 14 до 18 - подросток
	//от 18 до 60 - взрослый
	//после 60 - пожилой
	number=49
	fmt.Println("Возраст: ", number)
	if number<14 {
		fmt.Println("Ребенок")
	} else if number < 18 {
		fmt.Println("Подросток")
	} else if number < 60 {
		fmt.Println("Взрослый")
	} else {
		fmt.Println("Пожилой")
	}
}

	//Проверка делимости на 3 и 5
func check3and5() {
	fmt.Println(" Проверка делимости на 3 и 5 ")
	var number int
	//fmt.Print("Введите число: ")
	//fmt.Scan(&number)
	number=15
	fmt.Println("Число: ", number)
	if number%3==0 && number%5==0 {
		fmt.Println("Делится! ")
	} else {
		fmt.Println("Не делится! ")
	}
}

	//Задачи на циклы
	//Факториал числа
func factorial() {
	fmt.Println(" Факториал числа ")
	var number,finalN int
	finalN=1
	//fmt.Print("Введите число для вычисления факториала: ")
	//fmt.Scan(&number)
	number=5
	fmt.Println("Число: ", number)
	for i:=1; i < (number+1);i++ {
		finalN = finalN*i
	}
	fmt.Println("Факториал равен: ", finalN)
}

	//Числа Фибоначчи
func fibonachi() {
	fmt.Println(" Числа Фибоначчи ")
	var number int 
	var start1, start2, start3 int
	start1=0
	start2=1
	//fmt.Print("Введите количество чисел Фибоначчи: ")
	//fmt.Scan(&number)
	number=7
	fmt.Println("До числа: ", number)
	fmt.Print(start1, " ")
	for i:=1; i < number; i++ {
		fmt.Print(start2, " ")
		start3=start1
		start1=start2
		start2 = start3 + start2
	}
}

	//Реверс массива
func reversMassive() {
	fmt.Println(" Реверс массива ")
	var numbers [5] int = [5]int{1,2,3,4,5}
	/*fmt.Print("Введите 1-ое число: ")
	fmt.Scan(&numbers[0])
	fmt.Print("Введите 2-ое число: ")
	fmt.Scan(&numbers[1])
	fmt.Print("Введите 3-ое число: ")
	fmt.Scan(&numbers[2])
	fmt.Print("Введите 4-ое число: ")
	fmt.Scan(&numbers[3])
	fmt.Print("Введите 5-ое число: ")
	fmt.Scan(&numbers[4])*/
	fmt.Println("Исходный массив: ", numbers)
	var helper int
	for i:=0; i<2;i++ {
		helper = numbers[i]
		numbers[i] = numbers[4-i]
		numbers[4-i] = helper
	}
	fmt.Println("Итоговый массив: ", numbers)
}

	//Поиск простых чисел
func simpleNumber() {
	fmt.Println(" Поиск простых чисел ")
	var number,helper int
	helper=0
	//fmt.Print("Введите число: ")
	//fmt.Scan(&number)
	number=30
	fmt.Println("До числа: ", number)
	for j:=2; j <= number; j++ {
		for i:=1; i <= j; i++ {
			if (j%i==0) {
				helper++
			}
		}
		if helper==2 {
			fmt.Print(j, " ")
		}
		helper=0
	}
}

	// Сумма чисел в массиве
func sumMassive() {
	fmt.Println(" Сумма чисел в массиве ")
	var numbers [5] int = [5]int{1,2,3,4,5}
	/*fmt.Print("Введите 1-ое число: ")
	fmt.Scan(&numbers[0])
	fmt.Print("Введите 2-ое число: ")
	fmt.Scan(&numbers[1])
	fmt.Print("Введите 3-ое число: ")
	fmt.Scan(&numbers[2])
	fmt.Print("Введите 4-ое число: ")
	fmt.Scan(&numbers[3])
	fmt.Print("Введите 5-ое число: ")
	fmt.Scan(&numbers[4])*/
	fmt.Println("Исходный массив: ", numbers)
	var sum int
	for i:=0; i < 5; i++ {
		sum=sum+numbers[i]
	}
	fmt.Println("Итоговая сумма:", sum)
}

func main() {
	fmt.Println()
	sumDigits()
	fmt.Println()
	temperature()
	fmt.Println()
	doubleMassives()
	fmt.Println()
	sumString()
	fmt.Println()
	distanceBtwDot()
	fmt.Println()
	checkEvenOdd()
	fmt.Println()
	VisocosYear()
	fmt.Println()
	maxNumber()
	fmt.Println()
	ageCheck()
	fmt.Println()
	check3and5()
	fmt.Println()
	factorial()
	fmt.Println()
	fibonachi()
	fmt.Println()
	reversMassive()
	fmt.Println()
	simpleNumber()
	fmt.Println()
	sumMassive()
}
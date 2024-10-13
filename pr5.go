package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strconv"
	"sort"
)

func systSchil() {
	fmt.Println(" Перевод систем счисления ")
	var number,start1,end1,helper,helper1 int
	//Ввожу 100 в 36. Перевожу 100 из 36 в 5-чную
	fmt.Print("Введите трехзначное число, исходную систему счисления и конечную систему счисления: ")
	fmt.Scan(&number)
	fmt.Scan(&start1)
	fmt.Scan(&end1)
	helper =  ((number/100)%10) * start1*start1 + ((number/10)%10) * start1 + ((number)%10)
	helper1 = (helper%end1) + 10*((helper/end1)%end1) + 100*((helper/(end1*end1))%end1) + 1000*((helper/(end1*end1*end1))%end1) + 10000*((helper/(end1*end1*end1*end1))%end1)
	fmt.Print("100 из 36 в 5: ", helper1)
}

// решение квадратного уравнения
func SquareUrav() {
	fmt.Println("Введите a,b,c: ")
	var a,c, b, d, sq1,sq2 float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	d = b*b - 4*a*c
	sq1 = (-b + math.Sqrt(d)) / (2*a)
	sq2 = (-b - math.Sqrt(d)) / (2*a)
	fmt.Print("1 корень: ", sq1, " / 2 корень: ", sq2)
}

// Сортировка чисел по модулю
func sortModule(numbers []float64) []float64 {
	sort.Slice(numbers, func (i, j int) bool {
		return math.Abs(numbers[i]) < math.Abs(numbers[j])
	})
	return numbers
}

func sliyanie() {
	fmt.Println("Введите два отсортированных массива:")
	number := [5]int {1,2,3,4,5}
	number1 := [5]int {2,3,4,5,6}
	var number2 [10]int
	number2[0] = number[0]
	number2[1] = number[1] 
	number2[2] = number1[0] 
	number2[3] = number[2] 
	number2[4] = number1[1] 
	number2[5] = number[3] 
	number2[6] = number1[2] 
	number2[7] = number[4] 
	number2[8] = number1[3] 
	number2[9] = number1[4]
	fmt.Println("Итоговый массив:", number2)
}

func podStroka() {
	var str1, strIN1 string
	var helper,ii int
	helper = 0
	ii = 0
	fmt.Println("Введите строку и подстроку: ")
	fmt.Scan(&str1)
	fmt.Scan(&strIN1)
	str := []rune(str1)
	strIN := []rune(strIN1)
	for i:=0; i < len(str)-len(strIN)+1; i++ {
		for j :=0; j < len(strIN); j++ {
			if strIN[j] == str[i + j] {
				helper = helper + 1
				if j == 0 {
					ii = i+j
				}
			}
		}
		if helper == len(strIN) {
			fmt.Println("Индекс первого вхождения: ", ii)
			break
		} else {
			helper = 0
		}
	}
	if helper == 0 {
		fmt.Println("-1")
	}
}

//Калькулятор
func calc() {
	var operation string
	var a,b int
	fmt.Print("Введите операцию (+, -, *, /, ^, %): ")
	fmt.Scan(&operation)
	for operation != "+" && operation != "-" && operation != "*" && operation != "/" && operation != "^" && operation != "%" {
		fmt.Print("Неверная операция: ")
		fmt.Print("Введите операцию (+, -, *, /, ^, %): ")
		fmt.Scan(&operation)
	}
	fmt.Print("Введите числа для операции: ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	if operation == "+" {
		fmt.Println("Результат операции: ", a + b)
	}
	if operation == "-" {
		fmt.Println("Результат операции: ", a - b)
	}
	if operation == "*" {
		fmt.Println("Результат операции: ", a * b)
	}
	if operation == "/" {
		if b == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println("Результат операции: ", a / b)
		}
	}
	if operation == "^" {
		fmt.Println("Результат операции: ", a ^ b)
	}
	if operation == "%" {
		fmt.Println("Результат операции: ", a % b)
	}
}

func palindrom() bool {
	var strPal string
	fmt.Print("Введите слово: ")
	fmt.Scan(&strPal)
	runes := []rune(strPal)
	for i:=0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

func perecechenieLiniy() {
	var a1, a2, b1, b2, c1, c2 float64
	fmt.Println("Введите координаты первого отрезка (a1 a2):")
	 fmt.Scan(&a1, &a2)
 	fmt.Println("Введите координаты второго отрезка (b1 b2):")
 	fmt.Scan(&b1, &b2)
 	fmt.Println("Введите координаты третьего отрезка (c1 c2):")
 	fmt.Scan(&c1, &c2)
	 // Убедимся, что начальная точка меньше конечной для каждого отрезка
	if a1 > a2 {
		a1, a2 = a2, a1
	}
	if b1 > b2 {
		b1, b2 = b2, b1
	}
	if c1 > c2 {
		c1, c2 = c2, c1
	}
	// Находим максимальную начальную точку
	start := math.Max(float64(math.Max(a1, b1)), c1)
	// Находим минимальную конечную точку
	end := math.Min(float64(math.Min(a2, b2)), c2)
	if start <= end {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func maxLength() {
	var current  string
	var maxString string
	var j,maxJ int
	fmt.Println("Введите предложение: ")
	inputPred, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(inputPred)
	j=0
	maxJ=0
	current = ""
	for i:=0; i < len(runes); i++ {
		if (runes[i] != ',' && runes[i] != '.' && runes[i] != '!' && runes[i] != '?' && runes[i] != ' ') || i+1 == len(runes){
			current = current + string(runes[i])
			j = j + 1
			if i+1 == len(runes) {
				if j > maxJ {
					maxString = current
					current = ""
					maxJ=j
					j = 0
				}
			}
		} else {
			if j > maxJ {
				maxString = current
				current = ""
				maxJ=j
				j = 0
			} else {
				current = ""
				j = 0
			}
			continue
		}
	}
	fmt.Println("Максимальное слово:", maxString)
}

func VisocosYear1() {
	fmt.Println(" Проверка високосного года ")
	var number int
	fmt.Print("Введите год: ")
	fmt.Scan(&number)
	fmt.Println("Год: ", number)
	if number%4==0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func fibonachi1() {
	fmt.Println(" Числа Фибоначчи ")
	var number int 
	var start1, start2, start3 int
	start1=0
	start2=1
	fmt.Print("Введите количество чисел Фибоначчи: ")
	fmt.Scan(&number)
	fmt.Println("До числа: ", number)
	fmt.Print(start1, " ")
	for start2 < number {
		fmt.Print(start2, " ")
		start3=start1
		start1=start2
		start2 = start3 + start2
	}
}
	
func simpleNumber1() {
	fmt.Println(" Поиск простых чисел ")
	var number,number1,helper int
	helper=0
	fmt.Print("Введите два числа: ")
	fmt.Scan(&number)
	fmt.Scan(&number1)
	for j:=number; j <= number1; j++ {
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

func ArmstrongNumber() {
	fmt.Println(" Числа Армстронга ")
	var number,number1,h,helper int
	var helper1 float64
	helper=0
	helper1=0
	fmt.Print("Введите два числа: ")
	fmt.Scan(&number)
	fmt.Scan(&number1)
	for j:=number; j <= number1; j++ {
		helper = j
		h = len(strconv.Itoa(helper)) // степень
		for helper > 0 {
			helper1 = helper1 + math.Pow(float64(helper%10),float64(h))
			helper = helper/10
		}
		if helper1 == float64(j) {
			fmt.Print(" ", j)
		}
		helper1=0
	}
}

func ReversString() {
	fmt.Print("Введите строку: ")
	inputPred, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(inputPred)
	for i := len(runes)-1; i >= 0; i-- {
		fmt.Print(string(runes[i]))
	}
}

func Evklid() {
	var num,num1 int
	fmt.Print("Введите два числа для нахождения НОД: ")
	fmt.Scan(&num)
	fmt.Scan(&num1)
	for {
		if num!=0 && num1!=0 {
			if num > num1 {
				num = num % num1
			} else {
				num1 = num1 % num
			}
		} else{
			fmt.Print("НОД: ", num + num1)
			break
		}
	}
}

func main() {
	systSchil()
	SquareUrav()
	
	numbers := []float64{-5, 3, -1, 4, -2, 0, 7}
	fmt.Println("Исходный массив: ", numbers)
	fmt.Println("Отсортированный массив: ", sortModule(numbers))
	
	sliyanie()
	podStroka()
	calc()
	fmt.Print(palindrom())
	perecechenieLiniy()
	maxLength()
	VisocosYear1()
	fibonachi1()
	simpleNumber1()
	ArmstrongNumber()
	ReversString()
	Evklid()
}

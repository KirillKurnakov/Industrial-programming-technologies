package main

import (
	"fmt"
	"math"
)

// 1. Функция вычисления площади треугольника
func SquareTriangle(base float64, height float64) float64 {
	return 0.5 * (base * height)
}

// 2. Сортировка массива по возрастанию
func sortArray(arr []int) []int {
	fmt.Println("Исходный массив: ", arr)
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 3. Сумма квадратов чётных чисел
func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += int(math.Pow(float64(i), 2))
		}
	}
	return sum
}

// 4. Проверка палиндрома
func isPalindrome(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

// 5. Проверка числа на простоту
func isPrime(number int) bool {
	helper := 0
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			helper++
		}
		if helper >= 2 && i != number {
			return false
		}
	}
	if helper <= 2 {
		return true
	}
	return true
}

// 6. Генерация последовательности простых чисел
func generatePrimes(limit int) []int {
	var numbers []int
	for i := 1; i <= limit; i++ {
		if isPrime(i) {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

// 7. Перевод числа в двоичную систему
func toBinary(n int) string {
	var result string
	for i := 0; n > 0; i++ {
		result = fmt.Sprintf("%d", n%2) + result
		n /= 2
	}
	return result
}

// 8. Нахождение максимального элемента в массиве
func findMax(arr []int) int {
	maxN := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxN {
			maxN = arr[i]
		}
	}
	return maxN
}

// 9. Функция вычисления НОД (наибольший общий делитель)
func gcd(num int, num1 int) int {
	for {
		if num != 0 && num1 != 0 {
			if num > num1 {
				num = num % num1
			} else {
				num1 = num1 % num
			}
		} else {
			return num + num1
		}
	}
}

// 10. Сумма элементов массива
func sumArray(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	//fmt.Println("Площадь треугольника: ", SquareTriangle(5, 9))

	//arr := []int{64, 34, 25, 12, 22, 11, 90}
	//fmt.Println("Сортировка", sortArray(arr))

	//var sum int
	//fmt.Scan(&sum)
	//fmt.Println("Сумма квадратов до ", sum, " равна: ", sumOfSquares(sum))

	/*var strPal string
	fmt.Print("Введите слово: ")
	fmt.Scan(&strPal)
	fmt.Println(isPalindrome(strPal))*/

	/*var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Println(isPrime(number))*/

	/*var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Println(generatePrimes(number))*/

	/*var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Println("Число ", number, " в 2-ной системе счисления: ", toBinary(number))*/

	//numbers := [...]int{1,2,3,4,5,6,7,23,4,23,43223,0,2}
	//fmt.Println("Максимальное число в массиве: ", findMax(numbers[:]))

	/*var num, num1 int
	fmt.Print("Введите два числа для нахождения НОД: ")
	fmt.Scan(&num)
	fmt.Scan(&num1)
	fmt.Print("НОД для ", num, " и ", num1, " равен: ", gcd(num,num1))*/

	//numbers := [...]int{1, 2, 3, 4, 5}
	//fmt.Println("Сумма чисел в массиве ", numbers, " равна: ", sumArray(numbers[:]))
}

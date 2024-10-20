package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"bufio"
	"os"
	"time"
)

// 1. Простые числа
func simpleNumber2() {
	fmt.Println(" Поиск простых чисел ")
	var number, helper int
	helper = 0
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			helper++
		}
		if helper >= 2 && i != number {
			fmt.Println("Число не простое!")
			fmt.Println("Первый найденный делитель: ", i)
			break
		}
	}
	if helper <= 2 {
		fmt.Println("Число простое!")
	}
}

// 2. НОД (Евклид)
func Evklid1() {
	var num, num1 int
	fmt.Print("Введите два числа для нахождения НОД: ")
	fmt.Scan(&num)
	fmt.Scan(&num1)
	for {
		if num != 0 && num1 != 0 {
			if num > num1 {
				num = num % num1
			} else {
				num1 = num1 % num
			}
		} else {
			fmt.Print("НОД: ", num+num1)
			break
		}
	}
}

// 3. Сортировка пузырьком
func Bubble() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Исходный массив: ", arr)
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Отсортированный массив: ", arr)
}

// 4. Таблица умножения в формате матрицы
func Tablica() {
	const size = 10

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

// 5. Фибоначчи с мемоизацией
func fibonachi2(start1 int, start2 int, number int) {
	var start3 int
	if start2 <= number {
		fmt.Print(start2, " ")
		start3 = start1
		start1 = start2
		start2 = start3 + start2
		fibonachi2(start1, start2, number)
	}
}

// 6. Обратные числа
func ReversDigits() {
	var number int
	fmt.Println("Введите целое число: ")
	fmt.Scan(&number)
	for number > 0 {
		fmt.Print(number % 10)
		number = number / 10
	}
}

// 7. Треугольник Паскаля
func Pascal() {
	var levels int
	fmt.Print("Введите количество уровней треугольника Паскаля: ")
	fmt.Scan(&levels)

	pascal := make([][]int, levels)

	// Заполняем массив
	for i := 0; i < levels; i++ {
		// Инициализируем строку
		pascal[i] = make([]int, i+1)
		pascal[i][0] = 1
		pascal[i][i] = 1

		for j := 1; j < i; j++ {
			pascal[i][j] = pascal[i-1][j-1] + pascal[i-1][j]
		}
	}

	// Выводим треугольник Паскаля
	for i := 0; i < levels; i++ {
		for j := 0; j < levels-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", pascal[i][j])
		}
		fmt.Println()
	}
}

// 8. Число палиндром
func palindrome() {
	var number, helper, number1, lenN, check int
	fmt.Println("Введите число: ")
	fmt.Scan(&number)
	number1 = number
	check = 0
	for number > 0 {
		helper = number % 10
		number = number / 10
		lenN = lenN + 1
	}
	for i := 0; i <= lenN/2; i++ {
		if number1%10 != helper%10 && helper > 0 {
			fmt.Println("Число не палиндром!")
			check = 1
			break
		} else {
			helper = helper / 10
			number1 = number1 / 10
		}
	}
	if check == 0 {
		fmt.Println("Число палиндром!")
	}
}

// 9.Макс и мин в массиве
func MaxMinMassive(numbers []int) {
	var minN,maxN int
	minN=numbers[0]
	maxN=numbers[0]
	for i:=0; i < len(numbers); i++ {
		if numbers[i] < minN {
			minN = numbers[i]
		}
		if numbers[i] > maxN {
			maxN = numbers[i]
		}
	}
	fmt.Println("Максимальное число: ", maxN)
	fmt.Println("Минимальное число: ", minN)
}

//10. Игра "Угадай число"
func FindDigital() {
	var number int64
	var helper int64
	number = rand.Int63n(100) + 1
	fmt.Println(number)
	fmt.Println("Машина загадала число. У вас 10 попыток, чтобы отгадать!")
	for i:=10; i > 0; i-- {
		fmt.Println("Введите число от 1 до 100: ")
		fmt.Scan(&helper)
		if helper == number {
			fmt.Println("Вы угадали число!")
			break
		} else if helper > number {
			fmt.Println("Введенное число больше. Осталось попыток: ", i-1)
		} else if helper < number {
			fmt.Println("Введенное число меньше. Осталось попыток: ", i-1)
		}
	}
}

// 11. Армстронг
func ArmstrongNumber1() {
	fmt.Println(" Числа Армстронга ")
	var number, number1, h, helper int
	var helper1 float64
	helper = 0
	helper1 = 0
	fmt.Print("Введите два числа: ")
	fmt.Scan(&number)
	fmt.Scan(&number1)
	for j := number; j <= number1; j++ {
		helper = j
		h = len(strconv.Itoa(helper)) // степень
		for helper > 0 {
			helper1 = helper1 + math.Pow(float64(helper%10), float64(h))
			helper = helper / 10
		}
		if helper1 == float64(j) {
			fmt.Print(" ", j)
		}
		helper1 = 0
	}
}

//12. Подсчет слов в строке
func WordInString() {
	
	fmt.Println("Введите предложение: ")
	inputPred, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// приводим строку к нижнему регистру и разбиваем на слова
	words := strings.Fields(strings.ToLower(inputPred))

	// используем map для хранения слов и их количества
	wordCount := make(map[string]int)
   
	for _, word := range words {
		wordCount[word]++
	}
	// количество уникальных слов
	fmt.Println("Количество уникальных слов: ", len(wordCount))
}

// 13. Игра "Жизнь"

const (
	rows = 10
	cols = 10
)

func displayGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("● ")
			} else {
				fmt.Print("○ ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func countAliveNeighbors(grid [][]int, x, y int) int {
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	count := 0

	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]
		if newX >= 0 && newX < rows && newY >= 0 && newY < cols {
			count += grid[newX][newY]
		}
	}

	return count
}

func updateGrid(grid [][]int) [][]int {
	newGrid := make([][]int, rows)
	for i := range newGrid {
		newGrid[i] = make([]int, cols)
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			aliveNeighbors := countAliveNeighbors(grid, x, y)

			if grid[x][y] == 1 {
				if aliveNeighbors == 2 || aliveNeighbors == 3 {
					newGrid[x][y] = 1
				} else {
					newGrid[x][y] = 0
				}
			} else {
				if aliveNeighbors == 3 {
					newGrid[x][y] = 1
				} else {
					newGrid[x][y] = 0
				}
			}
		}
	}

	return newGrid
}

func Task_13() {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	for {
		displayGrid(grid)
		grid = updateGrid(grid)
		time.Sleep(1 * time.Second)
	}
}

// 14. Цифровой корень числа
func DigitalSquare(number int) {
	var helper int
	helper = 0
	if number/10 >= 1 {
		for number > 0 {
			helper = number%10 + helper
			number = number/10
		}
		DigitalSquare(helper)
	} else {
		fmt.Println("Цифровой корень числа: ", number)
	}
}

// 15. Римские цифры
func RimDigits() {

	var num int
	fmt.Println("Введите число: ")
	fmt.Scan(&num)
	// массив римских чисел и соответствующих значений
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
   
	result := ""
	for i := 0; num > 0; i++ {
		// Пока текущее число больше или равно значению из массива
		for num >= vals[i] {
			result += romans[i] // добавляем римское число к результату
			num -= vals[i]      // уменьшаем число
		}
	}
	fmt.Print("Результат:", result)
}

func main() {
	//simpleNumber2()
	//Bubble()
	//Tablica()

	/*var number int
	fmt.Print("Введите количество чисел Фибоначчи: ")
	fmt.Scan(&number)
	fmt.Println("До числа: ", number)
	//fmt.Print("0", " ")
	fibonachi2(0, 1, number)*/

	//ReversDigits()
	//Pascal()
	//palindrome()

	//var numbers = [...]int{1,2,3,4,5,6,7,23,4,23,43223,0,2}
	//MaxMinMassive(numbers[:])

	//FindDigital()

	//WordInString()

	//Task_13()

	/*var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	DigitalSquare(number)*/

	//RimDigits()

}

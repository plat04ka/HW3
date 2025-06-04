package main

import "fmt"

func main() {
	//Task 1
	fmt.Println("-------------------------------------------")
	fmt.Println("Makes the function that receives *int and the updates it w/o return↓")
	x := 10
	a := &x
	fmt.Println("Recieved: ", x)
	updateNumber(a)
	fmt.Println("Updated ", x)

	//Task2
	fmt.Println("-------------------------------------------")
	fmt.Println("make fibonacci function without LOOP↓")

	nFib := 7
	n := nFib - 1
	fmt.Println(nFib, " number fibonacci -> ", fibonacсi(n))

	//Bonus Task
	fmt.Println("-------------------------------------------")
	fmt.Println("reverse integer↓")
	integer := 123
	fmt.Println("integer:", integer)
	K(integer)
	fmt.Println("N numbers in integer:", D)
	tenpow(D)
	mirror(integer, Vpow, D)
	fmt.Println("Reverse integer:", Mirror)

	//Task 3
	fmt.Println("-------------------------------------------")
	fmt.Println("make 3 function that call one after one ↓")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()
	obsos := otsos1(7)
	fmt.Println(obsos)
}

// updateNumber функция которая получает обращение к числу и обновляет его значение
func updateNumber(igor *int) {
	*igor = 11
}

// fibonaachi функция которая вычисляет число фибоначи под заданным номером, используя рекурсию
func fibonacсi(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacсi(n-1) + fibonacсi(n-2)
}

// otsos1 вторая функция, к которой обращается первая функция main, а сама обращается к третьей функции
func otsos1(otsosal1 int) int {
	fmt.Println(otsosal1)
	otsos2(otsosal1 - 1)
	return 10
}

// otsos2 третья функция к которой обращаются, и та паникует
func otsos2(otsosal2 int) {
	fmt.Println(otsosal2)

	panic("panica")
}

// D переменная которая помогает считать функцию K
var D = 0

// K Функция которая считает количество цифр в числе
func K(k int) int {
	if k == 0 {
		return D
	}
	D++
	K(k / 10)
	return D
}

// V и Vpow глобальные переменные, которые помогают считать функцию tenpow
var V = 10
var Vpow = 1

// Функция которая создает 10 в степени D-1
func tenpow(t int) int {
	if t == 1 {
		return Vpow
	}
	Vpow = Vpow * V
	tenpow(t - 1)
	return Vpow
}

// Mirror - Глобальная переменная которая помогает считать функцию mirror
var Mirror = 0

// Функция которая создает перевернутое число
func mirror(m int, n int, p int) {
	if p == 0 {
		return
	}
	Mirror += (m % 10) * n
	n /= 10
	m /= 10
	mirror(m, n, p-1)
}

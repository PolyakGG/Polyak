package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Ввод слов,включая пробел(по дефолту считывает до пробела)
	text, _ := reader.ReadString('\n')  // удаляем перенос на новую строку
	rs := []rune(text)                  // Превращаем string в rune (т.к. в рунах 1-русская буква 1 байт(2байта без превращения))
	var lastindex int                   // Определяем последни(й/ю) букву/символ в строке, т.к. по заданию в конце предложения должна быть точка.
	for i := range rs {
		lastindex = i
	}
	if rs[0] >= 'А' && rs[0] <= 'Я' && rs[lastindex] == '.' { //Задаем условие, что первая буква-заглавная,и последний индекс равен точке.
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}

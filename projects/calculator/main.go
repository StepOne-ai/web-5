package main

import (
	"fmt"
)

func calculator(first <-chan int, second <-chan int,  stop <-chan struct{}) <-chan int {
	res := make(chan int)
	go func() {
		for {
			select {
			case num := <-first:
				res <- num*num
			case num := <-second:
				res <- num*3
			case <-stop:
				fmt.Println("stop")
				close(res)
				return
			}
		}
	}()
	return res
}

func main() {
	stop := make(chan struct{})

	first := make(chan int)

	second := make(chan int)
	
	go func() {
		first <- 3 // Ввод в нужный канал числа
		first <- 4 // Ввод в нужный канал числ
		second <- 100 // Ввод в нужный канал числ
		stop <- struct{}{} // Сигнал остановки
		return
	}()

	res := calculator(first, second, stop)
	for num := range res {
		fmt.Println(num)
	}
	fmt.Println(<-res) // Вывод результата
}
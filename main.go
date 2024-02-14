package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Создание канала для сигналов
	sigs := make(chan os.Signal, 1)

	// Обработка сигнала SIGINT
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		for {
			select {
			case <-sigs:
				fmt.Println("выхожу из программы")
				os.Exit(0)
			default:
				// Вывод квадратов натуральных чисел в бесконечном цикле
				for i := 1; i <= 10; i++ {
					fmt.Printf("%d\n", i*i)

				}
			}
		}
	}()
	// Бесконечный цикл для ожидания сигнала
	select {}
}

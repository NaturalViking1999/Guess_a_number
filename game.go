package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	target := rand.Intn(100) + 1

	fmt.Println("Я выбираю число от 1 до 100")
	fmt.Println("Число выбрано")
	// fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	success := false

	for i := 0; i < 10; i++ {
		fmt.Println("У вас осталось попыток: ", 10-i)

		fmt.Println("Напишите ваше число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Ваше число больше, чем загаданное")
		} else if guess < target {
			fmt.Println("Ваше число меньше, чем загаданное")
		} else if guess == target {
			success = true
			fmt.Println("Поздравляем, вы угадали число!")
			break
		}
	}

	if !success {
		fmt.Println("Вы не смогли отгадать число и проиграли, загаданное число: ", target)
	}
}

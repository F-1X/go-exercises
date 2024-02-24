package main

import (
	"bufio"
	"fmt"
	"hashes/mapcache"
	"log"
	"os"
	"strconv"
	"time"
)

func enterManually() ([]int, []int) {
	fmt.Println("Размер 1 массива")

	var n int
	fmt.Scan(&n)

	var array1 []int
	array1 = make([]int, n)

	fmt.Println("Размер 2 массива")
	fmt.Scan(&n)
	var array2 []int
	array2 = make([]int, n)
	fmt.Println(array2)

	fmt.Println("Данные 1 массива:")

	var input int

	for i := 0; i < len(array1); i++ {
		fmt.Scan(&input)
		array1[i] = input
	}
	fmt.Println(array1)
	fmt.Println("Данные 2 массива:")

	for i := 0; i < len(array2); i++ {
		fmt.Scan(&input)
		array2[i] = input
	}

	return array1, array2

}

func inputFromFile() (error, *[]int, *[]int) {
	var array1 []int
	var array2 []int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return err, nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := 0
	// 1 array
	/*=================================*/
	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	fmt.Println("число", n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()

		input, _ = strconv.Atoi(line)
		array1 = append(array1, input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
		return err, nil, nil
	}

	// 2 array
	/*=================================*/
	_ = scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	fmt.Println("число", m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		line := scanner.Text()

		input, _ := strconv.Atoi(line)
		array2 = append(array2, input)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
		return err, nil, nil
	}

	return nil, &array1, &array2
}

func transformToHashMap(array []int) {
	//hashmap := make(map[string][]int, len(array))

	for i := 0; i < len(array); i++ {

		fmt.Println(array[i])

	}

}

func main() {
	// для наглядности вынесем
	err, array1, array2 := inputFromFile()
	if err != nil {
		log.Fatal("ops")
	}

	fmt.Println(*array1)
	fmt.Println(*array2)

	// memcached()

}

// NO THREAD SAFETY
// НЕ ПОТОКОБЕЗОПАСНО
func memcached() {
	// задание 14.5.2, думаю можно прикрутить к итоговому заданию
	// и использовать как хранилище хеш карт, добавляется только временное хранилище
	// может пригодиться чтобы не лазить в бд и не высчитывать хеши повторно

	var a mapcache.Cache = mapcache.NewInMemoryCache(time.Second * 2)

	an := a.Get("1")
	fmt.Println(an)

	a.Set("1", "hello from 1")
	an = a.Get("1")
	fmt.Println(an)

	time.Sleep(time.Second * 3)
	a.Set("2", "hello from 2")

	time.Sleep(time.Second * 3)
	a.Set("3", "hello from 3")

	time.Sleep(time.Second * 3)

	fmt.Println(a.Get("1"))
	fmt.Println(a.Get("2"))
	fmt.Println(a.Get("3"))

	fmt.Println("видимо уже ничего нет с таким таймингом")

	select {}
}

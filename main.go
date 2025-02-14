package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
//Незабыть создать функцию конструктор и методы.
type Person struct{
	Name string
	Age int
	geteroFriend bool
	balance struct{
		Credit float32
		Point int
}
}*/

func main() {
	var mulA int = 4
	var mulB int = 6
	var mulN int = 2
	var someString string = "one two three four five six seven eight nine ten"
	var price []int
	var point []int
	var list []string

	fmt.Println(sum(mulA, mulB), "\t 1. вывести сумму чисел")                               // 1. вывести сумму чисел
	fmt.Println(itr(mulA, mulB), "\t 2. увеличить число на выбаное кол-во итераций")        // 2. увеличить число на выбаное кол-во итераций
	fmt.Println(itoa(mulA), "\t 3. привести число к строке")                                // 3. привести число к строке
	fmt.Println(fillSlice(price, mulA), "\t 4. заполнить слайс на выбаное кол-во символов") // 4.заполнить слайс на выбаное кол-во символов
	price = fillSlice(price, mulA)
	point = fillSlice(point, mulA)
	fmt.Println(addSlice(someString), "\t 5. залить текст в слайс c разделением по ` `") // 5. залить текст в слайс c разделением по " "
	list = addSlice(someString)
	list1 := list
	list = slicecenterspace(list, mulA, mulN) // 6. вырезать центр слайса c явным указанием начала и окончания элементов
	fmt.Println(list, "\t 6. вырезать центр слайса c явным указанием начала и окончания элементов")
	list = sliceSpaceEnd(list, mulA) // 7. обрезать слайс с конца до индекса
	fmt.Println(list, "\t 7. обрезать слайс с конца до индекса")
	list1 = sliceSpaceBegin(list1, mulN) // 8. обрезать слайс с начала до индекса
	fmt.Println(list1, "\t 8. обрезать слайс с начала до индекса")
	list1 = sliceSpaceBeginEnd(list1, mulN, mulN) // 9. обрезать слайс с начала и конца по индексам
	fmt.Println(list1, "\t 9. обрезать слайс с начала и конца по индексам")
	fmt.Println(findIndex(price, mulA), "\t 10. поиск индекса числа в слайсе")  // 10. поиск индекса числа в слайсе
	fmt.Println(sumslice(price, point), "\t 11. добавить слайс в слайс")        // 11. добавить слайс в слайс
	fmt.Println(summas(price, point), "\t 12. сложить по элементно два слайса") // 12. сложить по элементно два слайса
	fmt.Println(revers(price), "\t 13. переворот слайса")                       // 13. переворот слайса
	fmt.Println(setmap(list, point), " \t 14. создание мапы из двух слайсов")   // 14. создание мапы из двух слайсов //Уникальность ключей:Ключи в мапе должны быть уникальными. Если в срезе keys есть дубликаты, последнее значение перезапишет предыдущее.

}

// 1. вывести сумму чисел
func sum(a, b int) int {
	return a + b
}

// 2. увеличить число на выбаное кол-во итераций
func itr(a, b int) int {
	for i := a; i <= b; i++ {
		a = i
	}
	return a
}

// 3. привести число к строке
func itoa(a int) string {
	x := strconv.Itoa(a)
	return x
}

// 4.заполнить слайс на выбаное кол-во символов
func fillSlice(slice []int, a int) []int {
	for i := 1; i <= a; i++ {
		slice = append(slice, i)
	}
	return slice
}

// 5. залить текст в слайс c разделением по " "
func addSlice(str string) []string {
	slice := strings.Split(str, " ")
	return slice
}

// 6. вырезать центр слайса c явным указанием начала и окончания элементов
func slicecenterspace(slice []string, begin, end int) []string {
	slicecopy := make([]string, 0, len(slice))
	slicecopy = slice[len(slice)-end:]
	slice = append(slice[:begin], slicecopy...)
	return slice
}

// 7. обрезать слайс с конца до индекса
func sliceSpaceEnd(slice []string, i int) []string {
	return slice[:i]
}

// 8. обрезать слайс с начала до индекса
func sliceSpaceBegin(slice []string, i int) []string {
	return slice[i:]
}

// 9. обрезать слайс с начала и конца по индексам
func sliceSpaceBeginEnd(slice []string, i, i2 int) []string {
	return slice[i:i2]
}

// 10. поиск индекса числа в слайсе
func findIndex(slice []int, a int) int {
	for i, v := range slice {
		if v == a {
			return i
		}
	}
	return -1
} //Println непринимает 2 параметра. Тогда надо было делать так: inxex, value := findIndex(slice []int, a int)(int,int)

// 11. добавить слайс в слайс
func sumslice(slice, slice2 []int) []int {
	slice = append(slice, slice2...)
	return slice
}

// 12. сложить по элементно два слайса
func summas(slice, slice2 []int) []int {
	res := make([]int, len(slice)) //все создано верно. Иначе была ошибка: panic: runtime error: index out of range [0] with length 0
	if len(slice) == len(slice2) {
		for i := range slice {
			res[i] = slice[i] + slice2[i]
		}
		return res
	}
	return nil
}

// 13. переворот слайса
func revers(slice []int) []int {
	cntindex := len(slice)
	for i := 0; i < len(slice)/2; i++ {
		// Меняем местами элементы, пока не дойдем до середины
		slice[i], slice[cntindex-1-i] = slice[cntindex-1-i], slice[i]
	}
	return slice
}

// 14. создание мапы из двух слайсов
func setmap(keys []string, values []int) map[string]int {
	if len(keys) == len(values) {
		result := make(map[string]int, len(keys))
		for i := 0; i < len(keys); i++ {
			result[keys[i]] = values[i]
		}
		return result
	}
	return nil
}

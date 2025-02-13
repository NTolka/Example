package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Xepyz
// var a int = 4
// var b int = 6
// var n int = 2
// var z string = " "
// var c float64 //no use
// var t string
// var e []int
// var f []int
// var g []string

func main() {
	t = "one two three four five six seven eight nine ten"
	fmt.Println(t)
	fmt.Println(sum(a, b), "//1. сложить два числа `a` и `b`")                     //1. сложить два числа `a` и `b`
	fmt.Println(itr(a, b), "//2. увеличить число `a` на `b` итераций с шагом `n`") //2. увеличить число `a` на `b` итераций с шагом `n`
	fmt.Println(itoa(a), "//3. превести число `a` к строке")                       //3. превести число `a` к строке
	fmt.Println(fill(e, a), "//4. заполнить срез `e` на `a` символов")             //4. заполнить срез `e` на `a` символов
	e = fill(e, a)
	f = e
	fmt.Println(addslice(t, z), "//5. залить текст `t` в слайс по разделителю `z`") //5. залить текст `t` в слайс по разделителю `z` ("\n"," ",".")
	g = slicecenterspace(g, n)                                                      //6. вырезать кусок слайса
	fmt.Println(g, "//6. вырезать кусок слайса")
	g = slicespace(g, a) //7. обрезать слайс.
	fmt.Println(g, "//7. обрезать слайс.")
	fmt.Println(find(e, a), "//5. индекс числа `a` в срезе `e`")             //8. индекс числа `a` в срезе `e`
	fmt.Println(sumslice(e, f), "//6. сложить срезы `e` и `f` в один слайс") //9. сложить срезы `e` и `f` в один слайс
	fmt.Println(summas(e, f), "//7. сложить по элементно срезы `e` и `f`")   //10. сложить по элементно срезы `e` и `f`
	fmt.Println(revers(e), "//8. переворот слайса `e`")                      //11. переворот слайса `e`
	fmt.Println(setmap(g, e), " //9. сложение срезов `e` и `g` в мапу")      //12. сложение срезов `e` и `g` в мапу //Уникальность ключей:Ключи в мапе должны быть уникальными. Если в срезе keys есть дубликаты, последнее значение перезапишет предыдущее.
}

// 1. сложить два числа `a` и `b`
func sum(...int) int {
	return a + b
}

// 2. увеличить число `a` на `b` итераций с шагом `n`
func itr(a, b int) int {
	for i := a; i <= b; i++ {
		a = i + n
	}
	return a
}

// 3. превести число `a` к строке
func itoa(a int) string {
	x := strconv.Itoa(a)
	return x
}

// 4. заполнить срез `e` на `a` символов
func fill(e []int, a int) []int {
	for i := 1; i <= a; i++ {
		e = append(e, i)
	}
	return e
}

// 5. залить текст `t` в слайс по разделителю `z` ("\n"," ",".")
func addslice(t, z string) []string {
	g = strings.Split(t, z)
	return g
}

// 6. вырезать кусок слайса
func slicecenterspace(g []string, n int) []string {
	j := make([]string, 0, 10)
	j = g[len(g)-n:] // g[8:]
	g = append(g[:5], j...)
	return g
}

// 7. обрезать слайс.
func slicespace(g []string, a int) []string {
	return g[:a]
}

// 8. поиск числа `a` в срезе `e` с указаием индекса
func find(e []int, a int) int {
	for i, v := range e {
		if v == a {
			return i
		}
	}
	return 0
}

// 9. сложить срезы `e` и `f` в один слайс
func sumslice(e, f []int) []int {
	e = append(e, f...)
	return e
}

// 10. сложить по элементно срезы `e` и `f`
func summas(e, f []int) []int {
	res := make([]int, len(e))
	if len(e) == len(f) {
		for i := range e {
			res[i] = e[i] + f[i]
		}
		return res
	}
	return nil
}

// 11. переворот слайса `e`
func revers(e []int) []int {
	x := len(e)
	for i := 0; i < len(e)/2; i++ {
		e[i], e[x-1-i] = e[x-1-i], e[i]
	}
	return e
	// 	// Определяем начальный и конечный индексы
	// 	left, right := 0, len(e)-1

	// // Меняем местами элементы, пока не дойдем до середины
	//
	//	for left < right {
	//		e[left], e[right] = e[right], e[left]
	//		left++
	//		right--
	//	}
	//
	// return e
}

// 12. сложение срезов `e` и `g` в мапу
func setmap(g []string, e []int) map[string]int {
	if len(e) == len(f) {
		keys := g
		values := e
		result := make(map[string]int)
		for i := 0; i < len(keys); i++ {
			result[keys[i]] = values[i]
		}
		return result
	}
	return nil
}

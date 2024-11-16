package main

import (
	"errors"
	"fmt"
)

func reverse(sl []int) []int {
	a := 1
	errors.As(errors.New("123"), &a)
	ans := make([]int, len(sl))
	for i, v := range sl {
		// выполнится после return, return вернет пустой слайс, но после уже, в него подставятся значения
		defer func(ind, val int) { ans[ind] = val }(len(sl)-1-i, v)
	}
	fmt.Println(ans)
	return ans
}

func empty() []int {
	ans := []int{1, 2, 3, 4, 5}
	fmt.Printf("EMTPY %p %v\n", ans, ans)
	defer func() {
		fmt.Printf("EMTPY DEFER  %p %v\n", ans, ans)
		ans = nil
	}() // меняет указатель! вернется нормальный
	defer func() { ans[0] = 0 }()
	// мы можем только занулить все элементы
	return ans
}
func empty2() (ans []int) {
	ans = []int{1, 2, 3, 4, 5}
	fmt.Printf("EMTPY 2  %p %v\n", ans, ans)
	defer func() {
		fmt.Printf("EMTPY DEFER 2  %p %v\n", ans, ans)
		ans = nil
	}() // меняет указатель! вернется нормальный
	defer func() { ans[0] = 0 }()
	// мы можем только занулить все элементы
	return ans
}

func sl(a []int) {
	a[0] = 123
	a = nil
}

func main() {
	fmt.Println(reverse([]int{1, 2, 3}))
	fmt.Println(empty())
	v := empty2()
	fmt.Printf("%v %T\n", v, v)
	fmt.Println(v == nil)
	mySlice := []int{1, 2, 3}
	sl(mySlice)
	fmt.Println(mySlice)
}

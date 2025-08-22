package main

import "fmt"

type Stringer interface {
	comparable
	String() string
}

func StringifyWithoutZero[T Stringer](s []T) (ret []string) {
	var zero T
	for _, v := range s {
		if v == zero {
			continue
		}
		ret = append(ret, v.String())
	}
	return ret
}

type MyString string

func (s MyString) String() string {
	return string(s)
}

func main() {
	sl := StringifyWithoutZero([]MyString{"I", "", "love", "", "golang"}) // 输出：[I love golang]
	fmt.Println(sl)
}

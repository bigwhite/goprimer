package main

import "fmt"

/*
func Strcat(s string, others ...string) string {
	var r = s
	for _, o := range others {
		r += o
	}
	return r
}
*/

func Strcat(s string, others ...string) string {
	var r = s
	for _, o := range others {
		r = fmt.Sprintf("%s%s", r, o)
	}
	return r
}

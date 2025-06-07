package main

import (
	"practice/channels"
	"practice/concurrency"
	"practice/errors"
	"practice/generics"
	"practice/lib"
	"practice/mypointers"
	"practice/myreader"
	"practice/practice"
)

func main() {
	example := myreader.ReadExample()
	switch example {
	case 1:
		mypointers.Ex1()
	case 2:
		mypointers.Ex2()
	case 3:
		mypointers.Ex3()
	case 4:
		mypointers.Ex4()
	case 5:
		generics.Ex1()
	case 6:
		lib.Ex1()
	case 7:
		lib.Ex2()
	case 8:
		lib.Ex3()
	case 9:
		lib.Ex4()
	case 10:
		practice.Ex1()
	case 11:
		practice.Ex2()
	case 12:
		practice.Ex3()
	case 13:
		concurrency.Ex1()
	case 14:
		concurrency.Ex2()
	case 15:
		concurrency.Ex3()
	case 16:
		concurrency.Ex4()
	case 17:
		practice.Ex5()
	case 18:
		practice.Ex6()
	case 19:
		channels.Ex1()
	case 20:
		channels.Ex2()
	case 21:
		channels.Ex3()
	case 22:
		channels.Ex4()
	case 23:
		channels.Ex5()
	case 24:
		channels.Ex6()
	case 25:
		channels.Ex7()
	case 26:
		channels.Ex8()
	case 27:
		channels.Ex9()
	case 28:
		channels.Ex10()
	case 29:
		channels.FanIn()
	case 30:
		channels.FanOut()
	case 31:
		channels.Ex11()
	case 32:
		channels.Ex12()
	case 33:
		channels.Ex13()
		channels.Ex14()
		channels.Ex15()
		channels.Ex16()
		channels.Ex17()
		channels.Ex18()
	case 34:
		errors.Ex1("p") // fmt.println
	case 35:
		errors.Ex1("l") //log.Prntln
	case 36:
		errors.Ex1("f") // log.Fatalln
	case 37:
		errors.Ex1("i") // panic
	case 38:
		errors.Ex2()
	case 39:
		errors.Ex3()
	case 40:
		errors.Ex4()
	case 41:
		errors.Ex5()
	}
}

package main

import (
	"fmt"
	"piscine"
)

func main() {
	// fmt.Println(piscine.CheckNumber("Hello"))
	// fmt.Println(piscine.CheckNumber("Hello1")) //checkNumber
	// fmt.Println(piscine.CountAlpha("Hello world"))
	// fmt.Println(piscine.CountAlpha("H e l l o"))
	// fmt.Println(piscine.CountAlpha("H1e2l3l4o"))
	// fmt.Println(piscine.CountChar("Hello World", 'l'))
	// fmt.Println(piscine.CountChar("5  balloons", 5))
	// fmt.Println(piscine.CountChar("   ", ' '))
	// fmt.Println(piscine.CountChar("The 7 deadly sins", '7'))
	// fmt.Print(piscine.PrintIf("abcdefz"))
	// fmt.Print(piscine.PrintIf("abc"))
	// fmt.Print(piscine.PrintIf(""))
	// fmt.Print(piscine.PrintIf("14"))
	// fmt.Print(piscine.PrintIfNot("abcdefz"))
	// fmt.Print(piscine.PrintIfNot("abc"))
	// fmt.Print(piscine.PrintIfNot(""))
	// fmt.Print(piscine.PrintIfNot("14"))
	// fmt.Println(piscine.RectPerimeter(10, 2))
	// fmt.Println(piscine.RectPerimeter(434343, 898989))
	// fmt.Println(piscine.RectPerimeter(10, -2))
	// fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	// fmt.Println(piscine.RetainFirstHalf("A"))
	// fmt.Println(piscine.RetainFirstHalf(""))
	// fmt.Println(piscine.RetainFirstHalf("Hello World"))
	fmt.Println(piscine.CamelToSnakeCase("HelloWorld"))
	fmt.Println(piscine.CamelToSnakeCase("helloWorld"))
	fmt.Println(piscine.CamelToSnakeCase("camelCase"))
	fmt.Println(piscine.CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(piscine.CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(piscine.CamelToSnakeCase("hey2"))
}

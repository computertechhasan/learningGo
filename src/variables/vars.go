package main

import "fmt"

func main() {
	//go has string, bool, int, int8-64, uint (unsigned)8-64, float32-64, complex(64-128)

	var first string = "Hasan";
	const isAwesome bool = true;
	var age int = 22;

	//in a function  this is okay syntax, not outside of one // inAFunction := true;

	newFirst, last := "hasan", "intisar";

	fmt.Println(first + " is learning go! And bools are ", isAwesome, " and today is ", age);
	fmt.Println(newFirst, last);
}

https://www.youtube.com/watch?v=SqrbIlUwR0U&t=200s 24 mins
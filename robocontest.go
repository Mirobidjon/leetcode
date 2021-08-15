package main

import "fmt"

func Str() {
	var (
		n int
		s string
	)

	fmt.Scan(&n)
	fmt.Scanln(&s)

	a := Check(s)
	if a > 6-n {
		fmt.Println(a)
	} else {
		fmt.Println(6 - n)
	}

}

func Check(s string) int {
	var (
		z bool
		c int
	)
	for _, v := range s {
		if '0' <= v && v <= '9' {
			z = true
			break
		}
	}

	if !z {
		c++
	}
	z = false

	for _, v := range s {
		if 'a' <= v && v <= 'z' {
			z = true
			break
		}
	}

	if !z {
		c++
	}
	z = false

	for _, v := range s {
		if 'A' <= v && v <= 'Z' {
			z = true
			break
		}
	}

	if !z {
		c++
	}
	z = false

	for _, v := range s {
		if v == '!' || v == '@' || v == '#' || v == '$' || v == '%' || v == '^' || v == '&' || v == '*' || v == '(' || v == ')' || v == '-' || v == '+' {
			z = true
			break
		}
	}

	if !z {
		c++
	}
	return c
}

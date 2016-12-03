package functions

import "fmt"

func Greet(name, sname string) string {
	return fmt.Sprint(name, " ", sname)
}

func Namedgreet(name, sname string) (s string) {
	s = fmt.Sprint(name, " ", sname)
	return s
}

func Doublereturn(name, sname string) (string, string) {
	return fmt.Sprint(name, " ", sname), fmt.Sprint(sname, " ", name)
}

package Assignment1

import "fmt"

func (D Database) auth(n string, l string, p string) bool {
	var u = User{
		n,
		l,
		p,
	}
	for i := 0; i < len(D.users); i++ {
		if u == D.users[i] {
			fmt.Println("Welcome!")
			return true
		} else {
			fmt.Println("No such user!")
			break
		}
	}
	return false
}

package Assignment1

func (D Database) reg(n string, l string, p string) {
	var u = User{n, l, p}
	D.users = append(D.users, u)
}

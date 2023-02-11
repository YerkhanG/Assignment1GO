package Assignment1

func (D Database) addi(n string, p int, r int) {
	var item = Item{
		n,
		p,
		r,
	}
	D.items = append(D.items, item)
}

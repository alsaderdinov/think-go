package person

type Employee struct {
	age uint
}

type Customer struct {
	age uint
}

func (e *Employee) Age() uint {
	return e.age
}

func (c *Customer) Age() uint {
	return c.age
}

type Ager interface {
	Age() uint
}

func Eldest(a ...Ager) uint {
	var eldAge uint

	for _, v := range a {
		if v.Age() > eldAge {
			eldAge = v.Age()
		}
	}
	return eldAge
}

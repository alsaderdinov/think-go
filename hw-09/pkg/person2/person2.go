package person2

type Employee struct {
	age uint
}

type Customer struct {
	age uint
}

func Eldest(a ...any) any {
	var eldAge uint
	var eldest any

	for _, v := range a {
		switch val := v.(type) {
		case Employee:
			if val.age > eldAge {
				eldAge, eldest = val.age, val
			}
		case Customer:
			if val.age > eldAge {
				eldAge, eldest = val.age, val
			}
		default:
			continue
		}
	}
	return eldest
}

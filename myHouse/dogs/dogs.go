package dogs

type Dogs struct {
	Name string
	Age  int
}

func CreateObjects() []Dogs {
	return []Dogs{
		{Name: "Рекс", Age: 11},
		{Name: "Рой", Age: 1},
		{Name: "Джек", Age: 2},
		{Name: "Тим", Age: 1},
		{Name: "Бим", Age: 1},
	}
}

package cats

type Cats struct {
	Name string
	Age  int
}

func CreateObjects() []Cats {
	return []Cats{
		{Name: "Мурка", Age: 11},
		{Name: "Милка", Age: 1},
		{Name: "Машка", Age: 2},
		{Name: "Кашка", Age: 2},
		{Name: "Лизка", Age: 1},
	}
}

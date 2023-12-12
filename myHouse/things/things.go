package things

type Things struct {
	Name  string
	Size  float64
	Count int
}

func CreateObjects() []Things {
	return []Things{
		{Name: "Стул", Size: 0.8, Count: 4},
		{Name: "Стол", Size: 5.333, Count: 1},
		{Name: "Ковер", Size: 6, Count: 1},
		{Name: "Кровать", Size: 5.6, Count: 5},
		{Name: "Тумба", Size: 1.5, Count: 1},
	}
}

package children

type Children struct {
	Name string
	Age  int
}

func CreateObjects() []Children {
	return []Children{
		{Name: "Артем", Age: 23},
		{Name: "Саша", Age: 17},
		{Name: "Петя", Age: 3},
		{Name: "Дима", Age: 2},
		{Name: "Даша", Age: 2},
	}
}

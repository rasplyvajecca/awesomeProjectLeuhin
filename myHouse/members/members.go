package members

type Members struct {
	Member  string
	Name    string
	Surname string
}

func CreateObjects() []Members {
	return []Members{
		{Member: "Мама", Name: "Мария", Surname: "Леухина"},
		{Member: "Папа", Name: "Олег", Surname: "Леухин"},
		{Member: "Бабушка", Name: "Ирина", Surname: "Леухина"},
		{Member: "Дядя", Name: "Максим", Surname: "Леухин"},
		{Member: "Тетя", Name: "Зинаида", Surname: "Леухина"},
	}
}

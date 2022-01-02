package models

type Person struct {
	Name     string
	Birthday string
	City     string
}

type GetMethods interface {
	GetMethod(param string) string
}

func (person *Person) GetMethod(param string) string {
	var result string
	switch param {
	case "Name":
		result = person.Name
	case "Birthday":
		result = person.Birthday
	case "City":
		result = person.City
	default:
		result = "No result"
	}
	return result
}

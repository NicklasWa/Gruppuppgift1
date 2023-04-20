package data

type Employee struct {
	Id   int
	Age  int
	City string
	Namn string
}

func (emp Employee) IsCool() bool {
	if emp.Namn == "Stefan" {
		return true
	}
	return false
}

func IsCool(emp Employee) bool {
	if emp.Namn == "Stefan" {
		return true
	}
	return false
}

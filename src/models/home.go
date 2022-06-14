package models

func (home *Home) Reset() {
	home.Value1 = "Default"
	home.Value2 = 0
	home.Value3 = false
}

type Home struct {
	Value1 string
	Value2 int
	Value3 bool
}

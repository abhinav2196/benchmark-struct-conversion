package main


type Input struct {
	name string `json:"name"`
	f1 string	`json:"f1"`
	f2 string	`json:"f2"`
	s1 S1		`json:"s1"`
	s2 S2		`json:"s2"`
}

type S1 struct{
	f1 string `json:"f1"`
}

type S2 struct {
	s1 S1 	`json:"s1"`
	f2 string	`json:"f2"`
}



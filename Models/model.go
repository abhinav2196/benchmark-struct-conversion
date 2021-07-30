package Models


type Response1 struct {
	Page   int `copier:"must",json:"page"`
	Fruits string `copier:"must",json:"fruits"`
	Z *Response2 `copier:"must",json:"z"`
}



type Response2 struct {
	X   int `copier:"must",json:"x"`
	Y string `copier:"must",json:"y"`
}
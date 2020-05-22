package viewmodel


type Register struct{
	Title string
	Active string
	Email string
	Password string
	Fullname string
}


func NewRegister() Register {
	result := Register{
		Active: "register",
		Title: "register",
	}
	return result
}





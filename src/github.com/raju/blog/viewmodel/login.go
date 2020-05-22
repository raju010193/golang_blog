package viewmodel

type Login struct {
	Title    string
	Active   string
	Email    string
	Password string
}

func NewLogin() Login {
	result := Login{
		Active: "login",
		Title:  "Bloger login",
	}
	return result
}

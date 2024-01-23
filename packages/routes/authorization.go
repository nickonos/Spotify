package routes

type Login struct {
	Code string
}

func (id Login) Subject() string {
	return "login"
}

type LoginResponse struct {
	JWT string
}

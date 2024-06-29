package auth

type AuthUsecase interface {
	Login(name, password string) (string, error)
	SignUp(name, password string) (string, error)
}

type AuthRepository interface {
	CheckUser(name, password string) (int, error)
	CreateUser(name, password string) (int, error)
}

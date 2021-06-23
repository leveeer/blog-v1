package service

type IUserAuthService interface {
	GetLoginCode(username string) error
}

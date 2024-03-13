package controllers

type LoginService interface {
	LoginUser(username string) bool
}
type loginInformation struct {
	username string
}

func StaticLoginService() LoginService {
	return &loginInformation{}
}

func (info *loginInformation) LoginUser(username string) bool {
	return info.username == username
}

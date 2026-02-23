package user

import "syloria-demo/domain"

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (svr *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svr.usrRepo.Create(user)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}
	return usr, nil
}
func (svr *service) Find(email string, pass string) (*domain.User, error) {
	usr, err := svr.usrRepo.Find(email, pass)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}
	return usr, nil
}

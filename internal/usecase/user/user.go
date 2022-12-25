package user

import "github.com/wwwshwww/graphql-sandbox/internal/domain/user"

type UserUsecase interface {
	Get(uint64) (user.User, GetError)
	Save(uint64, user.UserPreferences) SaveError
}

func New(ur user.Repository) userUsecase {
	return userUsecase{
		userRepository: ur,
	}
}

type userUsecase struct {
	userRepository user.Repository
}

func (uc userUsecase) Get(id uint) (user.User, GetError) {
	u, gerr := uc.userRepository.Get(id)
	if gerr != nil {
		switch gerr := gerr.(type) {
		case user.RepositoryGetErrorInfrastructural:
			return nil, GetErrorInfrastructural{Err: gerr.Err}
		default:
			panic("unexpected error")
		}
	}
	return u, nil
}

func (uc userUsecase) Save(id uint, pr user.UserPreferences) SaveError {
	u, gerr := uc.userRepository.Get(id)
	if gerr != nil {
		switch gerr := gerr.(type) {
		case user.RepositoryGetErrorInfrastructural:
			return SaveErrorInfrastructural{Err: gerr.Err}
		default:
			panic("unexpected error")
		}
	}

	u.Overwrite(pr)

	serr := uc.userRepository.Save(u)
	if serr != nil {
		switch serr := serr.(type) {
		case user.RepositorySaveErrorInfrastructural:
			return SaveErrorInfrastructural{Err: serr.Err}
		default:
			panic("unexpected error")
		}
	}

	return nil
}

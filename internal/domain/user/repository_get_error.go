package user

type RepositoryGetError interface {
	error
	isRepositoryGetError()
}

// RepositoryGetErrorの想定されたパターン
var (
	_ RepositoryGetError = RepositoryGetErrorInfrastructural{}
)

type RepositoryGetErrorInfrastructural struct {
	Err error
}

func (RepositoryGetErrorInfrastructural) isRepositoryGetError() {}
func (e RepositoryGetErrorInfrastructural) Error() string {
	return e.Err.Error()
}

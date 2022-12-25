package user

type RepositorySaveError interface {
	error
	isRepositorySaveError()
}

// RepositorySaveErrorの想定されたパターン
var (
	_ RepositorySaveError = RepositorySaveErrorInfrastructural{}
)

type RepositorySaveErrorInfrastructural struct {
	Err error
}

func (RepositorySaveErrorInfrastructural) isRepositorySaveError() {}
func (e RepositorySaveErrorInfrastructural) Error() string {
	return e.Err.Error()
}

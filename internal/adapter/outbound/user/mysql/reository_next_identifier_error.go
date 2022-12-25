package user_mysql

type repositoryNextIdentifierError struct {
	touched         bool
	infrastructural error
}

func (e *repositoryNextIdentifierError) Error() string {
	return e.infrastructural.Error()
}

func (e *repositoryNextIdentifierError) IsRepositoryNextIdentifierError() {}

func (e *repositoryNextIdentifierError) touch() {
	e.touched = true
}

func (e *repositoryNextIdentifierError) setInfrastructural(err error) {
	e.touch()
	e.infrastructural = err
}

func (e *repositoryNextIdentifierError) Infrastructural() error {
	return e.infrastructural
}

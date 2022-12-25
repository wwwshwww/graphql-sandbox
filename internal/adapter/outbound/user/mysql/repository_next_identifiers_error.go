package user_mysql

type repositoryNextIdentifiersError struct {
	touched         bool
	infrastructural error
}

func (e *repositoryNextIdentifiersError) Error() string {
	return e.infrastructural.Error()
}

func (e *repositoryNextIdentifiersError) IsRepositoryNextIdentifiersError() {}

func (e *repositoryNextIdentifiersError) touch() {
	e.touched = true
}

func (e *repositoryNextIdentifiersError) setInfrastructural(err error) {
	e.touch()
	e.infrastructural = err
}

func (e *repositoryNextIdentifiersError) Infrastructural() error {
	return e.infrastructural
}

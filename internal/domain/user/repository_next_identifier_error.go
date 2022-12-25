package user

type RepositoryNextIdentifierError interface {
	error
	IsRepositoryNextIdentifierError()

	Infrastructural() error
}

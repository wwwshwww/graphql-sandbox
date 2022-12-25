package user

type RepositoryNextIdentifiersError interface {
	error
	IsRepositoryNextIdentifiersError()

	Infrastructural() error
}

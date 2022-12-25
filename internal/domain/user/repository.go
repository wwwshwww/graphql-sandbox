package user

type Repository interface {
	Get(uint) (User, RepositoryGetError)
	Save(User) RepositorySaveError
	NextIdentifier() (uint, RepositoryNextIdentifierError)
	NextIdentifiers(uint) ([]uint, RepositoryNextIdentifiersError)
}

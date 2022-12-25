package user_mysql

import (
	"github.com/wwwshwww/graphql-sandbox/internal/domain/user"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return Repository{db: db}
}

func (r Repository) Get(id uint) (user.User, user.RepositoryGetError) {
	var rows []User

	if err := r.db.Model(&User{}).Where("id = ?", id).Find(&rows).Error; err != nil {
		return nil, user.RepositoryGetErrorInfrastructural{Err: err}
	}

	if len(rows) > 0 {
		return unmarshal(rows[0]), nil
	}

	return user.New(id, user.UserPreferences{}), nil
}

func (r Repository) Save(u user.User) user.RepositorySaveError {
	if err := r.db.Model(&User{}).Where("id = ?", u.ID()).Save(marshal(u)).Error; err != nil {
		return user.RepositorySaveErrorInfrastructural{Err: err}
	}
	return nil
}

func (r Repository) NextIdentifier() (uint, user.RepositoryNextIdentifierError) {
	rerr := &repositoryNextIdentifierError{}

	ids, err := r.NextIdentifiers(1)
	if err != nil {
		rerr.setInfrastructural(err.Infrastructural())
		return 0, rerr
	}

	if len(ids) != 1 {
		panic("NextIdentifiers didn't provide exactly 1 identifier")
	}

	return ids[0], nil
}

func (r Repository) NextIdentifiers(n uint) ([]uint, user.RepositoryNextIdentifiersError) {
	if n == 0 {
		return nil, nil
	}

	rerr := &repositoryNextIdentifiersError{}

	rows := make([]User, n)
	if err := r.db.Save(&rows).Error; err != nil {
		rerr.setInfrastructural(err)
		return nil, rerr
	}
	if err := r.db.Delete(rows).Error; err != nil {
		rerr.setInfrastructural(err)
		return nil, rerr
	}

	ids := make([]uint, n)
	for i, row := range rows {
		ids[i] = row.ID
	}

	return ids, nil
}

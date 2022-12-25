package user_mysql

import "github.com/wwwshwww/graphql-sandbox/internal/domain/user"

func unmarshal(row User) user.User {
	return user.Restore(
		row.ID,
		user.UserPreferences{
			Name:     row.Name,
			Password: row.Password,
		},
	)
}

func marshal(u user.User) User {
	return User{
		ID:       u.ID(),
		Name:     u.Name(),
		Password: u.Password(),
	}
}

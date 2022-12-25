package user

type User interface {
	ID() uint
	Name() string
	Password() string

	SetName(string)
	SetPassword(string)

	Overwrite(UserPreferences)
}

type user struct {
	id       uint
	name     string
	password string
}

// ユーザのすべての設定可能項目をまとめた構造体
type UserPreferences struct {
	Name     string
	Password string
}

func New(id uint, pr UserPreferences) User {
	return &user{
		id:       id,
		name:     pr.Name,
		password: pr.Password,
	}
}

func Restore(id uint, pr UserPreferences) User {
	return New(id, pr)
}

func (u *user) ID() uint {
	return u.id
}

func (u *user) Name() string {
	return u.name
}

func (u *user) SetName(name string) {
	u.name = name
}

func (u *user) Password() string {
	return u.password
}
func (u *user) SetPassword(password string) {
	u.password = password
}

func (u *user) Overwrite(pr UserPreferences) {
	u.SetName(pr.Name)
	u.SetPassword(pr.Password)
}

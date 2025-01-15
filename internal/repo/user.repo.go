package repo

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUser(id string) (string, string) {
	return id, "John Doe"
}

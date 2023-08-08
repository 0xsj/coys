package account

type Repository interface {
	CreateAccount()
	GetAccountById()
}

type RepositoryImpl struct{}

func NewRepository() {}

func CreateAccount() {}

func GetAccountById() {}

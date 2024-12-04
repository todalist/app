package user

type IUserRepo interface {

	// basic crud
	Get(uint) (*User, error)

	Save(*User) (*User, error)

	List(*UserQuerier) ([]*User, error)

	Delete(uint) (uint, error)

}
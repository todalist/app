package user

type IUserStore interface {

	// basic crud
	Get(uint) (*User, error)

	Save(*User) (*User, error)

	List(*UserQuerier) ([]*User, error)

	Delete(uint) (uint, error)

}
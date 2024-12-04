package userToda

type IUserTodaRepo interface {

	// basic crud
	Get(uint) (*UserToda, error)

	Save(*UserToda) (*UserToda, error)

	List(*UserTodaQuerier) ([]*UserToda, error)

	Delete(uint) (uint, error)

}
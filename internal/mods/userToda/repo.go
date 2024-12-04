package userToda

type IUserTodaRepo interface {

	// basic crud
	Get(uint) (*UserToda, error)

	First(*UserTodaQuerier) (*UserToda, error)

	Save(*UserToda) (*UserToda, error)

	List(*UserTodaQuerier) ([]*UserToda, error)

	Delete(uint) (uint, error)

	DeleteByTodaId(uint) error

}
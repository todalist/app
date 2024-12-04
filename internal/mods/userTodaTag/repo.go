package userTodaTag

type IUserTodaTagRepo interface {

	// basic crud
	Get(uint) (*UserTodaTag, error)

	Save(*UserTodaTag) (*UserTodaTag, error)

	List(*UserTodaTagQuerier) ([]*UserTodaTag, error)

	Delete(uint) (uint, error)

}
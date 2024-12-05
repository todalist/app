package todaTagRef

type ITodaTagRefRepo interface {

	// basic crud
	Get(uint) (*TodaTagRef, error)

	First(*TodaTagRefQuerier) (*TodaTagRef, error)

	Save(*TodaTagRef) (*TodaTagRef, error)

	List(*TodaTagRefQuerier) ([]*TodaTagRef, error)

	Delete(uint) (uint, error)

}
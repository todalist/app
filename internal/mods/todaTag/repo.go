package todaTag

type ITodaTagRepo interface {

	// basic crud
	Get(uint) (*TodaTag, error)

	First(*TodaTagQuerier) (*TodaTag, error)

	Save(*TodaTag) (*TodaTag, error)

	List(*TodaTagQuerier) ([]*TodaTag, error)

	Delete(uint) (uint, error)

}
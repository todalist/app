package todaFlow

type ITodaFlowRepo interface {

	// basic crud
	Get(uint) (*TodaFlow, error)

	First(*TodaFlowQuerier) (*TodaFlow, error)

	Save(*TodaFlow) (*TodaFlow, error)

	List(*TodaFlowQuerier) ([]*TodaFlow, error)

	Delete(uint) (uint, error)

}
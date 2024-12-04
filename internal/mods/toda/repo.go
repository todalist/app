package toda

type ITodaRepo interface {

	// basic crud
	Get(uint) (*Toda, error)

	Save(*Toda) (*Toda, error)

	List(*TodaQuerier) ([]*Toda, error)

	Delete(uint) (uint, error)

}
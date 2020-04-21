package services

//IDataService implements the repository access controller
type IDataService interface {
	BeginTx() error
	Commit() error
	Initialize() error
	Create(request interface{}) (interface{}, error)
	Update(id string, request interface{}) (interface{}, error)
	GetById(id string) (interface{}, error)
	GetAll(params string) (interface{}, error)
	GetAllCSV() (interface{}, error)
}

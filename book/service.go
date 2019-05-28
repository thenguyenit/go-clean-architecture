package book

//Service service interface
type Service interface {
	Find(id int64) *Book
	GetRepo() Repository
}

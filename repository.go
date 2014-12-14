package repository
 
import ( 
	"gopkg.in/mgo.v2"
)

type Repository struct {
	databaseName string
	url string
	session *mgo.Session
}
 
func (repo *Repository) Initialize(url, dbname string) {
	repo.url = url
	repo.databaseName = dbname
	repo.session, _ = mgo.Dial(repo.url)
}
 
func (repo *Repository) OpenCollection(collection string) *mgo.Collection {
	return repo.session.DB(repo.databaseName).C(collection)
}
 
func (repo *Repository) Close() {
	repo.session.Close()
}

type M map[string]string
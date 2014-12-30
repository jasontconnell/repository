package repository
 
import ( 
	"gopkg.in/mgo.v2"
)

type Repository struct {
	databaseName string
	databaseServer string
	session *mgo.Session
}
 
func (repo *Repository) Initialize(server, dbname string) {
	repo.databaseServer = server
	repo.databaseName = dbname
	repo.session, _ = mgo.Dial(repo.databaseServer)
}
 
func (repo *Repository) OpenCollection(collection string) *mgo.Collection {
	return repo.session.DB(repo.databaseName).C(collection)
}
 
func (repo *Repository) Close() {
	repo.session.Close()
}
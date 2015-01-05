package repository
 
import ( 
	"gopkg.in/mgo.v2"
)

type Repository struct {
	databaseName string
	databaseServer string
	session *mgo.Session
}
 
func (repo *Repository) Initialize(server, dbname, dbuser, dbpwd string) {
	repo.databaseServer = server
	repo.databaseName = dbname

	upwd := ""
	if dbuser != "" && dbpwd != "" {
		upwd = dbuser + ":" + dbpwd + "@"
	}

	url := "mongodb://" + upwd + server + ":27017/" + dbname
	repo.session, _ = mgo.Dial(url)
}
 
func (repo *Repository) OpenCollection(collection string) *mgo.Collection {
	return repo.session.DB(repo.databaseName).C(collection)
}
 
func (repo *Repository) Close() {
	repo.session.Close()
}
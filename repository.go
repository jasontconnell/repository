package repository
 
import ( 
	"gopkg.in/mgo.v2"
)

type Repository struct {
	session *mgo.Session
}

type MasterConnection struct {
	session *mgo.Session
}

var master *MasterConnection
 
func (repo *Repository) Initialize(config Configuration) {
	upwd := ""
	if config.DatabaseUser != "" && config.DatabasePassword != "" {
		upwd = config.DatabaseUser + ":" + config.DatabasePassword + "@"
	}

	url := "mongodb://" + upwd + config.DatabaseServer + ":27017/" + config.Database

	if master == nil {
		master = new(MasterConnection)
	}

	if master.session == nil {
		master.session, _ = mgo.Dial(url)
		master.session.SetMode(mgo.Monotonic, true)
	}

	repo.session = master.session.Copy()
}
 
func (repo *Repository) OpenCollection(collection string) *mgo.Collection {
	return repo.session.DB("").C(collection)
}
 
func (repo *Repository) Close() {
	repo.session.Close()
}
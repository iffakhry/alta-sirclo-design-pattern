package singleton

var (
	instance *db
	Umum     string
)

type db struct {
	Status string
	Type   string
}

func (d db) GetUser(name string) string {
	return "Hello " + name
}

func GetDBInstance() *db {
	if instance == nil {
		instance = &db{}
		instance.Type = "db"
	}
	return instance
}

package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(posgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: posgreSQL, DatabaseMongoDB: mongoDB}
}

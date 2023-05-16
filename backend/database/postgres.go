package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type ConfigPq struct {
	Host string
	Port int
	User string
	password string
	dbname string
}
var PqConfig ConfigPq = ConfigPq{
	Host: "localhost",
	Port: 5432,
	User: "postgres",
	password: "buckets22",
	dbname: "pern",
}
type PostgresInstance struct {
	DB *sql.DB
}
var Goaldb PostgresInstance;

func Connect(config *ConfigPq){
	psqlConnection := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", 
	config.Host, config.Port, config.User, config.password, config.dbname);

	database, err := sql.Open("postgres", psqlConnection);
	if err != nil {
		panic(err)
	}
	Goaldb = PostgresInstance{
		DB: database,
	}
	//defer database.Close()
    err = database.Ping()
	if err != nil {
		panic(err)
	}
    fmt.Println("Connected to postgres.")
}




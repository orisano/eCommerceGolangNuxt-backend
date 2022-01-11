package db

//
import (
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

//
//var Client *ent.Client
//var err error

//func Open(databaseUrl string) *ent.Client {
//	db, err := sql.Open("pgx", databaseUrl)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// Create an ent.Driver from `db`.
//	drv := entsql.OpenDB(dialect.Postgres, db)
//	return ent.NewClient(ent.Driver(drv))
//}

//func Init() {
//	Client, err = ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=bongobitan password=123456 sslmode=disable")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Connected to database successfully")
//	//defer Client.Close()
//	// Run the auto migration tool.
//	if err := Client.Schema.Create(context.Background()); err != nil {
//		log.Fatalf("failed creating schema resources: %v", err)
//	}
//}

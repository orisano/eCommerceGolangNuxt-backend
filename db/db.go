package db

//
import (
	"bongo/ent"
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
	"log"
)

var Client *ent.Client
var err error

func Init() {
	Client, err = ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=bongo_ent password=123456 sslmode=disable")
	//Client, err = ent.Open("sqlite3", "file:bongobitan?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database successfully")

	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

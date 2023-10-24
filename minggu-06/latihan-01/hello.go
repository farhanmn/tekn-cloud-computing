package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectSQL() *sql.DB {
	db, err := sql.Open("mysql", "root:rootP@ss@/sample_db")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	println("\nSQL Connected!\n")
	return db
}

var ctx = context.Background()

func connectMongo() (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	println("\nMongoDB Connected!\n")
	return client.Database("sample_db"), nil
}

func selectSQL() {
	sql, err := connectSQL().Query("SELECT * FROM tb_users")
	if err != nil {
		panic(err.Error())
	}

	println("---------------MySQL Section----------------------")
	for sql.Next() {
		var user User
		err = sql.Scan(&user.ID, &user.Name, &user.Address)
		if err != nil {
			panic(err.Error())
		}

		log.Println(user.ID, user.Name, user.Address)
	}
	println("-------------------------------------------------")
}

func findWithMongo() {
	db, err := connectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("sample").Find(ctx, bson.M{"storeLocation": "Denver"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]Trx, 0)
	for csr.Next(ctx) {
		var row Trx
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		println("--------------Mongo Section-----------------------")
		for i := 0; i < 6; i++ {
			fmt.Println("Items  :", result[i].Items)
			fmt.Println("Location :", result[i].StoreLocation)
			fmt.Println("Method :", result[i].PurchaseMethod)
		}
		println("-------------------------------------------------")
	}
}

type User struct {
	ID      int    `json:"user_id"`
	Name    string `json:"user_name"`
	Address string `json:"user_address"`
}

type Item struct {
	Name     string               `json:"name"`
	Price    primitive.Decimal128 `json:"price"`
	Quantity int                  `json:"quantity"`
}

type Trx struct {
	Items          []Item `bson:"items"`
	StoreLocation  string `bson:"storeLocation"`
	PurchaseMethod string `bson:"purchaseMethod"`
}

func main() {
	selectSQL()

	findWithMongo()
}

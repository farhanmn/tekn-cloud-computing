package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"context"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = make(map[string]string)

func connectSQL() *sql.DB {
	db, err := sql.Open("mysql", "root:rootP@ss@/sample_db")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
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

func selectSQL() []User {
	sql, err := connectSQL().Query("SELECT * FROM tb_users")
	if err != nil {
		panic(err.Error())
	}
	var users []User
	for sql.Next() {
		var user User
		err = sql.Scan(&user.ID, &user.Name, &user.Address)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}
	if err = sql.Err(); err != nil {
		return users
	}
	return users
}

func findWithMongo() []Trx {
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
	finalResult := make([]Trx, 0)
	for csr.Next(ctx) {
		var row Trx
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		for i := 0; i < 4; i++ {
			finalResult = append(finalResult, result[i])
		}
	}

	return finalResult
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

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	r.GET("/sqlcheck", func(c *gin.Context) {
		result := selectSQL()
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	r.GET("/mongocheck", func(c *gin.Context) {
		result := findWithMongo()
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

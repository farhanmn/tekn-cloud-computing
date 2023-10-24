// credit - go-graphql hello world example
package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"net/http"

	"github.com/gin-gonic/gin"

	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

type User struct {
	ID      int    `json:"user_id"`
	Name    string `json:"user_name"`
	Address string `json:"user_address"`
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Address": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var userList = graphql.NewList(userType)

type Item struct {
	Name     string               `json:"name"`
	Price    primitive.Decimal128 `json:"price"`
	Quantity int                  `json:"quantity"`
}

var itemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Item",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Price": &graphql.Field{
				Type: graphql.Float,
			},
			"Quantity": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        userList,
				Description: "Get user by id",
				Args: graphql.FieldConfigArgument{
					"ID": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// id, ok := p.Args["ID"].(int)

					for _, user := range selectSQL() {
						log.Println(user)
						return user, nil
					}
					return nil, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery() *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: `
		{
			user(ID: 3) {
				ID,
				Name,
				Address
			}
		}`,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
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
		// connectSQL()
		result := executeQuery()
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

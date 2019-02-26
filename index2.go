package main

import (
	"log"

	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

type Car struct {
	Name string
	Type string
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// Database connection
	session, err := mgo.Dial("localhost")
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Database name and collection name
	// car-db is database name car is collation name
	c := session.DB("car-db").C("car")
	c.Insert(&Car{"Audi", "Luxury car"})

	// Get /car
	app.Get("/car", func(ctx iris.Context) {
		result := Car{}
		err = c.Find(bson.M{"name": "Audi"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(result)
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	// http://localhoist:8080/car
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

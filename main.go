package main

import (
	routes "e-vet/Routes"
	"e-vet/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// type details struct {
//     fname   string
//     lname   string
//     age     int
//     balance float64
// }

// type myType string

// func showDetails(i, j interface{}) {
//     t1 := reflect.TypeOf(i)
//     k1 := t1.Kind()
//     t2 := reflect.TypeOf(j)
//     k2 := t2.Kind()
//     fmt.Println("Type of first interface:", t1)
//     fmt.Println("Kind of first interface:", k1)
//     fmt.Println("Type of second interface:", t2)
//     fmt.Println("Kind of second interface:", k2)

//     fmt.Println("The values in the first argument are :")
//     if reflect.ValueOf(i).Kind() == reflect.Struct {
//         value := reflect.ValueOf(i)
//         numberOfFields := value.NumField()
//         for i := 0; i < numberOfFields; i++ {
//             fmt.Printf("%d.Type:%T || Value:%#v\n",
//               (i + 1), value.Field(i), value.Field(i))

//             fmt.Println("Kind is ", value.Field(i).Kind())
//         }
//     }
//     value := reflect.ValueOf(j)
//     fmt.Printf("The Value passed in "+
//       "second parameter is %#v", value)

// }

func main() {
	db.Dbs["main"].Connect()
	// db.DBConn.Migrate()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/mig/:mig?/:seed?", func(c *fiber.Ctx) error {
		ext := c.Params("ext")

		// db.DBConn.DropSchema("public")

		db.DBConn.Migrate()

		if ext == "seed" {
			db.DBConn.SeedSchema()
			return c.SendString("public dropped and created and seeded")

		}

		return c.SendString("public dropped and created")
	})

	routes.Setup(app)
	app.Listen(":8000")
}

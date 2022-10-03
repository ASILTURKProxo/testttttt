package controllers

import (
	"e-vet/db"
	models "e-vet/models"
	"fmt"
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func UserAbilityCreateController(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	dbC := db.DBConn.DB
	bearerToken := c.Get("Authorization")
	tokenString := bearerToken[7:]
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	//get user token in jwt claims
	userID := token.Claims.(jwt.MapClaims)["userID"]
	// dbC.Model(&user).Association("Languages").Append(&Language{Name: "DE"})
	user := models.User{} //User model from models/user.go
	// dbC.Omit("Name", "email").Find(&user, userID)  //Find user by id and ignore name and email
	dbC.Find(&user, userID)

	//Create ability map in token
	a := []int{1, 2}
	//user ability create in jwt

	boolen := sync(user, models.Ability{}, "user_abilities", a)

	if boolen {
		return c.JSON(map[string]interface{}{
			"message": "user ability not created",
		})
	}

	return c.JSON(map[string]interface{}{"token": user, "error": boolen})

	// return nil
}

func sync(firstModel, lastModel interface{}, intermediateTable string, a []int) bool {
	// t1 := reflect.TypeOf(getModels) //modelin ne olduğunu dönderiyor
	// k1 := t1.Kind() //tipini dönderiyor
	// value := reflect.ValueOf(firstModel)
	fmt.Println("---------------------------")

	fmt.Println(reflect.TypeOf(firstModel)) //modelin adı geliyor

	// lastModel = models.Ability{ID: 3}
	fmt.Println("-----------SORGUM----------------")

	// relation := reflect.TypeOf(firstModel).Field(reflect.TypeOf(lastModel).Name())

	db.DBConn.DB.Model(firstModel).Association(reflect.TypeOf(lastModel).Name()).Replace()
	fmt.Println("-----------SORGUM BİT----------------")

	// db.DBConn.DB.Model(firstModel).Association(reflect.TypeOf(lastModel).Name()).Append()

	fmt.Print(reflect.TypeOf(lastModel).Name())
	// if reflect.ValueOf(firstModel).Kind() == reflect.Struct {
	// 	value := reflect.ValueOf(firstModel)
	// 	numberOfFields := value.NumField()
	// 	for i := 0; i < numberOfFields; i++ {
	// 		if reflect.ValueOf(firstModel).Kind() == reflect.ValueOf(lastModel).Kind() {

	// 			fmt.Printf("Type:%+v ", value.(i))

	// 		}
	// 	}
	// }

	// fmt.Printf("%+v", value)
	// dbC := db.DBConn.DB
	// var removeAbilties []models.Ability
	// dbC.Table(intermediateTable).Where("user_id", &value.ID).Delete(&removeAbilties)
	// var tempMapAbility []map[string]interface{}

	// for _, casd := range a {
	// 	fmt.Println(casd)
	// 	tempMapAbility = append(tempMapAbility, map[string]interface{}{
	// 		"user_id":    getModels.ID,
	// 		"ability_id": casd,
	// 	})
	// }
	// err := dbC.Table("user_abilities").Create(&tempMapAbility).Error
	// if err != nil {
	// 	return false
	// }
	// fmt.Println(removeAbilties)
	// fmt.Print(intermediateTable)
	return true
}

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

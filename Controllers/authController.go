package controllers

import (
	"time"

	"e-vet/db"
	models "e-vet/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const SecretKey = "jwtSecret"

func HashPassword(password []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	return bytes, err
}

func CheckPasswordHash(password []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := HashPassword([]byte(data["password"]))

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	db.DBConn.DB.Create(&user)

	return c.JSON(map[string]interface{}{
		"user": user,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	//Token time to live
	tokenTime := time.Now().Add(time.Hour * 24)
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	password := data["password"]
	//User model from models/user.go
	var user models.User
	//Find user by email
	db.DBConn.DB.Preload("Ability", func(db *gorm.DB) *gorm.DB {
		return db.Joins("left JOIN modules ON modules.id = abilities.module_id").Select("modules.name as module_name", "modules.subject as module_key", "abilities.*")
	}).Where("email = ?", data["email"]).First(&user)
	//Check if user password is correct
	match := CheckPasswordHash([]byte(password), user.Password)
	if !match {
		c.Status(401)
		return c.JSON(map[string]interface{}{
			"message": "Incorrect password",
		})
	}
	//Create ability map in token
	var tempMapAbility = make(map[string]interface{})
	//user ability create in jwt
	for _, ability := range user.Ability {
		tempMapAbility[ability.ModuleKey] = []string{
			ability.Action,
			ability.Subject,
		}
	}
	//Create jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ability":   tempMapAbility,
		"email":     user.Email,
		"fullname":  user.Name,
		"userID":    user.ID,
		"ip":        c.IP(),
		"isMaster":  false,
		"product":   "bot",
		"ExpiresAt": tokenTime.Unix(), //1day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(401)
		return c.JSON(map[string]interface{}{
			"message": "User not found",
		})

	}

	return c.JSON(map[string]interface{}{
		"user":        user,
		"accessToken": token,
		"success":     "Success",
	})

}

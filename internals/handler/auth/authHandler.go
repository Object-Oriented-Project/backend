package auth

import (
	"log"
	"os"
	"time"

	"github.com/Object-Oriented-Project/backend/database"
	"github.com/Object-Oriented-Project/backend/internals/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoginToken(c *fiber.Ctx) error {
	userLogin := new(model.User)
	userInfoFormDB := new(model.User)

	// Parse login info to userLogin
	if err := c.BodyParser(userLogin); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	db := database.DB
	// find user in DB
	result := db.Where("Username = ?", userLogin.Username).First(&userInfoFormDB)

	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot find user",
		})
	}

	// Compare user password
	if err := bcrypt.CompareHashAndPassword([]byte(userInfoFormDB.PasswordHash), []byte(userLogin.PasswordHash)); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Read .env file to get jwt secret
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Create token
	jwtSecretKey := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Username"] = userInfoFormDB.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"Token": t,
	})
}

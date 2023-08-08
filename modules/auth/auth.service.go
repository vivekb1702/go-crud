package auth

import (
	"go-crud/initializers"
	"go-crud/models"
	"go-crud/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginUserService(email string, password string) (*LoginResponse, error) {

	user := models.User{}
	findUser := initializers.DB.Where("email = ?", email).First(&user)
	if findUser.RowsAffected == 0 {
		return nil, utils.CustomError{Code: 404, Message: "User not found"}
	}
	comparePasswords := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if comparePasswords != nil {
		return nil, utils.CustomError{Code: 401, Message: "Invalid Credentials"}
	}
	token, err := GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: token}, nil
}

func GenerateToken(u models.User) (string, error) {

	key := []byte(os.Getenv("DB_URL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    u.ID,
		"exp":   time.Now().Add(15 * time.Minute),
		"email": u.Email,
	})
	jToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return jToken, nil

}

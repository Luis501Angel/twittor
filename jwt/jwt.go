package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/luis501angel/go-course/models"
)

func GenerateJWT(user models.User) (string, error) {
	myKey := []byte("Luis501Angel")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastname":  user.LastName,
		"birthday":  user.Birthday,
		"biography": user.Biography,
		"location":  user.Location,
		"website":   user.Website,
		"_id":       user.Id.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

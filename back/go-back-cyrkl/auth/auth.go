package auth

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

/*
 This function generates a new JSON Web Token
 The function also accepts the users username as a input
*/
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user"] = username
	claims["aud"] = "surfspots.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	/*
	   The token will expire after 24 hours / 1 day
	*/
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		err = errors.New("Something went wrong")
		return "", err
	}
	return newToken, nil
}

/*
 This function checks if the JWT is still valid
*/
func CheckJWT(token string) error {
	t, err := jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return err
	} else if !t.Valid {
		err = errors.New("Invalid token")
		return err
	} else if t.Claims.(jwt.MapClaims)["aud"] != "surfspots.jwtgo.io" {
		err = errors.New("Invalid aud")
		return err
	} else if t.Claims.(jwt.MapClaims)["iss"] != "jwtgo.io" {
		err = errors.New("Invalid iss")
		return err
	}
	return nil
}

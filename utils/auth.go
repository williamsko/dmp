package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CreateToken : create jwt token for agent
func CreateToken(matricule string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["matricule"] = matricule
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// TokenValid : Tell if token is valid
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// VerifyToken : check validity of a token
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

//ExtractToken get the token from the request body
func ExtractToken(r *http.Request) string {

	bearToken := r.Header.Get("Authorization")

	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//ExtractTokenAuth get the token from the request header
func ExtractTokenAuth(r *http.Request) string {
	token, err := VerifyToken(r)

	if err != nil {
		return ""
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		return claims["matricule"].(string)
	}
	return ""
}

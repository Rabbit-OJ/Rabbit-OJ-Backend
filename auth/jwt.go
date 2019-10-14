package auth

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type StandardClaims = jwt.StandardClaims

type Claims struct {
	Uid      uint32 `json:"uid"`
	Username string `json:"username"`

	StandardClaims
}

func SignJWT(user *models.User) (string, error) {
	nextTime := time.Now()
	nextTime = nextTime.AddDate(0, 0, 1)

	jwtObject := &Claims{
		Uid:      user.Uid,
		Username: user.Username,
		StandardClaims: StandardClaims{
			ExpiresAt: nextTime.Unix(),
		},
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS512, jwtObject)

	return tokenObj.SignedString([]byte(utils.Secret))
}

func VerifyJWT(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.Secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetAuthObj(c *gin.Context) (*Claims, error) {
	authObjectRaw, exists := c.Get("AuthObject")

	if !exists {
		return nil, errors.New("access denied")
	}

	authObject, ok := authObjectRaw.(*Claims)
	if !ok {
		return nil, errors.New("token format invalid")
	}

	return authObject, nil
}

func GetAuthObjRequireAdmin(c *gin.Context) (*Claims, error) {
	authObj, err := GetAuthObj(c)

	if err != nil {
		return nil, err
	}

	if !authObj.IsAdmin {
		return nil, errors.New("admin required")
	}

	return authObj, nil
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("secret")

type user struct {
	UserID   uint64 `json:"userID"`
	UserName string `json:"userName"`
}

type userClaims struct {
	user
	jwt.StandardClaims
}

// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Router /testapi/get-string-by-int/{some_id} [get]
func handleAuth(c *gin.Context) {
	data := make(map[string]interface{})
	token, err := generateToken("admin", "admin")
	fmt.Println(err)
	if err != nil {
		data["error"] = err
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "ERROR",
			"data": data,
		})
		return
	}
	data["token"] = token
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "SUCCESS",
		"data": data,
	})
}

func handleJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.Query("token")
		if token == "" {
			code = 0
		} else {
			claims, err := parseToken(token)
			if err != nil {
				code = 1
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 2
			}
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

func generateToken(username, password string) (string, error) {
	param := struct {
		Username string
		Password string
	}{
		Username: username,
		Password: password,
	}

	if param.Username != "admin" || param.Password != "admin" {
		return "", fmt.Errorf("invalid login")
	}

	claims := userClaims{
		user: user{
			UserID:   1,
			UserName: param.Username,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
			Issuer:    "demo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func parseToken(jwtToken string) (*userClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	v, ok := token.Claims.(*userClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("error")
	}
	return v, nil
}

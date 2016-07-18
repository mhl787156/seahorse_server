package auth

import (
	"fmt"
	"time"

	//"github.com/saturi/go.uuid"
	//"golang.ord/x/crypt/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/mhl787156/seahorse_server/models"
)

const (
	tokenDuration = 72
	expireOffset  = 3600
)

func GenerateToken(userUUID string) (string, []byte, error) {
	secretkey := getRandomBytes()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(tokenDuration)).Unix(),
		"iat": time.Now().Unix(),
		"sub": userUUID,
	})
	tokenString, err := token.SignedString(secretkey)

	return tokenString, secretkey, err
}

func Authenticate(user *models.User) bool {
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)

	// testUser := models.User{
	// 	UUID:     uuid.NewV4(),
	// 	Username: "haku",
	// 	Password: string(hashedPassword),
	// }

	// return user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
	fmt.Println("Authenticating this User!")
	return false
}

func getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

func Logout(tokenString string, token *jwt.Token) error {
	// redisConn := redis.Connect()
	// return redisConn.SetValue(tokenString, tokenString, backend.getTokenRemainingValidity(token.Claims["exp"]))
	return nil
}

func IsInBlacklist(token string) bool {
	// redisConn := redis.Connect()
	// redisToken, _ := redisConn.GetValue(token)

	// if redisToken == nil {
	// 	return false
	// }

	return false
}

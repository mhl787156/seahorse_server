package auth

import (
	"fmt"
	"time"

	//"github.com/saturi/go.uuid"
	//"golang.ord/x/crypt/bcrypt"

	"github.com/dgrijalva/jwt-go"
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

func Authenticate(tokenString string, secretkey []byte) (string, bool) {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretkey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if str, ok := claims["sub"].(string); ok {
			/* check is string */
			return str, true
		}
	}

	fmt.Println("Error:", err)
	return "", false
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

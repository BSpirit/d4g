package handlers

import (
	"crypto/sha512"
	"d4g/app/models"
	"d4g/app/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func AuthenticationHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	login := r.FormValue("login")
	password := r.FormValue("password")
	access, err := models.GetAccessFromLogin(env.DB, login)
	if err != nil {
		return &StatusError{Code: 401, Err: utils.Trace(fmt.Errorf("login %s does not exist", login))}
	}

	sha512 := sha512.New()
	sha512.Write([]byte(password))
	hashedPassword := string(sha512.Sum(nil))
	if hashedPassword != access.Password {
		return &StatusError{Code: 401, Err: utils.Trace(fmt.Errorf("login %s: bad credentials", login))}
	}

	// hours, err := time.ParseDuration("2m")
	// if err != nil {
	// 	return &StatusError{Code: 500, Err: utils.Trace(err)}
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":    login,
		"is_admin": access.IsAdmin,
		"expire":   time.Now().Add(1.2E11).Unix(),
	})

	tokenString, err := token.SignedString(env.JWTKey)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	w.Header().Set("Content-Type", "application/json")
	data := map[string]interface{}{
		"JWT": tokenString,
	}

	json.NewEncoder(w).Encode(data)
	return nil
}

func AccessRoleHandler(env *Env, w http.ResponseWriter, r *http.Request) *StatusError {
	tokenString := r.Header.Get("Authorization")
	re, err := regexp.Compile(`Bearer (?P<token>.+)`)
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	tokenString = re.FindStringSubmatch(tokenString)[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, utils.Trace(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
		}

		return env.JWTKey, nil
	})
	if err != nil {
		return &StatusError{Code: 500, Err: utils.Trace(err)}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return &StatusError{Code: 403, Err: utils.Trace(fmt.Errorf("invalid JWT"))}
	}

	expireTime := time.Unix(int64(claims["expire"].(float64)), 0)
	if time.Now().After(expireTime) {
		return &StatusError{Code: 401, Err: utils.Trace(fmt.Errorf("expired JWT"))}
	}

	w.Header().Set("Content-Type", "application/json")
	data := map[string]interface{}{
		"is_admin": claims["is_admin"],
	}

	json.NewEncoder(w).Encode(data)
	return nil
}

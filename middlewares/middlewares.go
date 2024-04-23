package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// app issuer
var issuer string = "meetings.app"

type Claims struct {
    UserName string `json:"name"`
    jwt.RegisteredClaims
}

func IsAuthorized(next http.Handler) http.Handler {
    var myKey = []byte(os.Getenv("JWT_SECRET"))

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        _, _, err := verifyToken(w, r)
        if err != nil {
            return 
        }

        next.ServeHTTP(w, r)
        return 
    })
}

func verifyToken(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {
    w.Header().Add("Vary", "Authorization")

    // check for headers
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        return "", nil, errors.New("no auth header")
    }

    // check that there are two parts: bearer and the token
    headerParts := strings.Split(authHeader, " ")
    if len(headerParts) != 2 { // if it doesn't have Bearer Token...
        return "", nil, errors.New("unauthorized: no Bearer")
    }

    token := headerParts[1]
    claims := &Claims{}
    secret := os.Getenv("JWT_SECRET")

    // parse token
    _, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secret), nil
    })
    if err != nil {
       if strings.HasPrefix(err.Error(), "token is expired by") {
            return "", nil, errors.New("expired token")
        } 
        return "", nil, err
    }

    // is issuers of token don't match
    if claims.Issuer != issuer {
        return "", nil, errors.New("incorrect issure")
    }

    return token, claims, nil
}

package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/arturfil/meetings_app_server/helpers"
	"github.com/arturfil/meetings_app_server/types"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	store types.UserStore // interface type
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Get("/v1/healthcheck", h.healthCheck)
	router.Get("/v1/teachers", h.getAllTeachers)
	router.Post("/v1/signup", h.signupUser)
	router.Post("/v1/login", h.loginUser)
    router.Get("/v1/users/bytoken", h.getUserByToken)
}

func (h *Handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Msg  string
		Code int
	}{
		Msg:  "Making sure this works",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

func (h *Handler) signupUser(w http.ResponseWriter, r *http.Request) {
	var body types.RegisterUserPayload

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.ErrorJSON(w, fmt.Errorf("Couldn't read the json data %v", err), http.StatusBadRequest)
		return
	}

	// In the handler -> check that users is not already signed up
	_, err = h.store.GetUserByEmail(body.Email) // if users exists that should be wrong
	if err == nil {
		helpers.ErrorJSON(w, fmt.Errorf("user already exists %s", body.Email), http.StatusInternalServerError)
		return
	}

	// encrypt users password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		helpers.ErrorJSON(w, fmt.Errorf("something went wrong %v", err), http.StatusInternalServerError)
		return
	}

	// save user in the db
	err = h.store.CreateUser(types.RegisterUserPayload{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(hashedPassword),
	})

	helpers.WriteJSON(w, http.StatusNoContent, "Successfully signed up the user")
}

func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	// body
	var body types.LoginUserPayload

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helpers.ErrorJSON(w, fmt.Errorf("Couldn't read the json data %v", err), http.StatusBadRequest)
		return
	}

	// fmt.Print("body", body.Email)

	// check if users exists
	user, err := h.store.GetUserByEmail(body.Email)
	if err != nil {
		helpers.ErrorJSON(w, fmt.Errorf("User does not exist %v", err), http.StatusInternalServerError)
		return
	}

	// check if the passwords encrypted match
	if !passwordMatches(user.Password, body.Password) {
		helpers.ErrorJSON(w, fmt.Errorf("Invalid credentials"), http.StatusBadRequest)
	}

	// get the secret from the env & generate JWT
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := createJWT(secret, user.ID)
	if err != nil {
		helpers.ErrorJSON(w, fmt.Errorf("Couldn't generate the token %v", err))
		return
	}

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

// getUserByToken - you will get the user by when providing a jwt token
func (h *Handler) getUserByToken(w http.ResponseWriter, r *http.Request) {
	var myKey = []byte(os.Getenv("JWT_SECRET"))
    fmt.Println("sKey ", myKey)
	type TokenClaim struct {
		Sub string `json:"sub"`
		Aud string `json:"aud"`
		Iss string `json:"iss"`
		Exp string `json:"exp"`
		jwt.StandardClaims
	}

	claims := &TokenClaim{}

	token, err := jwt.ParseWithClaims(strings.Split(r.Header["Authorization"][0], " ")[1], claims, func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fmt.Errorf("there was an error")
		// }
		return myKey, nil
	})

    log.Println("before valid: ", token.Valid)

	if token.Valid {
		if err != nil {
			helpers.MessageLogs.InfoLog.Print(err)
		}

		user, err := h.store.GetUserById(claims.Sub)
		if err != nil {
			helpers.MessageLogs.InfoLog.Print(err)
		}

        log.Print("user", user)

		user.Password = ""
		helpers.WriteJSON(w, http.StatusOK, user)
		return
	}

}

func (h *Handler) getAllTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := h.store.GetTeachers()
	if err != nil {
		helpers.ErrorJSON(w, fmt.Errorf("Couldn't get the teachers %v", err))
		return
	}

	helpers.WriteJSON(w, http.StatusOK, teachers)
}

func passwordMatches(hashed, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText)) // plain text also hashed => return stored == entered
	return err == nil
}

func createJWT(secret []byte, id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	expiration := time.Now().Add(time.Hour * 24).Unix() // add in unix time

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = fmt.Sprint(id)
	claims["aud"] = types.Domain
	claims["iss"] = types.Domain
	claims["exp"] = expiration

	// maybe add admin here

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

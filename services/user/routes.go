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
	"github.com/arturfil/meetings_app_server/middlewares"
	"github.com/arturfil/meetings_app_server/types"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
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

	router.Route("/v1/teachers", func(router chi.Router) {
		router.Get("/", h.getAllTeachers)
		router.Get("/search", h.searchTeachers)
	})

	router.Route("/v1/auth", func(router chi.Router) {
		router.Post("/signup", h.signupUser)
		router.Post("/login", h.loginUser)
	})

	router.Route("/v1/users", func(router chi.Router) {
        router.Get("/bytoken", h.getUserByToken)
	})

    router.Route("/v1/users/admin", func(router chi.Router) {
		router.Use(middlewares.IsAuthorized)
		router.Get("/", h.getAllUsers)
    })


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

	// check if users exists
	user, err := h.store.GetUserByEmail(body.Email)
	if err != nil {
		helpers.ErrorJSON(w, fmt.Errorf("Invalid credentials, please try again"), http.StatusInternalServerError)
		log.Println("User does not exist", err)
		return
	}

	// check if the passwords encrypted match
	if !passwordMatches(user.Password, body.Password) {
		helpers.ErrorJSON(w, fmt.Errorf("Invalid credentials"), http.StatusBadRequest)
		return
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

func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetAllUsers()
	if err != nil {
		helpers.ErrorJSON(w, fmt.Errorf("Couldn't get the teachers", err))
		return
	}

	helpers.WriteJSON(w, http.StatusOK, users)
}

// getUserByToken - you will get the user by when providing a jwt token
func (h *Handler) getUserByToken(w http.ResponseWriter, r *http.Request) {
	var myKey = []byte(os.Getenv("JWT_SECRET"))

	claims := &types.TokenClaim{}

	tokenString := strings.Split(r.Header["Authorization"][0], " ")[1]

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return myKey, nil
	})

	if token.Valid {
		if err != nil {
			helpers.MessageLogs.InfoLog.Print(err)
		}

		user, err := h.store.GetUserById(claims.Sub)
		if err != nil {
			helpers.MessageLogs.InfoLog.Print(err)
		}

		helpers.WriteJSON(w, http.StatusOK, user)
		return
	}

	helpers.ErrorJSON(w, fmt.Errorf("Couldn't generate the token %v", err))
}

func (h *Handler) getAllTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := h.store.GetTeachers()
	if err != nil {
		helpers.ErrorJSON(w, fmt.Errorf("Couldn't get the teachers %v", err))
		return
	}

	helpers.WriteJSON(w, http.StatusOK, teachers)
}

func (h *Handler) searchTeachers(w http.ResponseWriter, r *http.Request) {
	searchWord := r.URL.Query().Get("queryWord")

	teachers, err := h.store.SearchTeachers(searchWord)
	if err != nil {
		helpers.ErrorJSON(w, fmt.Errorf("Something went wrong", err))
	}

	helpers.WriteJSON(w, http.StatusOK, teachers)
}

func passwordMatches(hashed, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText)) // plain text also hashed => return stored == entered
	return err == nil
}

func createJWT(secret []byte, id string) (string, error) {
	// token := jwt.New(jwt.SigningMethodHS256)

	expiration := time.Now().Add(time.Hour * 24).Unix() // add in unix time

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.TokenClaim{
		RegisteredClaims: jwt.RegisteredClaims{},
		Sub:              fmt.Sprint(id),
		Aud:              types.Domain,
		Iss:              types.Domain,
		Exp:              fmt.Sprint(expiration),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/arturfil/meetings_app_server/helpers"
	"github.com/arturfil/meetings_app_server/services"
	"github.com/golang-jwt/jwt/v4"
)

var user services.User

// Signup - this handler will manage the signup logic
func signup(w http.ResponseWriter, r *http.Request) {
    err := json.NewDecoder(r.Body).Decode(&user)    
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return 
    }
    
    _, err = user.Signup(user)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Panicln(err) 
        return 
    }

    res := Response{
        Msg: "succesfully signed up",
        Code: http.StatusOK,
    }

    helpers.WriteJSON(w, http.StatusOK, res)
}

func login(w http.ResponseWriter, r *http.Request) {
    myKey := []byte(os.Getenv("SECRET_KEY"))

    type credentials struct {
        Email string `json:"email"`
        Password string `json:"password"`
    }

    var creds credentials
    var payload services.JsonResponse

    // read json
    err := helpers.ReadJSON(w, r, &creds)
    if err != nil {
        payload.Error = true
        payload.Message = "Invalid json supplied"
        _ = helpers.WriteJSON(w, http.StatusBadRequest, payload)
        return 
    }

    // get user if user exists
    user, err := user.GetByEmail(creds.Email)
    if err != nil {
        str := fmt.Sprintf("invalid -> %v",  err)
        helpers.ErrorJSON(w, errors.New(str))
        return 
    }

    // check if password password matches
    validPassword, err := user.PasswordMatches(creds.Password)
    if err != nil || !validPassword {
        helpers.ErrorJSON(w, errors.New("wrong credentials, please try again"))
        return 
    }

    // if you reach this part, means user exists and password is valid => create jwt token
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)
    claims["authorized"] = true
    claims["name"] = user.FirstName
    claims["email"] = user.Email
    claims["exp"] = time.Now().Add(time.Minute * 60 * 4).Unix()

    tokenString, err := token.SignedString(myKey)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        helpers.ErrorJSON(w, err)
        return 
    }

    user.Password = "hidden"

    // create response
    response := services.TokenResponse{
        Token: tokenString,
        User: user,
    }

    // write response if no errors, else it will be logged
    err = helpers.WriteJSON(w, http.StatusOK, response)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
    }

}

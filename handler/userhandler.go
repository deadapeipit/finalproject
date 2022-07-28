package handler

import (
	"context"
	"encoding/json"
	"finalproject/database"
	"finalproject/entity"
	s "finalproject/server"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct{}

func InstallUsersHandler(r *mux.Router) {
	api := UserHandler{}
	r.HandleFunc("/users/{id}", api.UsersHandler)
	r.HandleFunc("/users", api.UsersHandler)
}

type UserHandlerInterface interface {
	UsersHandler(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler() UserHandlerInterface {
	return &UserHandler{}
}

func (h *UserHandler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodPost:
		if id == "login" {
			loginUserHandler(w, r)
		} else if id == "register" {
			registerUsersHandler(w, r)
		} else {
			updateUserHandler(w, r, id)
		}
	case http.MethodPut:
		updateUserHandler(w, r, id)
	case http.MethodDelete:
		deleteUserHandler(w, r, id)
	default:
		s.WriteJsonResp(w, s.ErrorNotFound, "PAGE NOT FOUND")
		return
	}
}

// loginUserHandler
// Method: POST
// Example: localhost/login
// JSON Body:
// {
// 	"email": "user@email.com",
// 	"password": "password"
// }
func loginUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	decoder := json.NewDecoder(r.Body)
	var inp entity.UserLogin
	if err := decoder.Decode(&inp); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
	id, pw, err := database.SqlDatabase.Login(ctx, inp.Username)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(pw), []byte(inp.Password))
	if err != nil {
		s.WriteJsonResp(w, s.ErrorUnauthorized, "UNAUTHORIZED")
		return
	}
	claims := entity.MyClaims{
		Iat: int(time.Now().UnixMilli()),
		Exp: int(time.Now().Add(time.Second * time.Duration(60)).UnixMilli()),
		Uid: id,
	}

	token := jwt.NewWithClaims(
		s.JWT_SIGNING_METHOD,
		claims,
	)

	tokenVal, err := token.SignedString([]byte(s.Config.SecretKey))
	if err != nil {
		s.WriteJsonResp(w, s.ErrorBadRequest, "BAD_REQUEST")
		return
	}
	retVal := map[string]string{
		"token": tokenVal,
	}
	s.WriteJsonResp(w, s.Success, retVal)
}

// registerUsersHandler
// Method: POST
// Example: localhost/register
// JSON Body:
// {
//		"username": "user1",
//		"email": "user@email.com",
//		"password": "password1",
//		"age": 22
// }
func registerUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	validate := validator.New()
	decoder := json.NewDecoder(r.Body)
	var inp entity.UserRegister
	if err := decoder.Decode(&inp); err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	err := validate.Struct(inp)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorBadRequest, err.Error())
		return
	}

	encrtedPwd, err := s.EncryptPassword(inp.Password)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	inp.Password = encrtedPwd

	users, err := database.SqlDatabase.Register(ctx, inp)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	s.WriteJsonResp(w, s.Success201, users)
}

// updateUserHandler
// Method: PUT
// Example: localhost/users/1
// JSON Body:
// {
//		"username": "user1",
//		"email": "user@email.com"
// }
func updateUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()

	if id != "" { // get by id
		if idInt, err := strconv.ParseInt(id, 10, 64); err == nil {
			decoder := json.NewDecoder(r.Body)
			validate := validator.New()
			var inp entity.UserUpdate
			if err := decoder.Decode(&inp); err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			err := validate.Struct(inp)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorBadRequest, err.Error())
				return
			}
			if idInt != s.LogonUser.ID {
				s.WriteJsonResp(w, s.ErrorUnauthorized, "UNAUTHORIZED")
				return
			}
			users, err := database.SqlDatabase.UpdateUser(ctx, idInt, inp.Email, inp.Username)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			retVal := users.ToUserUpdateOutput()
			s.WriteJsonResp(w, s.Success, retVal)
		}
	}
}

// deleteUserHandler
// Method: DELETE
// Example: localhost/users/1
func deleteUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()
	if id != "" {
		if idInt, err := strconv.ParseInt(id, 10, 64); err == nil {
			if idInt != s.LogonUser.ID {
				s.WriteJsonResp(w, s.ErrorUnauthorized, "UNAUTHORIZED")
				return
			}
			users, err := database.SqlDatabase.DeleteUser(ctx, idInt)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			retVal := map[string]string{
				"message": users,
			}
			s.WriteJsonResp(w, s.Success, retVal)

		}
	}
}

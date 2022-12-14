package handler

import (
	"context"
	"encoding/json"
	"finalproject/database"
	"finalproject/entity"
	s "finalproject/server"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type SocialMediaHandler struct{}

func InstallSocialMediaHandler(r *mux.Router) {
	api := SocialMediaHandler{}
	r.HandleFunc("/socialmedias/{id}", api.SocialMediasHandler)
	r.HandleFunc("/socialmedias", api.SocialMediasHandler)
}

type SocialMediaHandlerInterface interface {
	SocialMediasHandler(w http.ResponseWriter, r *http.Request)
}

func NewSocialMediaHandler() SocialMediaHandlerInterface {
	return &SocialMediaHandler{}
}

func (h *SocialMediaHandler) SocialMediasHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodGet:
		getSocialMediasHandler(w, r)
	case http.MethodPost:
		postSocialMediaHandler(w, r)
	case http.MethodPut:
		updateSocialMediaHandler(w, r, id)
	case http.MethodDelete:
		deleteSocialMediaHandler(w, r, id)
	default:
		s.WriteJsonResp(w, s.ErrorNotFound, "PAGE NOT FOUND")
		return
	}
}

// getSocialMediasHandler
// Method: GET
// Example: localhost/socialmedias
func getSocialMediasHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	retVal, err := database.SqlDatabase.GetSocialMedias(ctx)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}

	s.WriteJsonResp(w, s.Success, retVal)
}

// postSocialMediaHandler
// Method: POST
// Example: localhost/socialmedias
// JSON Body:
// {
// 	"name": "social media name",
// 	"social_media_url": "https://domainsocialmedia.com/user",
// 	"profile_image_url": "https://domainsocialmedia.com/userimage.jpg"
// }
func postSocialMediaHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	validate := validator.New()
	decoder := json.NewDecoder(r.Body)
	var inp entity.SocialMediaPost
	if err := decoder.Decode(&inp); err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	err := validate.Struct(inp)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorBadRequest, err.Error())
		return
	}
	p, err := database.SqlDatabase.PostSocialMedia(ctx, s.LogonUser.ID, inp)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}

	retVal := p.ToSocialMediaPostOutput()

	s.WriteJsonResp(w, s.Success201, retVal)
}

// updateSocialMediaHandler
// Method: PUT
// Example: localhost/socialmedias/1
// JSON Body:
// {
// 	"name": "social media name",
// 	"social_media_url": "https://domainsocialmedia.com/user",
// 	"profile_image_url": "https://domainsocialmedia.com/userimage.jpg"
// }
func updateSocialMediaHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()

	if id != "" { // get by id
		if idInt, err := strconv.ParseInt(id, 10, 64); err == nil {
			validate := validator.New()
			decoder := json.NewDecoder(r.Body)
			var inp entity.SocialMediaPost
			if err := decoder.Decode(&inp); err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			err := validate.Struct(inp)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorBadRequest, err.Error())
				return
			}
			c, err := database.SqlDatabase.GetSocialMediaByID(ctx, idInt)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			if c.UserID != s.LogonUser.ID {
				s.WriteJsonResp(w, s.ErrorUnauthorized, "UNAUTHORIZED")
				return
			}

			p, err := database.SqlDatabase.UpdateSocialMedia(ctx, s.LogonUser.ID, idInt, inp)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			retVal := p.ToSocialMediaUpdateOutput()
			s.WriteJsonResp(w, s.Success, retVal)
		}
	}
}

// deleteSocialMediaHandler
// Method: DELETE
// Example: localhost/socialmedias/1
func deleteSocialMediaHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()
	if id != "" {
		if idInt, err := strconv.ParseInt(id, 10, 64); err == nil {
			c, err := database.SqlDatabase.GetSocialMediaByID(ctx, idInt)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			if c.UserID != s.LogonUser.ID {
				s.WriteJsonResp(w, s.ErrorUnauthorized, "UNAUTHORIZED")
				return
			}
			msg, err := database.SqlDatabase.DeleteComment(ctx, s.LogonUser.ID, idInt)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			retVal := map[string]string{
				"message": msg,
			}
			s.WriteJsonResp(w, s.Success, retVal)

		}
	}
}

package handler

import (
	"context"
	"encoding/json"
	"finalproject/database"
	"finalproject/entity"
	s "finalproject/server"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PhotoHandler struct{}

func InstallPhotosHandler(r *mux.Router) {
	api := PhotoHandler{}
	r.HandleFunc("/photos/{id}", api.PhotosHandler)
	r.HandleFunc("/photos", api.PhotosHandler)
}

type PhotoHandlerInterface interface {
	UsersHandler(w http.ResponseWriter, r *http.Request)
}

func NewPhotoHandler() PhotoHandlerInterface {
	return &UserHandler{}
}

func (h *PhotoHandler) PhotosHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodGet:
		if id == "" {
			getPhotosHandler(w, r)
		} else if id == "myphoto" {
			getMyPhotosHandler(w, r)
		} else {
			getPhotosByUserIDHandler(w, r, id)
		}
	case http.MethodPost:
		postPhotoHandler(w, r)
	case http.MethodPut:
		updatePhotoHandler(w, r, id)
	case http.MethodDelete:
		deletePhotoHandler(w, r, id)
	default:
		s.WriteJsonResp(w, s.ErrorNotFound, "PAGE NOT FOUND")
		return
	}
}

// getPhotosHandler
// Method: GET
// Example: localhost/photos
func getPhotosHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	ret, err := database.SqlDatabase.GetPhotos(ctx)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	retVal := []entity.PhotoGetOutput{}
	for _, i := range ret {
		var tempout entity.PhotoGetOutput
		tempout.Photo = i
		tempout.User = entity.UserUpdate{
			Email:    s.LogonUser.Email,
			Username: s.LogonUser.Username,
		}
		retVal = append(retVal, tempout)
	}

	s.WriteJsonResp(w, s.Success, retVal)
}

// getMyPhotosHandler
// Method: GET
// Example: localhost/photos
func getMyPhotosHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	ret, err := database.SqlDatabase.GetPhotosByUserID(ctx, s.LogonUser.ID)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	retVal := []entity.PhotoGetOutput{}
	for _, i := range ret {
		var tempout entity.PhotoGetOutput
		tempout.Photo = i
		tempout.User = entity.UserUpdate{
			Email:    s.LogonUser.Email,
			Username: s.LogonUser.Username,
		}
		retVal = append(retVal, tempout)
	}

	s.WriteJsonResp(w, s.Success, retVal)
}

// getPhotosByUserIDHandler
// Method: GET
// Example: localhost/photos/1
func getPhotosByUserIDHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	ret, err := database.SqlDatabase.GetPhotosByUserID(ctx, idInt)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	retVal := []entity.PhotoGetOutput{}
	for _, i := range ret {
		var tempout entity.PhotoGetOutput
		tempout.Photo = i
		tempout.User = entity.UserUpdate{
			Email:    s.LogonUser.Email,
			Username: s.LogonUser.Username,
		}
		retVal = append(retVal, tempout)
	}

	s.WriteJsonResp(w, s.Success, retVal)
}

// postPhotoHandler
// Method: POST
// Example: localhost/photos
// JSON Body:
// {
// 	"title": "title photo",
// 	"caption": "caption photo",
// 	"photo_url": "https://photo.domain.com"
// }
func postPhotoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	decoder := json.NewDecoder(r.Body)
	var inp entity.PhotoPost
	if err := decoder.Decode(&inp); err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	p, err := database.SqlDatabase.PostPhoto(ctx, s.LogonUser.ID, inp)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}

	retVal := p.ToPhotoPostOutput()
	s.WriteJsonResp(w, s.Success201, retVal)
}

// updatePhotoHandler
// Method: PUT
// Example: localhost/photos/1
// JSON Body:
// {
// 	"title": "title photo",
// 	"caption": "caption photo",
// 	"photo_url": "https://photo.domain.com"
// }
func updatePhotoHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()

	if id != "" { // get by id
		if idInt, err := strconv.ParseInt(id, 10, 64); err == nil {
			decoder := json.NewDecoder(r.Body)
			var inp entity.PhotoPost
			c, err := database.SqlDatabase.GetPhotoByID(ctx, idInt)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			if c.UserID != s.LogonUser.ID {
				s.WriteJsonResp(w, s.ErrorUnauthorized, "UNAUTHORIZED")
				return
			}
			if err := decoder.Decode(&inp); err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			p, err := database.SqlDatabase.UpdatePhoto(ctx, s.LogonUser.ID, idInt, inp.Title, inp.Caption, inp.PhotoUrl)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			retVal := p.ToPhotoUpdateOutput()
			s.WriteJsonResp(w, s.Success, retVal)
		}
	}
}

// deletePhotoHandler
// Method: DELETE
// Example: localhost/photos/1
func deletePhotoHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()
	if id != "" {
		if idInt, err := strconv.ParseInt(id, 10, 64); err == nil {
			c, err := database.SqlDatabase.GetPhotoByID(ctx, idInt)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			if c.UserID != s.LogonUser.ID {
				s.WriteJsonResp(w, s.ErrorUnauthorized, "UNAUTHORIZED")
				return
			}
			msg, err := database.SqlDatabase.DeletePhoto(ctx, s.LogonUser.ID, idInt)
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

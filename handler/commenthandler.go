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

type CommentHandler struct{}

func InstallCommentHandler(r *mux.Router) {
	api := CommentHandler{}
	r.HandleFunc("/comments/{id}", api.CommentsHandler)
	r.HandleFunc("/comments", api.CommentsHandler)
}

type CommentHandlerInterface interface {
	CommentsHandler(w http.ResponseWriter, r *http.Request)
}

func NewCommentHandler() CommentHandlerInterface {
	return &CommentHandler{}
}

func (h *CommentHandler) CommentsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodGet:
		if id != "" {
			getCommentsByPhotoIDHandler(w, r, id)
		} else {
			getCommentsHandler(w, r)
		}
	case http.MethodPost:
		postCommentHandler(w, r)
	case http.MethodPut:
		updateCommentHandler(w, r, id)
	case http.MethodDelete:
		deleteCommentHandler(w, r, id)
	default:
		s.WriteJsonResp(w, s.ErrorNotFound, "PAGE NOT FOUND")
		return
	}
}

// getCommentsHandler
// Method: GET
// Example: localhost/comments
func getCommentsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	retVal, err := database.SqlDatabase.GetComments(ctx)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}

	s.WriteJsonResp(w, s.Success, retVal)
}

// getCommentsHandler
// Method: GET
// Example: localhost/comments/1
func getCommentsByPhotoIDHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	c, err := database.SqlDatabase.GetCommentsByPhotoID(ctx, idInt)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}

	retVal := []entity.CommentGetOutput{}
	for _, i := range c {
		var tempout entity.CommentGetOutput
		p, err := database.SqlDatabase.GetPhotoByID(ctx, int64(i.PhotoID))
		if err != nil {
			s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
			return
		}
		tempout.Comment = i
		tempout.User = entity.UserGetComment{
			ID:       s.LogonUser.ID,
			Email:    s.LogonUser.Email,
			Username: s.LogonUser.Username,
		}
		tempout.Photo = *p.ToPhotoGetComment()
		retVal = append(retVal, tempout)
	}

	s.WriteJsonResp(w, s.Success, retVal)
}

// postCommentHandler
// Method: POST
// Example: localhost/comments
// JSON Body:
// {
// 	"message": "comment message",
// 	"photo_id": 1
// }
func postCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	validate := validator.New()
	decoder := json.NewDecoder(r.Body)
	var inp entity.CommentPost

	if err := decoder.Decode(&inp); err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}
	err := validate.Struct(inp)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorBadRequest, err.Error())
		return
	}
	c, err := database.SqlDatabase.PostComment(ctx, s.LogonUser.ID, inp)
	if err != nil {
		s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
		return
	}

	retVal := c.ToCommentPostOutput()

	s.WriteJsonResp(w, s.Success201, retVal)
}

// updateCommentHandler
// Method: PUT
// Example: localhost/comments/1
// JSON Body:
// {
// 	"message": "comment message"
// }
func updateCommentHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()

	if id != "" { // get by id
		if idInt, err := strconv.ParseInt(id, 10, 64); err == nil {
			validate := validator.New()
			decoder := json.NewDecoder(r.Body)
			var inp entity.CommentUpdate
			if err := decoder.Decode(&inp); err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			err := validate.Struct(inp)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorBadRequest, err.Error())
				return
			}
			c, err := database.SqlDatabase.GetCommentByID(ctx, idInt)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			if c.UserID != s.LogonUser.ID {
				s.WriteJsonResp(w, s.ErrorUnauthorized, "UNAUTHORIZED")
				return
			}

			p, err := database.SqlDatabase.UpdateComment(ctx, s.LogonUser.ID, idInt, inp.Message)
			if err != nil {
				s.WriteJsonResp(w, s.ErrorDataHandleError, err.Error())
				return
			}
			retVal := p.ToCommentUpdateOutput()
			s.WriteJsonResp(w, s.Success, retVal)
		}
	}
}

// deleteCommentHandler
// Method: DELETE
// Example: localhost/comments/1
func deleteCommentHandler(w http.ResponseWriter, r *http.Request, id string) {
	ctx := context.Background()
	if id != "" {
		if idInt, err := strconv.ParseInt(id, 10, 64); err == nil {
			c, err := database.SqlDatabase.GetCommentByID(ctx, idInt)
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

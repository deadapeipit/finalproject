package database

import (
	"context"
	"finalproject/entity"
	"reflect"
	"testing"
)

func TestDatabase_PostComment(t *testing.T) {
	type args struct {
		ctx    context.Context
		userid int64
		i      entity.CommentPost
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    *entity.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.PostComment(tt.args.ctx, tt.args.userid, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.PostComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.PostComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_GetComments(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    []entity.CommentGetOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetComments(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.GetComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_GetCommentsByPhotoID(t *testing.T) {
	type args struct {
		ctx     context.Context
		photoid int64
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    []entity.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetCommentsByPhotoID(tt.args.ctx, tt.args.photoid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetCommentsByPhotoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.GetCommentsByPhotoID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_GetCommentByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    *entity.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetCommentByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetCommentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.GetCommentByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_UpdateComment(t *testing.T) {
	type args struct {
		ctx     context.Context
		userid  int64
		id      int64
		message string
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    *entity.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateComment(tt.args.ctx, tt.args.userid, tt.args.id, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.UpdateComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.UpdateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_DeleteComment(t *testing.T) {
	type args struct {
		ctx    context.Context
		userid int64
		id     int64
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteComment(tt.args.ctx, tt.args.userid, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.DeleteComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Database.DeleteComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

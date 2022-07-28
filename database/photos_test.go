package database

import (
	"context"
	"finalproject/entity"
	"reflect"
	"testing"
)

func TestDatabase_PostPhoto(t *testing.T) {
	type args struct {
		ctx context.Context
		u   int64
		i   entity.PhotoPost
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    *entity.Photo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.PostPhoto(tt.args.ctx, tt.args.u, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.PostPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.PostPhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_GetPhotos(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    []entity.PhotoGetOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetPhotos(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetPhotos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.GetPhotos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_GetPhotosByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userid int64
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    []entity.Photo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetPhotosByUserID(tt.args.ctx, tt.args.userid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetPhotosByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.GetPhotosByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_GetPhotoByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    *entity.Photo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetPhotoByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetPhotoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.GetPhotoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_UpdatePhoto(t *testing.T) {
	type args struct {
		ctx      context.Context
		userid   int64
		id       int64
		title    string
		caption  string
		photourl string
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    *entity.Photo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdatePhoto(tt.args.ctx, tt.args.userid, tt.args.id, tt.args.title, tt.args.caption, tt.args.photourl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.UpdatePhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.UpdatePhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_DeletePhoto(t *testing.T) {
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
			got, err := tt.s.DeletePhoto(tt.args.ctx, tt.args.userid, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.DeletePhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Database.DeletePhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

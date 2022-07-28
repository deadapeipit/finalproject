package database

import (
	"context"
	"finalproject/entity"
	"reflect"
	"testing"
)

func TestDatabase_PostSocialMedia(t *testing.T) {
	type args struct {
		ctx    context.Context
		userid int64
		i      entity.SocialMediaPost
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    *entity.SocialMedia
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.PostSocialMedia(tt.args.ctx, tt.args.userid, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.PostSocialMedia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.PostSocialMedia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_GetSocialMedias(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    []entity.SocialMediaGetOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetSocialMedias(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetSocialMedias() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.GetSocialMedias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_GetSocialMediaByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    *entity.SocialMedia
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetSocialMediaByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetSocialMediaByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.GetSocialMediaByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_UpdateSocialMedia(t *testing.T) {
	type args struct {
		ctx    context.Context
		userid int64
		id     int64
		i      entity.SocialMediaPost
	}
	tests := []struct {
		name    string
		s       *Database
		args    args
		want    *entity.SocialMedia
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateSocialMedia(tt.args.ctx, tt.args.userid, tt.args.id, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.UpdateSocialMedia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.UpdateSocialMedia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_DeleteSocialMedia(t *testing.T) {
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
			got, err := tt.s.DeleteSocialMedia(tt.args.ctx, tt.args.userid, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.DeleteSocialMedia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Database.DeleteSocialMedia() = %v, want %v", got, tt.want)
			}
		})
	}
}

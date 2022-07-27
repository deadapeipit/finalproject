package database

import (
	"context"
	"database/sql"
	"finalproject/entity"
	"time"
)

func (s *Database) PostSocialMedia(ctx context.Context, userid int64, i entity.SocialMediaPost) (*entity.SocialMedia, error) {
	result := &entity.SocialMedia{}
	qry := "insert into socialmedias (name, socialmediaurl, userid, createdat, updatedat) values (@name, @socialmediaurl, @userid, @createdat, @updatedat); select id, name, socialmediaurl, userid, createdat,updatedat from socialmedias"
	now := time.Now()
	rows, err := s.SqlDb.QueryContext(ctx, qry,
		sql.Named("name", i.Name),
		sql.Named("socialmediaurl", i.SocialMediaUrl),
		sql.Named("userid", userid),
		sql.Named("createdat", now),
		sql.Named("updatedat", now))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.SocialMediaUrl,
			&result.UserID,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (s *Database) GetSocialMedias(ctx context.Context, userid int64) ([]entity.SocialMedia, error) {
	var result []entity.SocialMedia
	qry := "select id, name, socialmediaurl, userid, createdat, updatedat from socialmedias where userid=@userid"
	rows, err := s.SqlDb.QueryContext(ctx, qry,
		sql.Named("userid", userid))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row entity.SocialMedia
		err := rows.Scan(
			&row.ID,
			&row.Name,
			&row.SocialMediaUrl,
			&row.UserID,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (s *Database) GetSocialMediaByID(ctx context.Context, id int64) (*entity.SocialMedia, error) {
	result := &entity.SocialMedia{}

	rows, err := s.SqlDb.QueryContext(ctx, "sselect id, name, socialmediaurl, userid, createdat, updatedat from socialmedias where id = @ID",
		sql.Named("ID", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.SocialMediaUrl,
			&result.UserID,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s *Database) UpdateSocialMedia(ctx context.Context, userid int64, id int64, name string, socialmediaurl string) (*entity.SocialMedia, error) {
	result := &entity.SocialMedia{}
	now := time.Now()
	qry := "update socialmedias set name=@name, socialmediaurl=@socialmediaurl, updatedat=@updatedat where id = @ID; select id, name, socialmediaurl, userid, updatedat from socialmedias where id = @ID"
	rows, err := s.SqlDb.QueryContext(ctx, qry,
		sql.Named("name", name),
		sql.Named("socialmediaurl", socialmediaurl),
		sql.Named("updatedat", now),
		sql.Named("userid", userid),
		sql.Named("ID", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.SocialMediaUrl,
			&result.UserID,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s *Database) DeleteSocialMedia(ctx context.Context, userid int64, id int64) (string, error) {
	var result string
	qry := "delete from socialmedias where id=@id and userid=@userid"
	_, err := s.SqlDb.ExecContext(ctx, qry,
		sql.Named("userid", userid),
		sql.Named("id", id))
	if err != nil {
		return "", err
	}

	result = "Your social media has been successfully deleted"

	return result, nil
}

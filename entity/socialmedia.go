package entity

import "time"

type SocialMedia struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int64     `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaPost struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type SocialMediaPostOutput struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int64     `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

func (s *SocialMedia) ToSocialMediaPostOutput() *SocialMediaPostOutput {
	out := &SocialMediaPostOutput{
		ID:             s.ID,
		Name:           s.Name,
		SocialMediaUrl: s.SocialMediaUrl,
		UserID:         s.UserID,
		CreatedAt:      s.CreatedAt,
	}
	return out
}

type SocialMediaUpdateOutput struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int64     `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (s *SocialMedia) ToSocialMediaUpdateOutput() *SocialMediaUpdateOutput {
	out := &SocialMediaUpdateOutput{
		ID:             s.ID,
		Name:           s.Name,
		SocialMediaUrl: s.SocialMediaUrl,
		UserID:         s.UserID,
		UpdatedAt:      s.UpdatedAt,
	}
	return out
}

type SocialMediaGetOutput struct {
	SocialMedia
	User UserGetComment `json:"user"`
}

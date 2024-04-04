package domain

import "context"

type Newsletter struct {
	Type        string `json:"type" gorm:"column:type" validate:"required"`
	Title       string `json:"title" gorm:"column:title" validate:"required"`
	SubTitle    string `json:"subtitle" gorm:"column:subtitle" validate:"required"`
	Description string `json:"description" gorm:"column:description" validate:"required"`
	Image       string `json:"image" gorm:"column:image" validate:"required"`
	TextColor   string `json:"textcolor" gorm:"column:textcolor" validate:"required"`
}

type NewsletterResponse struct {
	Id string `json:"id" gorm:"column:id" validate:"required"`
	Newsletter
}

type INewsletterRepository interface {
	Insert(ctx context.Context, newsletter Newsletter) error
}

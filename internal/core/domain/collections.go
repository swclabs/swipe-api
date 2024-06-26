package domain

// Collection use to query data from database, define in database, used to save a CollectionBody
type Collection struct {
	Id       int64  `json:"id" db:"collection"`
	Created  string `json:"created" db:"created"`
	Position string `json:"position" db:"position"` // Example: mac#1, mac#2
	Headline string `json:"headline" db:"headline"` // Ex: Get to know Mac
	Body     string `json:"body" db:"body"`
}

/***********************************************************************************************/

// CollectionBody used to bind from json mapping to structure
type CollectionBody struct {
	Title       string `json:"title" validate:"required"`
	SubTitle    string `json:"subtitle" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image"`
	TextColor   string `json:"text_color" validate:"required"`
}

// CollectionSliceBody is a body of CardBannerSlice
type CollectionSliceBody struct {
	CollectionBody
	Id      int64  `json:"id"`
	Created string `json:"created"`
}

// CollectionSchema is a type use to accept request and response
type CollectionSchema struct {
	Id       int64          `json:"id"`
	Created  string         `json:"created"`
	Position string         `json:"position" validate:"required"`
	Headline string         `json:"headline" validate:"required"`
	Body     CollectionBody `json:"body" validate:"required"`
}

// CollectionSliceSchema is a type use to accept request and response
type CollectionSliceSchema struct {
	Position   string                `json:"position"`
	Headline   string                `json:"headline"`
	CardBanner []CollectionSliceBody `json:"card_banner"`
}

type CollectionUploadSchema struct {
	Msg string `json:"msg"`
	Id  int64  `json:"id"`
}

/*** Swagger ***/

type CollectionSchemaSwagger struct {
	Position string                `json:"position" validate:"required"`
	Headline string                `json:"headline" validate:"required"`
	Body     CollectionBodySwagger `json:"body" validate:"required"`
}

type CollectionBodySwagger struct {
	Title       string `json:"title" validate:"required"`
	SubTitle    string `json:"subtitle" validate:"required"`
	Description string `json:"description" validate:"required"`
	TextColor   string `json:"text_color" validate:"required"`
}

/***********************************************************************************************/

type HeadlineBannerBody struct {
	Headline string `json:"headline" validate:"required"`
}

type HeadlineBannerSliceBody struct {
	HeadlineBannerBody
	Id      int64  `json:"id"`
	Created string `json:"created"`
}

// HeadlineBannerSchema user body request & response
type HeadlineBannerSchema struct {
	Position string             `json:"position" validate:"required"`
	Created  string             `json:"created"`
	Body     HeadlineBannerBody `json:"body" validate:"required"`
}

// HeadlineBannerSchemaSwagger used to generate swagger documents
type HeadlineBannerSchemaSwagger struct {
	Position string             `json:"position" validate:"required"`
	Body     HeadlineBannerBody `json:"body" validate:"required"`
}

// HeadlineBannerSlice response slices
type HeadlineBannerSlice struct {
	Position  string                    `json:"position"`
	Headlines []HeadlineBannerSliceBody `json:"headlines"`
}

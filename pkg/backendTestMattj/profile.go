package backendTestMattj

import (
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type (
	Profile struct {
		ID                   *uuid.UUID       `json:"id,omitempty"`
		FirstName            string           `json:"first_name,omitempty"`
		LastName             string           `json:"last_name,omitempty"`
		Email                string           `json:"email,omitempty"`
		DefaultLanguage      *DefaultLanguage `json:"default_language,omitempty"`
		StreetAddPrimary     string           `json:"street_address_primary,omitempty"`
		StreetAddSecondary   string           `json:"street_address_secondary,omitempty"`
		PhoneNumberPrimary   *PhoneNumber     `json:"phone_number_primary,omitempty"`
		PhoneNumberSecondary *PhoneNumber     `json:"phone_number_secondary,omitempty"`
		City                 string           `json:"city,omitempty"`
		StateProvince        string           `json:"state_province,omitempty"`
		ZipCode              string           `json:"zip_code,omitempty"`
		CreatedAt            *time.Time       `json:"created_at,omitempty" `
		UpdatedAt            *time.Time       `json:"updated_at,omitempty" `
	}

	DefaultLanguage struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Code string    `json:"code"`
	}

	PhoneNumber struct {
		Extension    string `json:"ext"`
		Country_code string `json:"country_code"`
		PhoneNumber  string `json:"phone_number,omitempty"`
	}
)

func (p *Profile) Validate() error {
	return validation.ValidateStruct(p, validation.Field(&p.FirstName, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.LastName, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.DefaultLanguage, validation.Required),
		validation.Field(&p.StreetAddPrimary, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.PhoneNumberPrimary, validation.Required),
		validation.Field(&p.City, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.StateProvince, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.ZipCode, validation.Required, validation.Length(1, 100)))
}

func (p *Profile) GetProfile(ctx *gin.Context) error {

	return nil
}

func (p *Profile) CreateProfile() error {
	return nil
}

func (p *Profile) UpdateProfile() error {
	return nil
}

func (p *Profile) DeleteProfile() error {
	return nil
}

package backendTestMattj

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type (
	Profile struct {
		ID                   uuid.UUID        `json:"id"`
		FirdtName            string           `json:"first_name"`
		LastName             string           `json:"last_name"`
		Email                string           `json:"email"`
		DefaultLanguage      *DefaultLanguage `json:"default_language,omitempty"`
		StreetAddPrimary     string           `json:"street_address_primary"`
		StreetAddSecondary   string           `json:"street_address_secondary,omitempty"`
		PhoneNumberPrimary   *PhoneNumber     `json:"phone_number_primary"`
		PhoneNumberSecondary *PhoneNumber     `json:"phone_number_secondary,omitempty"`
		City                 string           `json:"city"`
		StateProvince        string           `json:"state_province"`
		ZipCode              string           `json:"zip_code"`
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
	return validation.ValidateStruct(p,validation.Field(&p.FirstName, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.LastName, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.DefaultLanguage, validation.Required),
		validation.Field(&p.StreetAddPrimary, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.PhoneNumberPrimary, validation.Required),
		validation.Field(&p.City, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.StateProvince, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.ZipCode, validation.Required, validation.Length(1, 100))

}

func (p *Profile) GetProfile() error {
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





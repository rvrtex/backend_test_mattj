package backendTestMattj

import (
	"fmt"

	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	sq "github.com/Masterminds/squirrel"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "appliedsystems"
)

type (
	Profile struct {
		ID                   *uuid.UUID `json:"id,omitempty" db:"id"`
		FirstName            string     `json:"first_name,omitempty" db:"first_name"`
		LastName             string     `json:"last_name,omitempty" db:"last_name"`
		Email                string     `json:"email,omitempty" db:"email"`
		DefaultLanguage      *uuid.UUID `json:"default_language,omitempty" db:"default_language"`
		StreetAddPrimary     string     `json:"street_address_primary,omitempty" db:"street_address_primary"`
		StreetAddSecondary   string     `json:"street_address_secondary,omitempty" db:"street_address_secondary"`
		PhoneNumberPrimary   *uuid.UUID `json:"phone_number_primary,omitempty" db:"phone_number_primary"`
		PhoneNumberSecondary *uuid.UUID `json:"phone_number_secondary,omitempty" db:"phone_number_secondary"`
		City                 string     `json:"city,omitempty" db:"city"`
		StateProvidence      string     `json:"state_providence,omitempty" db:"state_providence"`
		ZipCode              string     `json:"zip_code,omitempty" db:"zip_code"`
		CreatedAt            *time.Time `json:"created_at,omitempty" db:"created_at"`
		UpdatedAt            *time.Time `json:"updated_at,omitempty" db:"updated_at"`
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

func (p *Profile) GetProfile(ctx *gin.Context) (*Profile, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbCon, err := sqlx.Connect("pgx", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer dbCon.Close()

	if err := dbCon.Ping(); err != nil {
		panic(err)
	}

	singleProfile := Profile{}

	query := sq.Select("*").
		PlaceholderFormat(sq.Dollar).
		From("profile").
		Where(sq.Eq{
			"id": ctx.Param("id"),
		})

	stmt, args, _ := query.ToSql()
	row := dbCon.QueryRowx(stmt, args...)
	err = row.StructScan(&singleProfile)
	if err != nil {
		return nil, err
	}
	return &singleProfile, nil
}

func (p *Profile) Validate() error {
	return validation.ValidateStruct(p, validation.Field(&p.FirstName, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.LastName, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.DefaultLanguage, validation.Required),
		validation.Field(&p.StreetAddPrimary, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.PhoneNumberPrimary, validation.Required),
		validation.Field(&p.City, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.StateProvidence, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.ZipCode, validation.Required, validation.Length(1, 100)))
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

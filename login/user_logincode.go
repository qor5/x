package login

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLoginCoder interface {
	FindUserByPhoneNumber(db *gorm.DB, model interface{}, phoneNumber string) (user interface{}, err error)
	GenerateLoginCode(db *gorm.DB, model interface{}) (token string, err error)
	ConsumeLoginCode(db *gorm.DB, model interface{}) error
	GetLoginCode() (token string, createdAt *time.Time, expired bool)
	SetConfirmTime(db *gorm.DB, model interface{}) error
}

type UserLoginCodeSender interface {
	SendLoginCode(r *http.Request, phoneNumber string, code string) error
}

type UserLoginCode struct {
	PhoneNumber        string `gorm:"index:,unique,where:phone_number!='' and deleted_at is null"`
	ConfirmedAt        *time.Time
	LoginCode          string `gorm:"index:,unique,where:login_code!=''"`
	LoginCreatedAt     *time.Time
	LoginCodeExpiredAt *time.Time
}

var _ UserLoginCoder = (*UserLoginCode)(nil)

const minPhoneNumberLength = 8

func numbersOnly(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsNumber(r) {
			return r
		}
		return -1
	}, s)
}

func (up *UserLoginCode) FindUserByPhoneNumber(db *gorm.DB, model interface{}, phoneNumber string) (user interface{}, err error) {
	phoneNumber = numbersOnly(phoneNumber)
	if len(phoneNumber) < minPhoneNumberLength {
		return nil, ErrAccountNumberInvalid
	}
	// Logic here is to try to find the phone number in the database, even
	// if the user did not enter the international code. So we use a LIKE
	// to check if there are more than one user with the same ending digits.
	// If there are, we need more numbers. Otherwise, we can assume it's the correct user.
	result := db.Model(model).Where("phone_number like ?", fmt.Sprintf("%%%s", phoneNumber)).
		First(model)
	if result.Error != nil {
		return nil, result.Error
	}
	switch result.RowsAffected {
	case 0:
		return nil, gorm.ErrRecordNotFound
	case 1:
		return model, nil
	default:
		return nil, ErrAccountNumberInvalid
	}
}

func (up *UserLoginCode) GetPhoneNumber() string {
	return up.PhoneNumber
}

func (up *UserLoginCode) GenerateLoginCodeExpiration(db *gorm.DB) (createdAt time.Time, expiredAt time.Time) {
	createdAt = db.NowFunc()
	return createdAt, createdAt.Add(10 * time.Minute)
}

func (up *UserLoginCode) GenerateLoginCode(db *gorm.DB, model interface{}) (code string, err error) {
	code = fmt.Sprintf("%06d", uuid.New().ID()%1000000)

	iface, ok := model.(interface {
		GenerateLoginCodeExpiration(db *gorm.DB) (createdAt time.Time, expiredAt time.Time)
	})
	if !ok {
		return "", fmt.Errorf("model does not have GenerateLoginCodeExpiration method, maybe it does not embed UserWhatsApp")
	}

	createdAt, expiredAt := iface.GenerateLoginCodeExpiration(db)

	result := db.Model(model).
		Where("phone_number = ?", numbersOnly(up.PhoneNumber)).
		Updates(map[string]interface{}{
			"login_code":            code,
			"login_created_at":      createdAt,
			"login_code_expired_at": expiredAt,
		})
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected != 1 {
		return "", gorm.ErrRecordNotFound
	}
	up.LoginCode = code
	up.LoginCreatedAt = &createdAt
	up.LoginCodeExpiredAt = &expiredAt
	return code, nil
}

func (up *UserLoginCode) ConsumeLoginCode(db *gorm.DB, model interface{}) error {
	now := time.Now()
	err := db.Model(model).
		Where("phone_number = ?", numbersOnly(up.PhoneNumber)).
		Updates(map[string]interface{}{
			"login_code_expired_at": now,
			"login_code":            "",
		}).
		Error
	if err != nil {
		return err
	}
	up.LoginCode = ""
	up.LoginCodeExpiredAt = &now
	return nil
}

func (up *UserLoginCode) GetLoginCode() (token string, createdAt *time.Time, expired bool) {
	if up.LoginCodeExpiredAt != nil && time.Since(*up.LoginCodeExpiredAt) > 0 {
		return "", nil, true
	}
	return up.LoginCode, up.LoginCreatedAt, false
}

func (up *UserLoginCode) SetConfirmTime(db *gorm.DB, model interface{}) error {
	now := time.Now()
	result := db.Model(model).
		Where("phone_number = ?", numbersOnly(up.PhoneNumber)).
		Update("confirmed_at", now)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return gorm.ErrRecordNotFound
	}

	up.ConfirmedAt = &now
	return nil
}

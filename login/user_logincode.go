package login

import (
	"encoding/base64"
	"fmt"
	"time"

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
	SendLoginCode(phoneNumber string, code string) error
}

type UserLoginCode struct {
	PhoneNumber            string `gorm:"index:,unique,where:phone_number!='' and deleted_at is null"`
	ConfirmedAt         *time.Time
	LoginToken          string `gorm:"index:,unique,where:login_token!=''"`
	LoginCreatedAt      *time.Time
	LoginTokenExpiredAt *time.Time
}

//var _ UserLoginCoder = (*UserLoginCode)(nil)

func (up *UserLoginCode) FindUserByPhoneNumber(db *gorm.DB, model interface{}, phoneNumber string) (user interface{}, err error) {
	err = db.Where("phone_number = ?", phoneNumber).
		First(model).
		Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (up *UserLoginCode) GetPhoneNumber() string {
	return up.PhoneNumber
}

func (up *UserLoginCode) GenerateLoginCodeExpiration(db *gorm.DB) (createdAt time.Time, expiredAt time.Time) {
	createdAt = db.NowFunc()
	return createdAt, createdAt.Add(10 * time.Minute)
}

func (up *UserLoginCode) GenerateLoginCode(db *gorm.DB, model interface{}) (token string, err error) {
	token = base64.URLEncoding.EncodeToString([]byte(uuid.NewString()))

	iface, ok := model.(interface {
		GenerateLoginCodeExpiration(db *gorm.DB) (createdAt time.Time, expiredAt time.Time)
	})
	if !ok {
		return "", fmt.Errorf("model does not have GenerateLoginCodeExpiration method, maybe it does not embed UserWhatsApp")
	}

	createdAt, expiredAt := iface.GenerateLoginCodeExpiration(db)

	err = db.Model(model).
		Where("login_code = ?", up.PhoneNumber).
		Updates(map[string]interface{}{
			"login_token":            token,
			"login_created_at":       createdAt,
			"login_token_expired_at": expiredAt,
		}).
		Error
	if err != nil {
		return "", err
	}
	up.LoginToken = token
	up.LoginCreatedAt = &createdAt
	up.LoginTokenExpiredAt = &expiredAt
	return token, nil
}

func (up *UserLoginCode) ConsumeLoginCode(db *gorm.DB, model interface{}) error {
	err := db.Model(model).
		Where("login_code = ?", up.PhoneNumber).
		Updates(map[string]interface{}{
			"login_token_expired_at": time.Now(),
		}).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (up *UserLoginCode) GetLoginCode() (token string, createdAt *time.Time, expired bool) {
	if up.LoginTokenExpiredAt != nil && time.Since(*up.LoginTokenExpiredAt) > 0 {
		return "", nil, true
	}
	return up.LoginToken, up.LoginCreatedAt, false
}

func (up *UserLoginCode) SetConfirmTime(db *gorm.DB, model interface{}) error {
	now := time.Now()
	err := db.Model(model).
		Where("login_code = ?", up.PhoneNumber).
		Update("confirmed_at", now).
		Error
	if err != nil {
		return err
	}
	up.ConfirmedAt = &now
	return nil
}

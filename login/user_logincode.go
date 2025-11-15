package login

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLoginCoder interface {
	FindUser(db *gorm.DB, model interface{}, identifier string) (user interface{}, err error)
	GenerateLoginCode(db *gorm.DB, model interface{}) (token string, err error)
	ConsumeLoginCode(db *gorm.DB, model interface{}) error
	GetLoginCode() (token string, createdAt *time.Time, expired bool)
	SetConfirmTime(db *gorm.DB, model interface{}) error
}

type UserLoginCodeSender interface {
	SendLoginCode(r *http.Request, identifier string, code string) error
}

type UserLoginCode struct {
	ConfirmedAt        *time.Time
	LoginCode          string `gorm:"index:,unique,where:login_code!=''"`
	LoginCreatedAt     *time.Time
	LoginCodeExpiredAt *time.Time
}

var _ UserLoginCoder = (*UserLoginCode)(nil)

func (up *UserLoginCode) FindUser(db *gorm.DB, model interface{}, identifier string) (user interface{}, err error) {
	// This is a generic finder, the actual implementation is on the user model.
	iface, ok := model.(UserLoginCoder)
	if !ok {
		return nil, fmt.Errorf("model does not implement UserLoginCoder")
	}
	return iface.FindUser(db, model, identifier)
}

func (up *UserLoginCode) GenerateLoginCodeExpiration(db *gorm.DB) (createdAt time.Time, expiredAt time.Time) {
	createdAt = db.NowFunc()
	return createdAt, createdAt.Add(10 * time.Minute)
}

func getModelPrimaryKey(db *gorm.DB, model interface{}) (pk string, pv any) {
	// get the primary key of the model
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(model)
	primaryField := stmt.Schema.PrioritizedPrimaryField
	primaryValue := primaryField.ReflectValueOf(stmt.Context, reflect.ValueOf(model)).Interface()
	return primaryField.DBName, primaryValue
}

func (up *UserLoginCode) GenerateLoginCode(db *gorm.DB, model interface{}) (code string, err error) {
	code = fmt.Sprintf("%06d", uuid.New().ID()%1000000)

	iface, ok := model.(interface {
		GenerateLoginCodeExpiration(db *gorm.DB) (createdAt time.Time, expiredAt time.Time)
	})
	if !ok {
		return "", fmt.Errorf("model does not have GenerateLoginCodeExpiration method, maybe it does not embed LoginCoder")
	}

	createdAt, expiredAt := iface.GenerateLoginCodeExpiration(db)


	pk, pv := getModelPrimaryKey(db, model)

	result := db.Model(model).
		Where(fmt.Sprintf("%s = ?", pk), pv).
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
	pk, pv := getModelPrimaryKey(db, model)

	err := db.Model(model).
		Where(fmt.Sprintf("%s = ?", pk), pv).
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
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(model)
	primaryField := stmt.Schema.PrioritizedPrimaryField
	primaryValue := primaryField.ReflectValueOf(stmt.Context, reflect.ValueOf(model)).Interface()
	result := db.Model(model).
		Where(fmt.Sprintf("%s = ?", primaryField.DBName), primaryValue).
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

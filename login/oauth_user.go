package login

import "gorm.io/gorm"

type OAuthUser interface {
	FindUserByOAuthUserID(db *gorm.DB, model interface{}, provider string, oid string) (user interface{}, err error)
	FindUserByOAuthIdentifier(db *gorm.DB, model interface{}, provider string, identifier string) (user interface{}, err error)
	// only update the o_auth_user_id when it's empty(null or '')
	InitOAuthUserID(db *gorm.DB, model interface{}, provider string, identifier string, oid string) error
	SetAvatar(v string)
	GetAvatar() string
}

type OAuthInfo struct {
	OAuthProvider string `gorm:"index:,unique,composite:oauth_provider_oauth_user_id,where:o_auth_provider!='' and o_auth_user_id!='' and deleted_at is null;index:,unique,composite:oauth_provider_oauth_identifier,where:o_auth_provider!='' and o_auth_identifier!='' and deleted_at is null"`
	OAuthUserID   string `gorm:"index:,unique,composite:oauth_provider_oauth_user_id,where:o_auth_provider!='' and o_auth_user_id!='' and deleted_at is null"`
	// OAuthIdentifier is an externally-facing account identifier, such as an email address for a Google account.
	// it is used to find the user record on the first login
	OAuthIdentifier string `gorm:"index:,unique,composite:oauth_provider_oauth_identifier,where:o_auth_provider!='' and o_auth_identifier!='' and deleted_at is null"`
	OAuthAvatar     string `gorm:"-"`
}

var _ OAuthUser = (*OAuthInfo)(nil)

func (oa *OAuthInfo) FindUserByOAuthUserID(db *gorm.DB, model interface{}, provider string, oid string) (user interface{}, err error) {
	err = db.Where("o_auth_provider = ? and o_auth_user_id = ?", provider, oid).
		First(model).
		Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (oa *OAuthInfo) FindUserByOAuthIdentifier(db *gorm.DB, model interface{}, provider string, identifier string) (user interface{}, err error) {
	err = db.Where("o_auth_provider = ? and o_auth_identifier = ?", provider, identifier).
		First(model).
		Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (oa *OAuthInfo) InitOAuthUserID(db *gorm.DB, model interface{}, provider string, identifier string, oid string) error {
	err := db.Model(model).
		Where("o_auth_provider = ? and o_auth_identifier = ? and coalesce(o_auth_user_id, '') = ''", provider, identifier).
		Updates(map[string]interface{}{
			"o_auth_user_id": oid,
		}).
		Error
	if err != nil {
		return err
	}
	oa.OAuthUserID = oid
	return nil
}

func (oa *OAuthInfo) SetAvatar(v string) {
	oa.OAuthAvatar = v
}

func (oa *OAuthInfo) GetAvatar() string {
	return oa.OAuthAvatar
}

package messages

import "github.com/qor5/x/v3/i18n/i18n-transfer/test/mock/out_messages"

const I18nModuleKey = "i18nModuleKey"

type UserMessage struct {
	name string
	Email
	out_messages.Detail
}

type Email struct {
	Email string
	Phone
}

var User_CN = &UserMessage{
	name:   "New User CN",
	Email:  Email_CN,
	Detail: out_messages.Detail_CN,
}

var Email_CN = Email{
	Email: "New CNEmail",
	Phone: Phone_CN,
}

var User_JP = &UserMessage{
	name:   "New User JP",
	Email:  Email_JP,
	Detail: out_messages.Detail_JP,
}

var Email_JP = Email{
	Email: "New JPEmail",
	Phone: Phone_JP,
}

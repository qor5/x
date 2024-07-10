package messages

const I18nModuleKey = "i18nModuleKey"

type UserMessage struct {
	name string
	Email
}

type Email struct {
	Email string
	Phone
}

type Phone struct {
	PhoneNumber string
}

var User_CN = &UserMessage{
	name:  "User CN",
	Email: Email_CN,
}

var Email_CN = Email{
	Email: "terry@theplant.cn",
	Phone: Phone_CN,
}

var Phone_CN = Phone{
	PhoneNumber: "+86",
}

var User_JP = &UserMessage{
	name:  "User JP",
	Email: Email_JP,
}

var Email_JP = Email{
	Email: "terry@theplant.jp",
	Phone: Phone_JP,
}

var Phone_JP = Phone{
	PhoneNumber: "+100",
}

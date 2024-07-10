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

var User_CN = &UserMessage{
	name:  "User CN",
	Email: Email_CN,
}

var Email_CN = Email{
	Email: "terry@theplant.cn",
	Phone: Phone_CN,
}

var User_JP = &UserMessage{
	name:  "User JP",
	Email: Email_JP,
}

var Email_JP = Email{
	Email: "terry@theplant.jp",
	Phone: Phone_JP,
}

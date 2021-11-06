package web_model

type WebsiteModel struct {
	Id int64
	WebsiteTitle string `xorm:"varchar(80)" json:"website_title"`
	WebsiteUrl string `xorm:"varchar(255)" json:"website_url"`
	WebsiteManager string `xorm:"varchar(30)" json:"website_manager"`
	WebsiteNickname string `xorm:"varchar(50)" json:"website_nickname"`
	WebsiteBanner string `xorm:"varchar(255)" json:"website_banner"`
	WebsiteAvatar string `xorm:"varchar(255)" json:"website_avatar"`
	WebsiteDesc string `xorm:"varchar(255)" json:"website_desc"`
	WebsiteEmail string `xorm:"varchar(120)" json:"website_email"`
	WebsitePhone string `xorm:"varchar(11)" json:"website_phone"`
	WebsiteWechar string `xorm:"varchar(255)" json:"website_wechar"`
	WebsiteWebo string `xorm:"varchar(255)" json:"website_webo"`
	WebsiteQq string `xorm:"varchar(20)" json:"website_qq"`
	WebsiteCopy string `xorm:"varchar(255)" json:"website_copy"`
	WebsiteLogo string `xorm:"varchar(255)" json:"website_logo"`
	WebsiteStaticUrl string `xorm:"varchar(255)" json:"website_static_url"`
	WebsiteLicenceOne string `xorm:"varchar(255)" json:"website_licence_one"`
	WebsiteLicenceTwo string `xorm:"varchar(255)" json:"website_licence_two"`
	WebsiteLicenceThree string `xorm:"varchar(255)" json:"website_licence_three"`
	WebsiteLicenceFour string `xorm:"varchar(255)" json:"website_licence_Four"`
	WebsiteApk string `xorm:"varchar(255)" json:"website_apk"`
	WebsiteIos string `xorm:"varchar(255)" json:"website_ios"`
	WebsiteVersion string `xorm:"varchar(50)" json:"website_version"`
	WebsiteNotice string `xorm:"varchar(255)" json:"website_notice"`
	WebsiteKeyword string `xorm:"varchar(255)" json:"website_keyword"`
	WebsiteAddress string `xorm:"varchar(255)" json:"website_address"`
	WebsiteAgreement string `xorm:"text" json:"website_agreement"`
	WebsitePrivacy string `xorm:"text" json:"website_privacy"`
	WebsiteAbout string `xorm:"text" json:"website_about"`
}

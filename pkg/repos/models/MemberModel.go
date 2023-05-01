package models

const (
	MEMBER_GENDER_MALE   = "MALE"
	MEMBER_GENDER_FEMALE = "FEMALE"
	MEMBER_ONLINE        = 1
	MEMBER_OFFLINE       = 0
)

type MemberModel struct {
	Id         int    `json:"id"`
	MemberId   string `gorm:"type:varchar(36);not null;index:member_id"`
	Name       string `gorm:"type:varchar(100);not null"`
	Email      string `gorm:"type:varchar(50);not null;index:email"`
	Gender     string `gorm:"type:varchar(10);not null"`
	Password   string `gorm:"type:varchar(255);not null"`
	IsOnline  int    `gorm:"type:tinyint;not null"`
	DeleteTime int    `gorm:"type:int(10);default:null"`
	CreateTime int    `gorm:"type:int(10);default:null"`
}

func (MemberModel) TableName() string {
	return "member"
}

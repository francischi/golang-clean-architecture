package models

import h "golang/pkg/helpers"

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

func (m *MemberModel) TableName() string {
	return "member"
}

func (m *MemberModel)Check()(bool,string){
	if !h.CheckUuid(m.MemberId){
		return false,"invalid uuid"
	}
	if len(m.Name) == 0{
		return false,"invalid name"
	}
	if m.Gender != MEMBER_GENDER_MALE && m.Gender != MEMBER_GENDER_FEMALE {
		return false,"invalid gender"
	}
	if !h.IsValidEmail(m.Email){
		return false,"invalid email"
	}
	return true ,""
}
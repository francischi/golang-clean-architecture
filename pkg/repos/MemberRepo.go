package repos

import (
	// "fmt"
	"gorm.io/gorm"
	"golang/pkg/repos/models"
)

type MemberRepo struct {
	DBconn *gorm.DB
}

func NewMemberRepo(db *gorm.DB) *MemberRepo{
	return &MemberRepo{DBconn:db}
}

func (m *MemberRepo) Create(member models.MemberModel)(err error){
	error := m.DBconn.Create(&member).Error
	return error
}

func (m *MemberRepo) GetMember(memberId string)(member models.MemberModel , err error){
	var model models.MemberModel
	error := m.DBconn.First(&model , "member_id=?",memberId).Error
	return model , error
}

func (m *MemberRepo) GetMemberByEmail(email string)(member models.MemberModel , err error){
	var model models.MemberModel
	error := m.DBconn.First(&model , "email=?",email).Error
	return model,error
}

func (m *MemberRepo ) ChangePwd(memberId string , newPwd string)(err error){
	var model models.MemberModel
	error := m.DBconn.First(&model , "member_id=?",memberId).Update("password", newPwd).Error
	return error
}

func (m *MemberRepo) IsEmailExist(email string) bool {
	var model models.MemberModel
	var count int64

	m.DBconn.Model(&model).Where("email", email).Count(&count)
	if( count != 0 ){
		return true
	}else{
		return false
	}
}
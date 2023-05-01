package interfaces

import(
	"golang/pkg/repos/models"
)

type MemberRepo interface {
	
	Create(member models.MemberModel)(err error)

	GetMember(memberId string)(member models.MemberModel , err error)

	GetMemberByEmail(email string)(member models.MemberModel , err error)

	ChangePwd(memberId string , newPwd string)(err error)

	IsEmailExist(email string) bool
}
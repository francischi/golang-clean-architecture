//+build wireinject

package memberModule

import (
	"github.com/google/wire"
	"golang/pkg/repos"
	"golang/pkg/helpers"
)

func InitMemberController() *MemberController{
	wire.Build(NewMemberController , NewMemberService , repos.NewMemberRepo , helpers.NewSqlSession)
	return &MemberController{}
}
//+build wireinject

package memberModule

import (
	"github.com/google/wire"
	impl "golang/pkg/repos/implement"
	"golang/pkg/helpers"
)

func InitMemberController() *MemberController{
	wire.Build(NewMemberController , NewMemberService , impl.NewMemberRepo , helpers.NewSqlSession)
	return &MemberController{}
}
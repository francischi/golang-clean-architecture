//+build wireinject

package memberModule

import (
	"golang/pkg/helpers"
	"golang/pkg/repos/implement"
	impl "golang/pkg/repos/implement"
	"golang/pkg/repos/interfaces"
	"golang/pkg/tokenModule"

	"github.com/google/wire"
)

var MemberRepo = wire.NewSet(impl.NewMemberRepo,wire.Bind(new(interfaces.MemberRepo), new(*implement.MemberRepo)))

func InitMemberController() *MemberController{
	wire.Build(
		NewMemberController ,
		tokenModule.NewTokenService ,  
		NewMemberService , 
		MemberRepo , 
		helpers.NewSqlSession ,
	)
	return &MemberController{}
}
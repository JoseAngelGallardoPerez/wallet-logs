package middlewares

import (
	pbUsers "github.com/Confialink/wallet-users/rpc/proto/users"
	"github.com/gin-gonic/gin"

	"github.com/Confialink/wallet-logs/internal/logs/acl"
	"github.com/Confialink/wallet-logs/internal/logs/errcodes"
)

func CanViewSystemLogs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("_user").(*pbUsers.User)
		if !acl.HasPermission(user, acl.PermissionViewSystemLogs) {
			errcodes.AddError(ctx, errcodes.CodeForbidden)
			ctx.Abort()
			return
		}
	}
}

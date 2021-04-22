package acl

import (
	"github.com/Confialink/wallet-logs/internal/srvdiscovery"
	"context"
	"net/http"

	"github.com/Confialink/wallet-logs/internal/logs/config/logs"
	pbPermissions "github.com/Confialink/wallet-permissions/rpc/permissions"
	pbUsers "github.com/Confialink/wallet-users/rpc/proto/users"
)

const PermissionViewSystemLogs = "view_system_log"

func HasPermission(user *pbUsers.User, permission string) bool {
	switch user.RoleName {
	case "root":
		return true
	case "admin":
		return checkAdminPermission(user, permission)
	default:
		return false
	}
}

func checkAdminPermission(user *pbUsers.User, permission string) bool {
	checker := getChecker()
	if checker == nil {
		return false
	}

	resp, err := checker.Check(context.Background(),
		&pbPermissions.PermissionReq{
			UserId:    user.UID,
			ActionKey: permission,
		},
	)
	if err != nil {
		logs.Logger.Error("Failed to get permission response", "error", err)
		return false
	}
	return resp.IsAllowed
}

func getChecker() pbPermissions.PermissionChecker {
	permissionsUrl, err := srvdiscovery.ResolveRPC(srvdiscovery.ServiceNamePermissions)
	if err != nil {
		logs.Logger.Error("Failed to get permissions rpc url", "error", err)
		return nil
	}
	return pbPermissions.NewPermissionCheckerProtobufClient(permissionsUrl.String(), http.DefaultClient)
}

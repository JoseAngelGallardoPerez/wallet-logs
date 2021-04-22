module github.com/Confialink/wallet-logs

go 1.13

replace github.com/Confialink/wallet-logs/rpc/logs => ./rpc/logs

require (
	github.com/Confialink/wallet-logs/rpc/logs v0.0.0-00010101000000-000000000000
	github.com/Confialink/wallet-permissions/rpc/permissions v0.0.0-20210218072732-21caf4a66e86
	github.com/Confialink/wallet-pkg-discovery/v2 v2.0.0-20210217105157-30e31661c1d1
	github.com/Confialink/wallet-pkg-env_config v0.0.0-20210217112253-9483d21626ce
	github.com/Confialink/wallet-pkg-env_mods v0.0.0-20210217112432-4bda6de1ee2c
	github.com/Confialink/wallet-pkg-errors v0.1.1
	github.com/Confialink/wallet-pkg-json_response v0.0.0-20210218075032-4cb33035b8f5
	github.com/Confialink/wallet-pkg-list_params v0.0.0-20210217104359-69dfc53fe9ee
	github.com/Confialink/wallet-pkg-model_serializer v0.0.0-20210217111055-c5e1cb1a75c7
	github.com/Confialink/wallet-pkg-service_names v0.0.0-20210217112604-179d69540dea
	github.com/Confialink/wallet-pkg-utils v0.0.0-20210217112822-e79f6d74cdc1
	github.com/Confialink/wallet-settings/rpc/proto/settings v0.0.0-20210218070334-b4153fc126a0
	github.com/Confialink/wallet-users/rpc/proto/users v0.0.0-20210218071418-0600c0533fb2
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/inconshreveable/log15 v0.0.0-20201112154412-8562bdadbbac
	github.com/jinzhu/gorm v1.9.15
)

print("Wallet Logs")

load("ext://restart_process", "docker_build_with_restart")

cfg = read_yaml(
    "tilt.yaml",
    default = read_yaml("tilt.yaml.sample"),
)

local_resource(
    "logs-build-binary",
    "make fast_build",
    deps = ["./cmd", "./internal"],
)
local_resource(
    "logs-generate-protpbuf",
    "make gen-protobuf",
    deps = ["./rpc/logs/logs.proto"],
)

docker_build(
    "velmie/wallet-logs-db-migration",
    ".",
    dockerfile = "Dockerfile.migrations",
    only = "migrations",
)
k8s_resource(
    "wallet-logs-db-migration",
    trigger_mode = TRIGGER_MODE_MANUAL,
    resource_deps = ["wallet-logs-db-init"],
)

wallet_logs_options = dict(
    entrypoint = "/app/service_logs",
    dockerfile = "Dockerfile.prebuild",
    port_forwards = [],
    helm_set = [],
)

if cfg["debug"]:
    wallet_logs_options["entrypoint"] = "$GOPATH/bin/dlv --continue --listen :%s --accept-multiclient --api-version=2 --headless=true exec /app/service_logs" % cfg["debug_port"]
    wallet_logs_options["dockerfile"] = "Dockerfile.debug"
    wallet_logs_options["port_forwards"] = cfg["debug_port"]
    wallet_logs_options["helm_set"] = ["containerLivenessProbe.enabled=false", "containerPorts[0].containerPort=%s" % cfg["debug_port"]]

docker_build_with_restart(
    "velmie/wallet-logs",
    ".",
    dockerfile = wallet_logs_options["dockerfile"],
    entrypoint = wallet_logs_options["entrypoint"],
    only = [
        "./build",
        "zoneinfo.zip",
    ],
    live_update = [
        sync("./build", "/app/"),
    ],
)
k8s_resource(
    "wallet-logs",
    resource_deps = ["wallet-logs-db-migration"],
    port_forwards = wallet_logs_options["port_forwards"],
)

yaml = helm(
    "./helm/wallet-logs",
    # The release name, equivalent to helm --name
    name = "wallet-logs",
    # The values file to substitute into the chart.
    values = ["./helm/values-dev.yaml"],
    set = wallet_logs_options["helm_set"],
)

k8s_yaml(yaml)

package main

//导入包含负载测试的包
import (
	"github.com/informalsystems/tm-load-test/pkg/loadtest"
)

// 使用说明常量
const appLongDesc = `Load testing application for Tendermint with optional coordinator/worker mode.
Generates large quantities of arbitrary transactions and submits those
transactions to one or more Tendermint endpoints. By default, it assumes that
you are running the kvstore ABCI application on your Tendermint network.

To run the application in a similar fashion to tm-bench (STANDALONE mode):
    tm-load-test -c 1 -T 10 -r 1000 -s 250 \
        --broadcast-tx-method async \
        --endpoints ws://tm-endpoint1.somewhere.com:26657/websocket,ws://tm-endpoint2.somewhere.com:26657/websocket

To run the application in COORDINATOR mode:
    tm-load-test \
        coordinator \
        --expect-workers 2 \
        --bind localhost:26670 \
        --shutdown-wait 60 \
        -c 1 -T 10 -r 1000 -s 250 \
        --broadcast-tx-method async \
        --endpoints ws://tm-endpoint1.somewhere.com:26657/websocket,ws://tm-endpoint2.somewhere.com:26657/websocket

To run the application in WORKER mode:
    tm-load-test worker --coordinator localhost:26680

NOTES:
* COORDINATOR mode exposes a "/metrics" endpoint in Prometheus plain text
* format
  which shows total number of transactions and the status for the coordinator
  and all connected workers.
* The "--shutdown-wait" flag in COORDINATOR mode is specifically to allow your
  monitoring system some time to obtain the final Prometheus metrics from the
  metrics endpoint.
* In WORKER mode, all load testing-related flags are ignored. The worker always
  takes instructions from the coordinator node it's connected to.
`

func main() {
	loadtest.Run(&loadtest.CLIConfig{ //启动负载测试
		AppName:              "tm-load-test",
		AppShortDesc:         "Load testing application for Tendermint kvstore", //应用的简单描述
		AppLongDesc:          appLongDesc,                                       //应用的详细描述
		DefaultClientFactory: "kvstore",                                         //默认的客户端类型是kvstore
	})
}

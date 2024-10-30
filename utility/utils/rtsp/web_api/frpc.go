package web_api

import (
	"fmt"
	frpc "github.com/kshdb/ks_frpc"
)

/*
运行rpc反向代理
*/
func RunRpc() {
	//frpc.Run("./frpc.ini")
	_content := fmt.Sprintf(`
[common]
server_addr = %s
server_port = %s
token=%s

[api]
type = %s
local_ip = %s
local_port = %s
remote_port = %s
`, "test2.api.cnwtn.com", "7005", "AvbuYer!35.com", "tcp", "127.0.0.1", "8011", "18083")
	frpc.RunContent(_content)
}

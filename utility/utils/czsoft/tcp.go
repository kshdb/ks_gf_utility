package czsoft

import (
	"fmt"
	"log"
	"net"
)

/*
tcp通讯对象
*/
type Tcp struct {
	//ip地址
	Ip string `json:"ip"`
	//端口号
	Port int `json:"port"`
}

/*
tcp打开通讯连接
*/
func (t *Tcp) GetConn() (conn net.Conn, err error) {
	// 服务端建立连接
	conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", t.Ip, t.Port))
	return
}

/*
tcp读数
*/
func (t *Tcp) Read(_conn net.Conn) (data []byte) {
	// 读取数据
	buffer := make([]byte, 128) // 适当的缓冲区大小
	n, err := _conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	// 处理读取到的数据
	data = buffer[:n]
	return
}

/*
tcp写数
*/
func (t *Tcp) Write(_conn net.Conn, buffer []byte) (n int, err error) {
	// 读取数据
	n, err = _conn.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return
}

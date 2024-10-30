package czsoft

import (
	"flag"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

/*
websocket对象
*/
type Ws struct {
	//ip地址
	Ip string `json:"ip"`
	//端口号
	Port int `json:"port"`
}

/*
Ws打开通讯连接
*/
func (t *Ws) GetConn(_ctx g.Ctx) (conn *websocket.Conn, err error) {
	addr := flag.String("addr", fmt.Sprintf("%s:%d", t.Ip, t.Port), "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	return
}

/*
Ws读数
*/
func (t *Ws) Read(_conn *websocket.Conn) (data []byte) {
	// 读取数据
	_, data, err := _conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}
	return
}

/*
Ws写数
*/
func (t *Ws) Write(_conn *websocket.Conn, buffer []byte) (err error) {
	// 读取数据
	err = _conn.WriteMessage(websocket.TextMessage, buffer)
	if err != nil {
		log.Fatal(err)
	}
	return
}

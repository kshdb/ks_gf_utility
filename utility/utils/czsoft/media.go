package czsoft

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

/*
直播流对象
*/
type Media struct {
	//视频流id
	Uuid string `json:"uuid"`
	//视频流单元
	ChannelID string `json:"channel_id"`
	//摄像头IP地址
	Ip string `json:"ip"`
	//摄像头端口号
	Port int `json:"port"`
	//摄像头账号
	Uid string `json:"uid"`
	//摄像头密码
	Pwd string `json:"pwd"`
}

/*
获取直播流地址
*/
func (m *Media) GetWsFlvUrl(ctx g.Ctx) (_url string) {
	_address, _ := g.Cfg().Get(ctx, "server.address")
	_url = fmt.Sprintf("ws://127.0.0.1:%s/stream/%s/channel/%s/ws.flv", _address, m.Uuid, m.ChannelID)
	return
}

/*
截取图片
*/
func (m *Media) GetImg(ctx g.Ctx) {
	_address, _ := g.Cfg().Get(ctx, "server.address")
	go g.Client().Get(ctx, fmt.Sprintf("http://127.0.0.1:%s/stream/%s/channel/%s/create_img", _address, m.Uuid, m.ChannelID))
}

/*
截取小视频
*/
func (m *Media) GetVideo(ctx g.Ctx) {

}

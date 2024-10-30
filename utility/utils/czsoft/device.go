package czsoft

import (
	"encoding/binary"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tidwall/gjson"
	"net"
)

var (
	//用于获取当前车牌号
	Onplate = make(chan string)
)

/*
称重软件硬件设备对象
*/
type CzSoftDevice struct {
	//仪表(读取地磅数据)
	YiBiao []YiBiaoModel `json:"yi_biao"`
	//道闸(控制抬杆)
	DaoZha []DaoZhaModel `json:"dao_zha"`
	//红外光栅(防作弊提醒)
	GuangShan []GuangShanModel `json:"guang_shan"`
	//手动按钮(不保存手动干预)
	AnNiu []AnNiuModel `json:"an_niu"`
	//车牌识别(识别车牌自动抬杆)
	ChePai []ChePaiModel `json:"che_pai"`
	//全景摄像头(保存时截图或截取视频)
	QuanJing []QuanJingModel `json:"quan_jing"`
	//红绿灯(规范车辆调度)
	HongLvDeng []any `json:"hong_lv_deng"`
	//Led屏幕(可视化提醒)
	Led []any `json:"led"`

	//读卡器
	DukaQi []DukaQiModel `json:"duka_qi"`
	//雷达(控制落杆)
	LeiDa []any `json:"lei_da"`
	//地感线圈(控制落杆功能与雷达可选其一)
	DiGan []any `json:"di_gan"`
}

/*
硬件基础对象
*/
type BaseModel struct {
	//品牌
	Brand string `json:"brand"`
	//型号
	TypeName string `json:"type_name"`
	//自定义名称
	NickName string `json:"nick_name"`
}

/*
通讯基础对象
*/
type BaseTxModel struct {
	BaseModel
	//串口信息
	SerialInfo Serial `json:"serial_info"`
	//Tcp信息
	TcpInfo Tcp `json:"tcp_info"`
	//websocket信息
	WsInfo Ws `json:"ws_info"`
	//通讯方式 [com,tcp,ws]
	Method string `json:"method"`
}

/*
仪表对象
*/
type YiBiaoModel struct {
	BaseTxModel
	//稳定时长
	WdLen int `json:"wd_len"`
	//稳定次数
	WdNum int `json:"wd_num"`
}

/*
道闸对象
*/
type DaoZhaModel struct {
	BaseTxModel
	//控制时长
	WdLen int `json:"wd_len"`
}

/*
车牌识别对象
*/
type ChePaiModel struct {
	BaseModel
	//流媒体信息
	MediaInfo Media `json:"media_info"`
	//信令端口
	TcpPort int `json:"port"`
}

/*
全景摄像头对象
*/
type QuanJingModel struct {
	BaseModel
	//流媒体信息
	MediaInfo Media `json:"media_info"`
}

/*
光栅对象
*/
type GuangShanModel struct {
	BaseTxModel
	//控制时长
	WdLen int `json:"wd_len"`
}

/*
按钮对象
*/
type AnNiuModel struct {
	BaseTxModel
	//控制时长
	WdLen int `json:"wd_len"`
}

/*
读卡器对象
*/
type DukaQiModel struct {
	BaseTxModel
	//控制时长
	WdLen int `json:"wd_len"`
}

// --------------------------车牌识别方法------------------------------//
//
//	func (c *ChePaiModel) GetPlate1() {
//		for {
//			c.OnPlate <- "测试" + gtime.Now().Format("Y-m-d H:i:s")
//			time.Sleep(time.Millisecond * 200)
//		}
//	}
/*
获取车牌号
@摄像机序列号
*/
func (c *ChePaiModel) GetPlate() {
	var _ctx g.Ctx
	_id := ""
	g.Try(_ctx, func(ctx g.Ctx) {
	_start:
		_conn := c.getConn()
		defer _conn.Close()
		//-------------------------基础指令------------------------------------//
		_cmd := fmt.Sprintf(`{"cmd":"getsn"}`)
		//_cmd := fmt.Sprintf(`{"cmd":"getsn","id":"%s"}`, _id)
		c.SendCmd(_conn, _cmd)
		data := make([]byte, 1024)
		_num, _err := _conn.Read(data)
		if _num == 0 || _err != nil {
			_conn.Close()
			goto _start
		} else {
			_id = gconv.String(gjson.Get(string(data), "value"))
		}
		if _id != "" {
			//发送指令--获取设备的序列号
			//SendCmd(_conn, `{"cmd":"getsn","id":"630b5d5d-1ae86241"}`)
			//发送指令--获取设备的硬件版本信息
			//SendCmd(_conn, `{"cmd":"get_hw_board_version","id":"630b5d5d-1ae86241"}`)
			//发送指令--获取设备当前时间戳
			//SendCmd(_conn, `{"cmd":"get_device_timestamp","id":"630b5d5d-1ae86241"}`)
			//发送指令--设置系统时间
			//SendCmd(_conn, `{"cmd":"set_time","id":"630b5d5d-1ae86241","timestring" : "2015-03-17 20:47:02"}`)
			//发送指令--设置网络参数
			// SendCmd(_conn, `{"cmd":"set_networkparam","id":"630b5d5d-1ae86241","body":{ "ip":"192.168.1.177", "netmask":"255.255.255.0", "gateway":"192.168.1.1", "dns":"0.0.0.0", "source":0
			// }}`)
			//发送指令--获取网络参数
			//SendCmd(_conn, `{"cmd":"get_networkparam","id":"630b5d5d-1ae86241","source":0}`)
			//发送指令--设置中心服务器网络参数
			//SendCmd(_conn, `{"cmd":"set_centerserver_net","id":"630b5d5d-1ae86241","body":{ "hostname":"192.168.1.106", "port":80, "enable_ssl":false, "ssl_port":443, "http_timeout":5}}`)
			//发送指令--设置当前配置为用户默认配置
			// SendCmd(_conn, `{"cmd":"set_user_default_cfg","id":"630b5d5d-1ae86241"}`)
			//发送指令--修改设备 admin 密码
			//SendCmd(_conn, `{"cmd":"set_adminpass","id":"630b5d5d-1ae86241","body":{ "old_pass":"asgwe4AGSAD45", "new_pass":"fdas213asfdgad"}}`)
			//发送指令--重启设备
			//SendCmd(_conn, `{"cmd":"reboot_dev","id":"630b5d5d-1ae86241"}`)
			//发送指令--重启设备
			//SendCmd(_conn, `{"cmd":"reboot_dev","id":"630b5d5d-1ae86241"}`)
			//发送指令-- 获取设备版本信息
			//SendCmd(_conn, `{"cmd":"get_product_info","id":"630b5d5d-1ae86241"}`)
			//发送指令-- 获取设备 4G 参数信息
			//SendCmd(_conn, `{"cmd":"get_4g_param","id":"630b5d5d-1ae86241"}`)
			//发送指令-- 设置设备 4G 参数信息
			//SendCmd(_conn, `{"cmd":"set_4g_param","id":"630b5d5d-1ae86241","body" : { "sub_cmd" : "set_apn", "apn_param" : { "apnaddr" : "192.168.1.12", "username" : "apn_user", "passwd" : "apn_pass", "authentication" : 0}}}`)

			//发送指令-- 设置设备名称
			//SendCmd(_conn, `{"cmd":"set_dev_name","id":"630b5d5d-1ae86241","body":{ "title":"Ivs"}}`)

			//-------------------------车牌识别------------------------------------//
			//发送指令-- 配置推送数据方式
			//SendCmd(_conn, `{"cmd":"ivsresult","id":"630b5d5d-1ae86241","enable": true,"format":"json","image":false,"image_type":0}`)
			//发送指令-- 获取最近一次识别结果
			//SendCmd(_conn, `{"cmd":"getivsresult","id":"630b5d5d-1ae86241","image" : true, "format" : "json"}`)
			//发送指令-- 手动触发车牌识别
			//SendCmd(_conn, `{"cmd":"trigger","id":"630b5d5d-1ae86241"}`)
			//发送指令-- 获取记录图片
			//SendCmd(_conn, `{"cmd":"get_image","id":"630b5d5d-1ae86241"}`)
			//发送指令-- 获取记录图片
			//SendCmd(_conn, `{"cmd":"get_offline_image","id":"630b5d5d-1ae86241"}`)
			//发送指令-- 抓取当前图片
			//SendCmd(_conn, `{"cmd":"get_snapshot","id":"630b5d5d-1ae86241"}`)
			//发送指令-- 获取视频播放的 uri:
			//SendCmd(_conn, `{"cmd":"get_rtsp_uri","id":"630b5d5d-1ae86241"}`)
			//发送指令-- 获取虚拟线圈参数
			//SendCmd(_conn, `{"cmd":"get_virloop_para","id":"630b5d5d-1ae86241"}`)
			//发送指令-- 设置客户定制 sn 序列号
			//SendCmd(_conn, `{"cmd":"set_oem_sn_info","id":"630b5d5d-1ae86241","body":{ "oem_sn":"1234-5689"}}`)
			//发送指令-- 获取客户定制 SN 序列号
			//SendCmd(_conn, `{"cmd":"get_oem_sn_info","id":"630b5d5d-1ae86241"}`)
			//发送指令-- 得到在线设备信息，含自己
			//SendCmd(_conn, `{"cmd":"dg_json_request","id":"630b5d5d-1ae86241","body":{"type":"online_devices"}}`)
			//接收指令
			//snLen := recvPacketSize(_conn)
			//if snLen > 0 {
			//	data := make([]byte, snLen)
			//	recvLen, _ := recvBlock(_conn, data, snLen)
			//	sn := string(data[:recvLen])
			//	fmt.Println("接收到回复的指令是--", sn)
			//}

			//snLen = recvPacketSize(_conn)
			//if snLen > 0 {
			//	data := make([]byte, snLen)
			//	recvLen, _ := recvBlock(_conn, data, snLen)
			//	sn := string(data[:recvLen])
			//	fmt.Println("登录结果是：", sn)
			//}
			//发送指令-- 配置推送数据方式
			_cmd = fmt.Sprintf(`{"cmd":"ivsresult","id":"%s","enable":true,"format":"json","image":false,"image_type":1}`, _id)
			c.SendCmd(_conn, _cmd)
			for {
				data := make([]byte, 1024)
				_num, _err := _conn.Read(data)
				if _num == 0 || _err != nil {
					_conn.Close()
					goto _start
				}
				//fmt.Println("监听到服务器发来的信息是？====", string(data), _num, _err)
				_plate := gjson.Get(string(data), "PlateResult.license")
				if gconv.String(_plate) != "" {
					//定义转码
					enc := mahonia.NewDecoder("gbk")
					Onplate <- enc.ConvertString(gconv.String(_plate))
					//fmt.Println("识别到的车牌是", gtime.Now().Format("Y-m-d H:i:s"), enc.ConvertString(gconv.String(_plate)))
				}
			}

		}

	})
}

/*
建立tcp链接
*/
func (c *ChePaiModel) getConn() (conn net.Conn) {
	_address := fmt.Sprintf("%s:%d", c.MediaInfo.Ip, c.TcpPort)
	conn, err := net.Dial("tcp", _address)
	if err != nil {
		fmt.Println("tcp创建失败:", err)
		return
	}
	return conn
}

/*
发送指令
*/
func (c *ChePaiModel) SendCmd(conn net.Conn, cmd string) {
	len := len(cmd)
	header := make([]byte, 8)
	header[0] = 'V'
	header[1] = 'Z'
	binary.BigEndian.PutUint32(header[4:8], uint32(len))
	conn.Write(header)
	conn.Write([]byte(cmd))
}

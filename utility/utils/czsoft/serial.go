package czsoft

import (
	"go.bug.st/serial"
	"log"
)

/*
串口对象
*/
type Serial struct {
	//串口号
	PortName string
	//波特率
	BaudRate int
	//数据位
	DataBits int
	//停止位 0 1 2
	StopBits int
	//校验位 0 1 2 3 4
	Parity int
}

/*
获取串口列表
*/
func (s *Serial) ComList() (_list []string) {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		return
	}
	for _, port := range ports {
		_list = append(_list, port)
	}
	return
}

/*
串口打开通讯连接
*/
func (s *Serial) GetConn() (conn serial.Port, err error) {
	// 配置串行端口
	_mode := &serial.Mode{
		BaudRate: s.BaudRate,                  //波特率
		DataBits: s.DataBits,                  //数据位
		StopBits: serial.StopBits(s.StopBits), //停止位
		Parity:   serial.Parity(s.Parity),     //校验位
	}
	conn, err = serial.Open(s.PortName, _mode)
	return
}

/*
串口读数
*/
func (s *Serial) Read(_conn serial.Port) (data []byte) {
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
串口写数
*/
func (s *Serial) Write(_conn serial.Port, buffer []byte) (n int, err error) {
	// 读取数据
	n, err = _conn.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return
}

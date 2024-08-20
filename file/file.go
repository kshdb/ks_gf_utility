package file

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
)

/*
文件对象
*/
type FileInfo struct {
	Ctx        g.Ctx             //上下文
	File       *ghttp.UploadFile `json:"avatarfile" type:"file" dc:"选择上传文件"`
	Name       string            `json:"name"  dc:"自定义文件名称"`        // 自定义文件名称
	PathR      string            `json:"pathR"  dc:"自定义二级文件目录"`     // 自定义二级文件目录
	RandomName bool              `json:"randomName"  dc:"是否随机命名文件"` // 是否随机命名文件
	FileType   string            `json:"fileType"  dc:"文件类型"`       // 文件类型
}

/*
文件上传
*/
func (f *FileInfo) Upload() (fileName string, err error) {
	uploadPath := g.Cfg().MustGet(f.Ctx, "ks_gf_file.file.upload.path").String()
	if uploadPath == "" {
		err = gerror.New("上传文件路径配置不存在")
	} else {
		dateDirName := gtime.Now().Format("Ymd")
		fileName, err = f.File.Save(gfile.Join(uploadPath, f.PathR, dateDirName), f.RandomName)
		fileName = fmt.Sprintf("/%s/%s/%s/%s", uploadPath, f.PathR, dateDirName, fileName)
	}
	return
}

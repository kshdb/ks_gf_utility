package file

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"io"
	"mime/multipart"
	"os"
)

/*
文件对象
*/
type FileInfo struct {
	Ctx        g.Ctx //上下文
	File       *ghttp.UploadFile
	Files      []*multipart.FileHeader
	Name       string `json:"name"  dc:"自定义文件名称"`        // 自定义文件名称
	PathR      string `json:"pathR"  dc:"自定义二级文件目录"`     // 自定义二级文件目录
	RandomName bool   `json:"randomName"  dc:"是否随机命名文件"` // 是否随机命名文件
	FileType   string `json:"fileType"  dc:"文件类型"`       // 文件类型
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

/*
多文件上传
*/
func (f *FileInfo) Uploads() (fileNames []string, err error) {
	uploadPath := g.Cfg().MustGet(f.Ctx, "ks_gf_file.file.upload.path").String()
	if len(f.Files) > 0 {
		for _, _f := range f.Files {
			//打开上传文件
			file, _ := _f.Open()
			defer file.Close()
			//创建上传目录
			if !gfile.Exists(uploadPath) {
				os.Mkdir(uploadPath, os.ModePerm)
			}
			dateDirName := gtime.Now().Format("Ymd")
			//创建上传文件
			fileName := fmt.Sprintf("%s/%s/%s/%s", uploadPath, f.PathR, dateDirName, _f.Filename)
			cur, _ := os.Create(fileName)
			defer cur.Close()
			io.Copy(cur, file)
			fileNames = append(fileNames, _f.Filename)
		}
	}
	return
}

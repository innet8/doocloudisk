// Code generated by hertz generator. DO NOT EDIT.

package aliyun

import (
	aliyun "github.com/cloudisk/biz/handler/aliyun"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_file := _api.Group("/file", _fileMw()...)
			{
				_content := _file.Group("/content", _contentMw()...)
				_content.POST("/office", append(_officeuploadMw(), aliyun.OfficeUpload)...)
				_content.POST("/save", append(_saveMw(), aliyun.Save)...)
				_content.POST("/upload", append(_uploadMw(), aliyun.Upload)...)
			}
		}
	}
}
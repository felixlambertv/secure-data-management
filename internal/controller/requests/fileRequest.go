package requests

import "mime/multipart"

type UploadFileRequest struct {
	Files      []*multipart.FileHeader `form:"files[]"`
	Permission []string                `form:"permission[]"`
}

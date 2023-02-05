package upload

type Group struct{}

func (g *Group) Upload() *UploadApi {
	return &insUpload
}

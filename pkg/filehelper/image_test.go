package filehelper

import (
	"server/pkg/mydev"
	"testing"
)

func TestFind(t *testing.T) {
	mydev.StartTrack()
	imgPath := FindFirstImage(`D:\data\camera\GrpcCamera\GrpcCamera\GrpcCamera`)
	println(imgPath)
	mydev.EndTrack()
}

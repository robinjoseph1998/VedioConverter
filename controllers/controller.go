package controllers

import (
	"net/http"

	"os/exec"

	"github.com/gin-gonic/gin"
	"gopkg.in/vansante/go-ffprobe.v2"
)

func Mp4ToMkvConverter(c *gin.Context) {
	vedioFile, err := c.FormFile("vedio")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	UploadFilePath := "/tmp" + vedioFile.Filename
	if err := c.SaveUploadedFile(vedioFile, UploadFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't save the vedio", "error": err.Error()})
		return
	}
	probeResult, err := ffprobe.ProbeURL(c, UploadFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	InputFormat := probeResult.Format.FormatName
	if InputFormat != "mp4" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the input file is not Mp4", "error": err.Error()})
		return
	}
	//Transcode Mp4 To Mkv
	outFile := "converted.mkv"
	cmd := exec.Command("ffmpeg", "-i", UploadFilePath, outFile)
	err = cmd.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't convert the vedio", "error": err.Error()})
		return
	}
	c.File(outFile)
	c.JSON(http.StatusOK, gin.H{"message": "vedio converted successfully"})
}

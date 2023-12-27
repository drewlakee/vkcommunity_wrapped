package utils

import (
	"log"
	"os"
	"vkcommunity_wrapped/internal/models"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type NewFile struct {
	file *os.File
}

func CreateNewFile(context models.Context, filename string) *NewFile {
	_ = os.MkdirAll(context.OutputDir, os.ModePerm)
	_ = os.Remove(context.OutputDir + "/" + filename)
	output, err := os.Create(context.OutputDir + "/" + filename)
	CheckError(err)
	return &NewFile{file: output}
}

func (file *NewFile) Close() {
	file.file.Close()
}

func (file *NewFile) Write(data string) {
	_, err := file.file.WriteString(data)
	CheckError(err)
}

package upload

import (
	"fmt"
	"io"
	"library/global"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	TypeDoc FileType = iota + 1
	// TypePdf
	// TypeEpub
	// TypeMobi
)

func GetFileName(name string) string {
	// ext:=GetFileExt(name)
	// fileName:=
	return name
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

//如果存在 aa.txt，就返回aa(1).txt，一直到...aa(n).txt
func GetUniqueFilePath(filename string) string {
	filePath := path.Join(GetSavePath(), filename)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return filename
	} else {
		ext := GetFileExt(filename)
		partName := filename[:len(filename)-len(ext)]
		for i := 1; ; i++ {
			newFileName := partName + fmt.Sprintf("(%v)", i) + ext
			filePath = path.Join(GetSavePath(), newFileName)
			_, err = os.Stat(filePath)
			if os.IsNotExist(err) {
				return newFileName
			}
		}
	}
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	switch t {
	case TypeDoc:
		for _, allowExt := range global.AppSetting.UploadBookAllowExts {
			if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	return false
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return nil
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

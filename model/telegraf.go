package model

import (
	"os"
	"time"
)

type File struct {
	Name string
}

// GetFile
// @author: [lhuiy119](https://github.com/lhuiy119)
// @description: 获取文件内容
func GetFile(path string) (content string, err error) {
	contentBytes, err := os.ReadFile(telegrafRoot + path)
	return string(contentBytes), err
}

// GetDir
// @author: [lhuiy119](https://github.com/lhuiy119)
// @description: 获取文件夹内的文件的路径
func GetDir(path string) (filesInfo []FileInfo, err error) {
	dirEntry, err := os.ReadDir(telegrafRoot + path)
	for _, v := range dirEntry {
		if !v.IsDir() {
			filesInfo = append(filesInfo, FileInfo{
				Path: path + string(os.PathSeparator) + v.Name(),
			})
		}
	}
	return filesInfo, err
}

// Update
// @author: [lhuiy119](https://github.com/lhuiy119)
// @description: 更新文件
func Update(path string, content string) error {
	err := os.WriteFile(telegrafRoot+path, []byte(content), os.ModePerm)
	return err
}

// Touch
// @author: [lhuiy119](https://github.com/lhuiy119)
// @description: 更新文件的访问时间和修改时间
func Touch(path string) error {
	err := os.Chtimes(telegrafRoot+path, time.Now(), time.Now())
	return err
}

// Delete
// @author: [lhuiy119](https://github.com/lhuiy119)
// @description: 删除文件
func Delete(path string) error {
	err := os.Remove(telegrafRoot + path)
	return err
}

// Exist
// @author: [lhuiy119](https://github.com/lhuiy119)
// @description: 删除文件
func Exist(path string) error {
	file, err := os.Open(telegrafRoot + path)
	defer file.Close()
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	//"path/filepath"
)

func main() {
	dir, _ := os.Getwd() // 获取当前路径，pwd
	fmt.Println(dir)
	curFile := path.Join(dir, "go练习代码", "src", "01基础", "15io操作", "02创建文件", "main.go") // 拼接项目路径
	fmt.Println(curFile)
	dir2 := path.Dir(curFile) // 获取指定文件路径，dirname（相对路径优先）
	fmt.Println(dir2)
	name := "fuck.txt"
	filePath := path.Join(dir2, name)
	newFile, err := os.Create(filePath) // 创建文件
	if err != nil {
		log.Fatal(err)
	}
	defer func(newFile *os.File) {
		err = newFile.Close() // 关闭文件
		if err != nil {
			log.Fatal(err)
		}
	}(newFile)
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0777)
	if err != nil {
		if os.IsPermission(err) {
			log.Fatal("Error: Write permission denied")
		}
	}
	file.Close()
	file, err = os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		if os.IsPermission(err) {
			log.Fatal("Error: Read permission denied")
		}
	}
	file.Close()

	n, err := newFile.WriteString("fuck") // 返回写入文件的长度
	fmt.Println("Write Size: ", n)
	if err != nil {
		log.Fatal(err)
	}
	stat, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Error: File does not exist")
		}
		log.Fatal(err)
	}
	fmt.Printf("File name: %s, Size in bytes: %d, Permissions: %v, Last Modified: %s, Is Dir: %v", stat.Name(), stat.Size(), stat.Mode(), stat.ModTime(), stat.IsDir())
	baseDir := path.Join(dir2, "fuck")
	os.MkdirAll(baseDir, 0777) // go 有点特殊，创建文件夹后还必须重新给遍权限
	os.Chmod(baseDir, 0777)
	filePath2 := path.Join(baseDir, "fuckme.txt")

	err = os.Rename(filePath, filePath2) // 移动 or 重命名文件，相当于 move？
	if err != nil {
		log.Fatal(err)
	}
	//err = os.Rename(filePath2, filePath)
	err = os.Truncate(filePath2, 100) // 裁剪文件到100字节；如果本来就少于100字节，以null填充；如果超过则丢弃；传0则会清空文件
	os.Remove(filePath2)
	stat, err = os.Stat(filePath2)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Error: File does not exist")
		}
		log.Fatal(err)
	} else {
		fmt.Println("delete filepath2 ok")
	}
	os.Remove(baseDir)
	stat, err = os.Stat(baseDir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Error: File does not exist")
		}
		log.Fatal(err)
	} else {
		fmt.Println("delete base dir ok")
	}

}

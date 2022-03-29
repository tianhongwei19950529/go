package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 输入输出的底层原理：
// 终端其实是一个文件，相关实例如下：
// os.Stdin: 标准输入，类型为 *File，文件描述符 0
// os.Stdout: 标准输出，类型为 *File，文件描述符 1
// os.Stderr: 标准错误输出，类型为 *File，文件描述符 2

// 以文件形式操作终端, 返回的第一个参数为 字节数...
func main1() {
	var buf [16]byte
	_, err := os.Stdin.Read(buf[:]) // 阻塞，等待输入
	if err != nil {
		return
	}
	_, err = os.Stdin.WriteString(string(buf[:]) + " me----")
	if err != nil {
		return
	}
	_, err = os.Stdout.WriteString(string(buf[:]) + " you----")
	if err != nil {
		return
	}
	_, err = os.Stderr.WriteString(string(buf[:]) + " all----")
	if err != nil {
		return
	}
}

// 文件操作api
// os.Create	// 根据文件名创建文件
// os.NewFile	// 根据文件描述符创建爱你文件
// os.Open		// 根据文件名打开文件
// os.OpenFile	// 根据文件名打开文件，可设置读写模式，权限等...
// os.Write		// 写 byte 数组
// os.WriteAt
// os.WriteString   // 直接写字符串
// os.Read
// os.ReadAt
// os.Remove

// 打开和关闭文件
func main2() {
	file, err := os.Open("./.gitignore") // 默认路径为go的项目路径，go run 运行目录...
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
}

// 写文件
func main3() {
	file, err := os.Create(".fuck.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)
	for i := 0; i < 5; i++ {
		// _, _ = file.WriteString("ab\n")  // 暴力写法
		_, err := file.WriteString("ab\n")
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = file.Write([]byte("cd\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// NOTE: []byte() 和 string() 一对，进行了类型转换，中间是数组...
	// fmt.Println([]byte{'a', 'b'})
	// fmt.Println([]byte("ab"))

	// fmt.Println([]string{"aaaa", "bbb"})
	// fmt.Println([]string("aaaa", "bbb"))

	// fmt.Println(string([]byte{'f', 'u', 'c', 'k'}))
	// 区别：string 进制修改，byte[] 可以修改；注意对应到内存的行为
}

// 读文件
func main4() {
	var file *os.File
	fmt.Printf("%p\n", file)
	file, err := os.Open("./.fuck.txt")
	fmt.Printf("%p\n", file)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer func(f *os.File) {
		fmt.Printf("%p\n", f)
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}(file)

	var buf [128]byte  // 数组    // 指定收取128字节
	var content []byte // 切片
	// var cc []string
	for {
		// n, err := file.Read(buf[:])	buf[:] 数组转切片
		n, err := file.Read(buf[:]) // Read 函数只能接收切片
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...) // append 第二个参数只能接收切片
		// content = append(content, buf)  // append 第二个参数只能接收切片
	}
	fmt.Println(string(content))
}

func readFile() error {
	file, err := os.Open("./.fuck.txt")
	buf, content := [128]byte{}, []byte{}
	if err != nil {
		return err
	}
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(content)
	// gopher 直接 return f.Close()，不放defer里面呢
	return file.Close()
}

// fuck buf 切片
func main5() {
	var w [128]byte
	var q []byte
	// w = []byte("fuck")    // 为啥子这行报错！！
	// w = [128]byte("fuck") // 为啥子这行报错！！
	q = []byte("fuck")
	fmt.Printf("%T----%v\n", w, w)
	fmt.Printf("%T----%v\n", w[:], w[:])
	fmt.Printf("%T----%v\n", q, q)
	fmt.Printf("%T----%v\n", q[:], q[:])
	var strss = []string{
		"aaa",
		"bbb",
		"ccc",
	}
	func(args ...string) {
		fmt.Println(args)
		ag := []string{
			"xxx",
			"yyy",
			"zzz",
		}
		ag = append(ag, args...)
		fmt.Println(ag)
		ag = append(ag, ag...)
		fmt.Println(ag)
	}(strss...)
}

// 拷贝文件
func main6() {
	srcFile, _ := os.Open("./.fuck.txt")
	dstFile, _ := os.Create("./.fucku.txt")
	// buf := make([]byte, 0)   // 大小绝对不能给0，不然会阻塞
	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			break
		}
		_, _ = dstFile.Write(buf[:n])
	}
	io.Copy(dstFile, srcFile)
	_ = srcFile.Close()
	_ = dstFile.Close()
}

// bufio 包实现了带缓冲区的读写，是对文件读写的封装
// bufio 写数据
func main7() {
	file, _ := os.OpenFile("./.fuck.txt", os.O_CREATE|os.O_WRONLY, 0666)
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, _ = writer.WriteString("hello\n")
	}
	// 刷新缓冲区，强制写出
	_ = writer.Flush()
	_ = file.Close()
}

// bufio 读数据
func main8() {
	file, _ := os.OpenFile("./.fuck.txt", os.O_RDONLY, 0666) // file 实现了 reader 接口？
	reader := bufio.NewReader(file)                          // strings.NewReader
	for {
		// content, err := reader.ReadString('\n')
		// fmt.Printf(content)
		var buf [2048]byte
		var content []byte
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		content = append(content, buf[:n]...)
		fmt.Println(string(content))
	}
}

// ioutil 工具包，也可以用来读写文件, api 更简单
func main9() {
	_ = ioutil.WriteFile("./.fuckall.txt", []byte("fuck all"), 0666)
}
func main10() {
	content, _ := ioutil.ReadFile("./.fuckall.txt")
	fmt.Println(string(content))
}

func main() {
	// main1()
	// main2()
	// main3()
	// main4()
	// main5()
	// _ = readFile()
	main6()
	// main7()
	// main8()
	// main9()
	// main10()
}

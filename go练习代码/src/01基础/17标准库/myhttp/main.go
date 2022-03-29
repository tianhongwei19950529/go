package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main1() {
	var resp *http.Response
	resp, err := http.Get("https://postman-echo.com/get?foo1=bar1&foo2=bar2")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("__________________")
	fmt.Println(resp.StatusCode)
	fmt.Println("__________________")
	fmt.Println(resp.ContentLength)
	fmt.Println("__________________")
	fmt.Println(resp.Body)
	fmt.Println("__________________")
	fmt.Println(resp.Proto)
	fmt.Println("__________________")
	fmt.Println(resp.Header)
	fmt.Println("__________________")
	fmt.Println(resp.Cookies())
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("__________________")
	fmt.Println(string(body))
}

func main2() {
	apiUrl := "https://www.7timer.info/bin/astro.php"
	data := url.Values{}
	data.Set("lon", "113.2")
	data.Set("lat", "23.1")
	data.Set("ac", "0")
	data.Set("unit", "metric")
	data.Set("output", "json")
	data.Set("tzshift", "0")
	fmt.Println(data.Encode())
	u, _ := url.ParseRequestURI(apiUrl)
	u.RawQuery = data.Encode()
	resp, _ := http.Get(u.String())
	fmt.Println("u.String()", u.String())
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func main3() {
	//apiUrl := "https://http.cat/200"   // 返回图片
	// https://dog.ceo/api/breeds/image/random
	//https://dog-api.kinduff.com/api/facts?number=10

	//https://postman-echo.com/post
}

// post json
func main4() {
	apiUrl := "https://postman-echo.com/post"
	contentType := "application/json"
	data := `{"name": "裤头", "age": 16}`
	resp, _ := http.Post(apiUrl, contentType, strings.NewReader(data)) // 从字符串创建 reader
	//resp, _ := http.Post(apiUrl, contentType, bufio.NewReader(data))	// 从文件创建 reader?
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

// post form
func main5() {
	apiUrl := "https://postman-echo.com/post"
	data := url.Values{}
	data.Set("name", "枯藤")
	data.Set("age", "12")
	resp, _ := http.PostForm(apiUrl, data)
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(b))
}

// post img, fuck!!!
func main6() {
	buf := new(bytes.Buffer)
	apiUrl := "https://postman-echo.com/post"
	//buf := new(bytes.Buffer)
	//buf := [128]byte{}
	//bufslice := make(buf, 10)
	f, _ := os.OpenFile("./cat.jpeg", os.O_WRONLY, 0666)
	writer := multipart.NewWriter(f)

	writer = multipart.NewWriter(buf) // This is import!!!
	//f, _ := os.Open("./cat.jpeg")
	//writer := bufio.NewWriter(f)
	formFile, _ := writer.CreateFormFile("file", "cat.jpeg") // writer 的这个方法...
	srcFile, _ := os.Open("./cat.jpeg")
	defer srcFile.Close()
	io.Copy(formFile, srcFile)
	contentType := writer.FormDataContentType()
	writer.Close()
	resp, _ := http.Post(apiUrl, contentType, buf)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	defer resp.Body.Close()
}

// 自定义client
func main7() {
	client := &http.Client{}
	jsonStr := []byte("{\"name\": \"aaa\", \"age\":16}")
	//f, _ := os.Open("./cat/jpeg")
	req, _ := http.NewRequest("POST", "https://postman-echo.com/post", bytes.NewReader(jsonStr)) // reader 或者 buffer？ bytes.NewReader(jsonStr) or bytes.NewBuffer(jsonStr)
	//req, _ := http.NewRequest("POST", "https://postman-echo.com/post", bufio.NewReader(f)) // reader 或者 buffer？ bytes.NewReader(jsonStr) or bytes.NewBuffer(jsonStr)
	//req, _ := http.NewRequest("GET", "https://postman-echo.com/get", nil) // GET 请求将最后的参数给 nil
	//req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

// 超时，cancel：https://colobu.com/2016/07/01/the-complete-guide-to-golang-net-http-timeouts/

// transport: pass

// 自定义 client
func main() {
	//main1()
	//go main2()
	//main4()
	//main5()
	//main6()
	main7()
}

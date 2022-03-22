package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func UpdateRulePoint(c *gin.Context) {
	firstname := c.DefaultQuery("firname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func TestParamGetInPath(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func TestParamPostInPath(c *gin.Context) {
	b := c.FullPath() == "/v1/users/user/:name/*action"
	c.String(http.StatusOK, "%t", b)
}

func TestQueryString(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func TestFormPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

//func TestParamFormMap(c *gin.Context) {
//	ids := c.QueryMap("ids")
//	names := c.PostFormMap("names")
//	fmt.Printf("ids: %v; names: %v", ids, names)
//}

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func TestJsonValidator(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.User != "manu" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in " + json.User + " " + json.Password})
}

func TestXmlValidator(c *gin.Context) {
	var xml Login
	if err := c.ShouldBindXML(&xml); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if xml.User != "manu" || xml.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in " + xml.User + " " + xml.Password})
}

func TestFormValidator(c *gin.Context) {
	var form Login
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if form.User != "manu" || form.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in " + form.User + " " + form.Password})
}

func TestQueryValidator(c *gin.Context) {
	var query Login
	//if err := c.ShouldBindQuery(&query); err != nil {    // ShouldBindQuery 和 ShouldBind 都可以
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if query.User != "manu" || query.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in " + query.User + " " + query.Password})
}

type Time time.Time

func (t Time) String() string {
	return time.Time(t).Format("2006-01-02")
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}

type Person struct {
	Name string `json:"name"`
	//Birthday   Time   `json:"birthday"`
	CreateTime Time `json:"create_time"`
	//CreateTime time.Time `form:"create_time" time_format:"unixNano"`
}

func TestTimeFormat(c *gin.Context) {
	var person Person
	//if err := c.ShouldBind(&person); err != nil {
	if err := c.ShouldBindJSON(&person); err != nil {
		//c.String(http.StatusOK, "fuck!!!")
		c.String(http.StatusOK, err.Error())
		return
	}
	log.Println(person.Name)
	//log.Println(person.Birthday)
	log.Println(person.CreateTime)
	c.String(http.StatusOK, "Success")
}

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "user comments")
	})
	//curl "http://localhost:5000/v1/users/user/luke/asdasd"
	users.GET("/user/:name/*action", TestParamGetInPath)
	//curl -X POST "http://localhost:5000/v1/users/user/luke/asdasd"
	users.POST("/user/:name/*action", TestParamPostInPath)
	//curl "http://localhost:5000/v1/users/user/welcome?firname=luke&lastname=rivel"
	users.GET("/user/welcome", TestQueryString)
	//curl -X POST -d "message=hello&nick=luke" "http://localhost:5000/v1/users/user/form_post"
	users.POST("/user/form_post", TestFormPost)
	//users.POST("/user/map", TestParamFormMap)
	//curl -H "Content-Type:application/json" -X POST -d '{"user":"luke", "password":"aqwe1213"}' "http://localhost:5000/v1/users/user/login_json"
	users.POST("/user/login_json", TestJsonValidator)
	// <>解码有问题 curl -H "Content-Type:text/xml" -X POST -d '<?xml version="1.0" encoding="UTF-8"?><root><user>manu</user><password>123</password></root>' "http://localhost:5000/v1/users/user/login_json"
	users.POST("/user/login_xml", TestXmlValidator)
	//curl -X POST -d 'user=luke&password=aqwe1213' "http://localhost:5000/v1/users/user/login_form"
	users.POST("/user/login_form", TestFormValidator)
	// curl "http://localhost:5000/v1/users/user/login_query?user=manu&password=123"
	users.GET("/user/login_query", TestQueryValidator)
	// 时间戳和时间字符格式
	users.POST("/user/time", TestTimeFormat)
	users.GET("/rules", UpdateRulePoint)

}

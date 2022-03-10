1.  简单的注意事项

    1.  open函数并不会连接数据库,他只会验证参数是不是正确.

    2.  Ping才会真正完成连接,验证是不是真的可以连接.

    3.  返回的DB对象是并发安全的,可以被多个goroutine使用,因此Open函数只应该调用一次,并且很少关闭.

    4.  初始化数据库应该放在额外的一个函数里面,不应该放在主函数里面

        1.  一定不要忘记引入引擎包

        2.  定义一个全局的db变量,记得大写,此时就不要在采用短变量声明.

            ```go
            package main
            
            import (
            	"database/sql"
            	"fmt"
            	_ "github.com/go-sql-driver/mysql"
            )
            
            var db *sql.DB
            
            func InitMysql() (err error) {
            	db, err = sql.Open("mysql", "root:123456@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local")
            	if err != nil {
            		return
            	}
            
            	err = db.Ping()
            	if err != nil {
            		return
            	}
            	db.SetMaxOpenConns(200)
            	db.SetMaxIdleConns(5)
            	return
            }
            
            func main() {
            	err := InitMysql()
            	if err != nil {
            		fmt.Println(err.Error())
            	}
            	defer db.Close()
            
            }
            
            ```

    5.  实现数据库连接池的时候不需要其他写法,只需要打开两个参数SetMaxOpenConns  SetMaxIdleConns

        1.  SetMaxIdleConns最大空闲连接数
        2.  SetMaxOpenConns最大连接数

2.  curd

    1.  单行查询用queryRow  多行查询用query  增删改 用exec
    2.  查询完毕之后必须用scan方法扫描到对应结构体里面,如果不调用scan 会导致一直占用当前连接.最终达到程序设置的上限
    3.  多行查询的时候最好手动加一个close 方法.
    4.  增加可以看到通过方法查看增加的ID,删除和修改可以看到影响行数
    
3.  mysql预处理

    1.  预处理用于处理批量的情况   先用Prepare将sql语句进行编译.之后再把参数放进去去执行.

4.  sql注入

    1.  任意时候都不能自己拼接sql语句
    2.  通过参数传入即可

5.  sql事务


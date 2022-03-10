1.  连接数据库

    1.  使用gorm连接数据的时候,需要导入对应的数据库驱动的包.

        ```go
        import (
        	"fmt"
        	"github.com/jinzhu/gorm"
        	_ "github.com/jinzhu/gorm/dialects/mysql" //必须导入驱动,但是我们用不到其他的东西所以用匿名变量赋值
        )
        
        func main()  {
        	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local")
        	if err != nil{
        		fmt.Printf(err.Error())
        	}
        	defer db.Close()
        }
        ```

2.  创建表

    1.  使用自动同步表进行同步,如果定义结构体有数据表没有的字段,会将这个字段加到数据表的后面(即更改数据表的结构).

    2.  关于表名

        1.  gorm 默认将驼峰体转成下划线+小写并且加上复数形式   (UserInfo > user_infos).

        2.  去除复数形式

            ```go
            //默认参数为false  即采用复数形式,可以更改为true  取消复数形式.
            db.SingularTable(true)
            ```

        3.  指定表名的几种方式

            1.  给结构体增加方法 方法名必须是TableName其他的并不行.

                ```go
                //此时表名就是return的user 而不是user_infos了
                func (UserInfo) GetTableName()  string  {
                	return "user"
                }
                //同步到数据库中
                db.AutoMigrate(&UserInfo{})
                ```

            2.  直接指定表名

                ```go
                db.Table("users").CreateTable(&UserInfo{})
                ```

            3.  可以通过gorm增加对应的表规则

    3.  关于列名

        1.  gorm 默认将驼峰体转成下划线+小写.

        2.  可以利用tags指定对应的字段名.

        3.  可以在tag中增加响应的限制.

            ```go
            type Animal struct {
            	AnimalId    int64     `gorm:"column:beast_id;size:150"`         // set column name to `beast_id`
            	Birthday    time.Time `gorm:"column:day_of_the_beast"` // set column name to `day_of_the_beast`
            	Age         int64     `gorm:"column:age_of_the_beast"` // set column name to `age_of_the_beast`
            }
            ```

        4.  修改列名的时候,并不会直接把列名修改,而是增加一个新的列,但是数据并不会同步过去!!!

    4.  关于嵌套

        1.  结构体是可以嵌套的,同样都会构建到数据库
        2.  gorm.Model  是一个gorm自带的结构体 里面有四个字段  id  创建时间  更新时间  删除时间

3.  查询

    1.  所有的情况都可以用结构体切片去接受值,哪怕只有一个也没有问题.

    2.  一般查询

        1.  方法的顺序是根据主键来定义,主键默认就是Id列,除非自定额外定义.
        2.  Frist (第一条),Take(随机一条),Last(最后一条),Find(所有记录)
        3.  Frist 还支持第二个参数 相当于where id = ??
        
    3.    where查询条件

          1.    where 并不会触发直接直接生成sql语句.也就是说可以在拼接其他条件(内联条件)

          2.    where方法支持两种写法,第一种是拼接sql的情况  第二种是直接用结构体或者map查询

                ```go
                	db.Where("name = ?", "stu01").First(&StuSlice)
                	fmt.Println(StuSlice)
                	db.Where(&Student{Name: "stu02"}).First(&StuSlice)
                	fmt.Println(StuSlice)
                ```

    4.    Not查询

          1.    类似于where

          2.    不同点在于in

                ```go
                db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
                //// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
                db.Not("name = ?", "jinzhu").First(&user)
                //// SELECT * FROM users WHERE NOT(name = "jinzhu");
                ```

    5.    or条件

          1.    拼接在where条件之后,用or方法连接 ,or方法使用与where差不多

                ```go
                	db.Where("age = ? or name = ?", 20, "stu08").Find(&StuSlice)
                	fmt.Println(StuSlice)
                    //相当于下面的情况
                	db.Where("age = ?", 20).Or("name = ?", "stu08").Find(&StuSlice)
                	fmt.Println(StuSlice)
                ```

    6.    内敛条件

          1.  作用与where相似,

    7.    获取结果如果没有的话构建一个(并不放入数据库)(仅支持 struct 和 map 条件)

    8.    获取结果如果没有的话构建一个(放入数据库)(仅支持 struct 和 map 条件)

    9.    对于8和9来说还有两种情况分别为 Attrs 和 Assign

          1.    #### Attrs   如果记录未找到，将使用参数创建 struct 和记录.

          2.    #### Assign  不管记录是否找到，都将参数赋值给 struct 并保存至数据库.

    10.    其他高级查询

           1.  查看官网

    11.  链式操作

         1.  orm的操作分为立即执行方法 和  普通的.
         2.  立即执行方法就是生成sql语句的方法.

    12.  多个立即执行方法

         1.  在 GORM 中使用多个立即执行方法时，后一个立即执行方法会复用前一个**立即执行方法**的条件 (不包括内联条件)

         ```go
         db.Where("name LIKE ?", "jinzhu%").Find(&users, "id IN (?)", []int{1, 2, 3}).Count(&count)
         //截至到Count之前 SELECT * FROM users WHERE name LIKE 'jinzhu%' AND id IN (1, 2, 3)
         // 此时Find 里面的id in (1,2,3)就不会再有
         //SELECT count(*) FROM users WHERE name LIKE 'jinzhu%'
         ```

    13.  生成函数

         1.  必须返回一个gorm.db的对象.

         2.  Scopes方法可以接受参数??

             ```go
             func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
               return db.Where("amount > ?", 1000)
             }
             
             func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
               return db.Where("pay_mode_sign = ?", "C")
             }
             
             func PaidWithCod(db *gorm.DB) *gorm.DB {
               return db.Where("pay_mode_sign = ?", "C")
             }
             
             func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
               return func (db *gorm.DB) *gorm.DB {
                 return db.Scopes(AmountGreaterThan1000).Where("status IN (?)", status)
               }
             }
             
             db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&orders)
             // 查找所有金额大于 1000 的信用卡订单
             
             db.Scopes(AmountGreaterThan1000, PaidWithCod).Find(&orders)
             // 查找所有金额大于 1000 的 COD 订单
             
             db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
             // 查找所有金额大于 1000 且已付款或者已发货的订单
             ```

4.  新增

    1.  新增之前看看数据库有没有当前记录

        ```go
        //返回是一个bool值,true就是有,false就是没有
        db.NewRecord(user) 
        ```

    2.  当前字段有默认值的时候,最好采用传递一个当前类型的指针类型,利用new 和 make实现

5.  更改

    1.  sava会更改全部字段.(数据库默认更新字段例如时间也会更改)

    2.  直接用save 会保存一条新的记录并不会更新.

    3.  只更新指定字段的时候,单个字段使用update 多个使用updates

    4.  对于映射来说,你可以选择更新某个值  用select  或者忽略某个值 Omit

    5.  更新单条记录的时候会触发钩子函数,比如有更新时间的字段,就会将次字段更新

    6.  批量更新不会触发钩子函数,但是可以通过map来手动更新次字段

        ```go
        db.Table("student").Debug().Where("id in (?)", []int{2, 3, 4}).Update(map[string]interface{}{"name": "stu", "age": 19, "updated_at": time.Now()})
        ```

    7.  更新某一列可以利用sql表达式

        ```go
        db.Model(&user).Update("age", gorm.Expr("age * ? + ?", 2, 100))
        ```

    8.  修改钩子函数的值????

        ```go
        func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
          if pw, err := bcrypt.GenerateFromPassword(user.Password, 0); err == nil {
            scope.SetColumn("EncryptedPassword", pw)
          }
        }
        ```

6.  删除

    1.  软删除和物理删除,软删除是将deleted_at的字段增加一个时间,并且正常查询的时候过滤出去.物理删除就是直接删除这个数据.
    
    2.  如果一个 model 有 `DeletedAt` 字段，他将自动获得软删除的功能,但是没有这个字段的话就没有软删除的功能,会直接删除
    
    3.  删除记录的时候需要确保你的结构体主键有值,如果没有的话就会全部删除数据.
    
    4.  当有软删除的时候,Unscoped方法可以查询到软删除的记录,同时利用Unscoped也可以完成物理删除
    
    5.  删除的时候不要自己去构建结构体.通过可以查询条件来进行删除.
    
        ```go
        db.Delete(&user)
        //// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;
        
        // 批量删除
        db.Where("age = ?", 20).Delete(&User{})
        //// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;
        
        // 查询记录时会忽略被软删除的记录
        db.Where("age = 20").Find(&user)
        //// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
        
        // Unscoped 方法可以查询被软删除的记录
        db.Unscoped().Where("age = 20").Find(&users)
        //// SELECT * FROM users WHERE age = 20;
        
        // Unscoped 方法可以物理删除记录
        db.Unscoped().Delete(&order)
        //// DELETE FROM orders WHERE id=10;
        ```


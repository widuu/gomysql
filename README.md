gomysql
[![Build Status](https://drone.io/github.com/widuu/gomysql/status.png)](https://drone.io/github.com/widuu/gomysql/latest)
===
**[官方网站](http://www.widuu.com)**

###介绍

 gomysql是基于go-sql-driver基础上开发的orm，这是一个轻量级的库。它会使数据库的增删改查变得非常容易。当然也是测试开发版，会一直优化和更新！请时刻关注我们

###安装
    
	go get github.com/go-sql-driver/mysql
    go get github.com/widuu/goini
	go get github.com/widuu/gomysql

###依赖的包

  mysql :[github.com/Go-SQL-Driver/MySQL](https://github.com/Go-SQL-Driver/MySQL)

  goini :[github.com/widuu/goini](https://github.com/widuu/goini)

###使用教程

>设置配置文件

	[database]
	username = mysql username
	password = mysql password
	hostname = mysql host
	charset =  mysql charset
	database = database name
	port = 	   mysql port

>**初始化配置**


    c, _ := gomysql.SetConfig("./conf/conf.ini") //根据配置文件，连接数据库

>查询数据

    t := c.SetTable("user")                      //设置要处理的表名
    data := t.FindOne()                          //查询表的一条数据，返回map[int]map[string]string格式
    gomysql.Print(data)                          //打印数据信息，返回如下截图的数据格式

![gomysql_print](http://yun.widuu.com/gomysql_print.png)

    data := t.Fileds("id", "password", "username").Where("id =12").FindOne() //Fileds 查询字段,Where where条件FindOne()查询一条
    //limit sql中的limit OrderBy sql中查询的OrderBy
    data  = c.SetTable("user").Fileds("id", "password", "username").Where("id>1").Limit(1, 5).OrderBy("id Desc").FindAll() 
    data  = t.FindAll() //查询所有数据，其中OrderBy() Limit() Where() Fileds()等设置条件
    gomysql.Print(data) //查询的数据都可以用Print()方法友好打印出来信息

>添加数据

     var value = make(map[string]interface{})  //设置map存储数据，map[key]value
     value["password"] = "mima3"    
     value["username"] = "xiaowei"
     value["id"] = 10
     t := c.SetTable("user")                   //设置增加的数据表
     t.SetPk("id")                             //设置主键自增的字段名称
     i,err := t.Insert(value)                  // 插入数据，返回增加个数和错误信息。返回最后增长的id和错误信息

>修改数据

    var value = make(map[string]interface{})
    value["password"] = "mima3"
    value["username"] = "xiaowei"
    n, err := c.SetTable("user").Where("id =5").Update(value)  //设置表，条件和修改的内容，返回影响条数和错误信息

>删除数据

    n, err := c.SetTable("user").Delete("id = 6")   //设置删除的表和数据，返回影响条数和错误信息

>关联操作

INNER JOIN 

	t := c.SetTable("user")
	//下面相当于sql语句，打印出是Select user.id,data.keywords,user.username,user.password from user INNER JOIN data ON user.id = data.id
	data := t.Fileds("user.id", "data.keywords", "user.username", "user.password").Join("data", "user.id = data.id").FindAll()
	fmt.Println(data)

LEFT JOIN

	 t.Fileds("user.id", "data.keywords", "user.username", "user.password").LeftJoin("data", "user.id = data.id").FindAll()
	 fmt.Println(data)

RIGHT JOIN
 	
	 t.Fileds("user.id", "data.keywords", "user.username", "user.password").RightJoin("data", "user.id = data.id").FindAll()
	 fmt.Println(data)

FULL　JOIN
	
     data := t.Fileds("user.id", "data.keywords", "user.username", "user.password").RightJoin("data", "user.id = data.id").FindAll()
	 fmt.Println(data)

>自定义sql语句

	// Query()方法 是自定义sql语句的
    
insert类型的数据

	n := t.Query("INSERT INTO user (`username`,`password`) VALUES ('xiaoweitest','xiaowei')") //返回最后增长的id

update|delete

	 //update
     n := c.Query("update user set username='ceshishenma' where id =17 ")
	 fmt.Println(n) //1 返回受影响行数
	
     //delete
	 n := c.Query("delete from user where id=16 ")
	 fmt.Println(n) //1 返回受影响行数

select 

    data := c.Query("select username,password from user")
	fmt.Println(data) //返回map[int]map[string]string 结构的所有数据

   

>关闭数据库

    c.DbClose()     //关闭数据库


---

##English Document

**[The Offcial Website](http://www.widuu.com)**

###Introduction

 Gomysql is based on the go - SQL - driver based on orm, this is a lightweight library.It allows the database to add delete very easily.Of course is also testing the development version, will continue to optimize and update!Please attention to us

###Install

    go get github.com/widuu/gomysql

###Rely on the package

  mysql :[github.com/Go-SQL-Driver/MySQL](https://github.com/Go-SQL-Driver/MySQL)

  goini :[github.com/widuu/goini](https://github.com/widuu/goini)

###Using the tutorial

>Set the configuration file

    [database]
    username = mysql username
    password = mysql password
    hostname = mysql host
    charset =  mysql charset
    database = database name
    port =     mysql port

>**Initialize the configuration**


    c, _ := gomysql.SetConfig("./conf/conf.ini") //According to the configuration files, connect to the database

>Query data

    t := c.SetTable("user")                      //set up the table name
    data := t.FindOne()                          //A data query，return map[int]map[string]string format
    gomysql.Print(data)                          //Print data information, and return the following screenshots data formats

![gomysql_print](http://yun.widuu.com/gomysql_print.png)

    data := t.Fileds("id", "password", "username").Where("id =12").FindOne() //Fileds Query field,Where where conditions FindOne() query a data
    //limit - sql limit,OrderBy - sql OrderBy
    data  = c.SetTable("user").Fileds("id", "password", "username").Where("id>1").Limit(1, 5).OrderBy("id Desc").FindAll() 
    data  = t.FindAll() //Query all the data, including OrderBy () Limit () the Where () Fileds () set conditions
    gomysql.Print(data) //Query data can use the Print () method of friendly printed information

>Insert data

     var value = make(map[string]interface{})  //Set the map data is stored, the map [key] the value
     value["password"] = "mima3"    
     value["username"] = "xiaowei"
     value["id"] = 10
     t := c.SetTable("user")                   //Set up to increase the data tables
     t.SetPk("id")                             //Set the primary key on the field name
     i,err := t.Insert(value)                  // Insert data, returns the number increase and error information.Returns the last growth id and error message

>Updata data

    var value = make(map[string]interface{})
    value["password"] = "mima3"
    value["username"] = "xiaowei"
    n, err := c.SetTable("user").Where("id =5").Update(value)  //Set the table, the condition and modify the content of the article returns affect the number and the error information

>Delete data

    n, err := c.SetTable("user").Delete("id = 6")   //Article, set the table and data delete, return to influence the number and the error information

>Associated operation

INNER JOIN 

	t := c.SetTable("user")
	
	//Equivalent to a SQL statement below, print out is 
	//Select user.id,data.keywords,user.username,user.password from user INNER JOIN data ON user.id = data.id

	data := t.Fileds("user.id", "data.keywords", "user.username", "user.password").Join("data", "user.id = data.id").FindAll()
	fmt.Println(data)

LEFT JOIN

	 t.Fileds("user.id", "data.keywords", "user.username", "user.password").LeftJoin("data", "user.id = data.id").FindAll()
	 fmt.Println(data)

RIGHT JOIN
 	
	 t.Fileds("user.id", "data.keywords", "user.username", "user.password").RightJoin("data", "user.id = data.id").FindAll()
	 fmt.Println(data)

FULL　JOIN
	
     data := t.Fileds("user.id", "data.keywords", "user.username", "user.password").RightJoin("data", "user.id = data.id").FindAll()
	 fmt.Println(data)

>自定义sql语句

	// Query()function Is a custom SQL statement
    
insert

	n := t.Query("INSERT INTO user (`username`,`password`) VALUES ('xiaoweitest','xiaowei')") //Return the id of the growth

update|delete

	 //update
     n := c.Query("update user set username='ceshishenma' where id =17 ")
	 fmt.Println(n) //1 Return the affected rows
	
     //delete
	 n := c.Query("delete from user where id=16 ")
	 fmt.Println(n) //1 Return the affected rows

select 

    data := c.Query("select username,password from user")
	fmt.Println(data) //return map[int]map[string]string All of the data structure

   

>Close the database

    c.DbClose()     //close the database











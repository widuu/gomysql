package main

import (
	"fmt"
	"github.com/widuu/gomysql"
)

func main() {
	c, err := gomysql.SetConfig("./conf.ini")
	if err != nil {
		fmt.Println(err)
	}

	t.Query("INSERT INTO user (`username`,`password`) VALUES ('xiaoweitest','xiaowei')")
	d := t.Query("select * from user")

	n := t.Query("INSERT INTO user (`username`,`password`) VALUES ('xiaoweitest','xiaowei')")
	fmt.Println(n)

	fmt.Println(t.Fileds("id,username").FindAll())

	data := t.Fileds("user.id", "data.keywords", "user.username", "user.password").Join("data", "user.id = data.id").FindAll()
	fmt.Println(data)

	n = c.Query("update user set username='ceshishenma' where id =17 ")
	fmt.Println(n)

	n = c.Query("delete from user where id=16 ")
	fmt.Println(n)

	data = c.Query("select username,password from user")
	fmt.Println(data)

	var value = make(map[string]interface{})
	value["username"] = "widuu"
	value["password"] = "widuu"
	_, err = t.Insert(value)
	fmt.Println(err)

	n, err = c.SetTable("user").Delete("id = 16")
	fmt.Println(n, err)

	var sss = make(map[string]interface{})
	sss["username"] = "widuuweb"
	r, err := t.Where("username = 'widuu'").Update(sss)
	fmt.Println(r, err)

	data = c.SetTable("user").SetParam([]string{"*"}).Where("id>1").Limit(1, 5).OrderBy("id Desc").FindAll()
	for _, v := range data {
		for k, value := range v {
			fmt.Println(k, value)
		}
	}
}

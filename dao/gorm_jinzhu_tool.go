package dao

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hello_grpc/dao/proto/service/test_orm"
)

func JianzhuTool() {
	db, err := gorm.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	defer db.Close()

	req := &test_orm.User{Id: 4,
		Sex:  1,
		Name: "dandy4",
	}
	ctx := context.Background()
	user, err := test_orm.DefaultCreateUser(ctx, req, db)
	fmt.Println("create1", user)

	req.Sex = 1
	res, err := test_orm.DefaultReadUser(ctx, req, db)
	fmt.Println(res, "|", err)
}

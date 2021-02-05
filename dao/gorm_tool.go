package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID    uint `gorm:"primarykey"`
	Ctime time.Time
	Code  string
	Price uint
}

type User struct {
	Name     string
	Age      uint
	Birthday time.Time
}

//message BaseInfo {
//option (gorm.opts).ormable = true;
//uint32 id = 1;
//int32 sex = 2;
//string name = 3;
//}

func GormTool() {

	dsn := "root:2501002@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	// 迁移 schema
	// db.AutoMigrate(&Product{})
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.CreateTable(&User{})
	result := db.Create(&user) // pass pointer of data to Create
	fmt.Println("user result", result)
	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	tx := db.First(&product, 1) // 根据整形主键查找
	fmt.Print("tx:", tx)
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)
}

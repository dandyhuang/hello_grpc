package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"
)

const dsn2 = "root:cole+123@tcp(192.168.207.160:3307)/go_db01?charset=utf8mb4&parseTime=True"

func TestMysqlJsonColumns(t *testing.T) {
	db, err := sql.Open("mysql", dsn2)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.ExecContext(context.Background(), `
CREATE TABLE IF NOT EXISTS test_model_json(
    id INTEGER PRIMARY KEY,
    json_v TEXT NOT NULL
)`)
	if err != nil {
		t.Fatal(err)
	}
	res, err := db.ExecContext(context.Background(), "INSERT INTO `test_model_json`(`id`, `json_v`) VALUES (?, ?)", 1, JsonColumn[Addresses]{Val: Addresses{Province: "广东", City: "东莞"}})
	if err != nil {
		t.Fatal(err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	if affected != 1 {
		t.Fatal(err)
	}
	fmt.Printf("insert successfully, insert rows:%d.\n", affected)

	row := db.QueryRowContext(context.Background(), "SELECT `id`,`json_v` FROM `test_model_json` LIMIT 1")
	if row.Err() != nil {
		t.Fatal(row.Err())
	}
	//tm := &User{
	//	Addresses: JsonColumn{
	//		Val: Address{},
	//	},
	//}

	tm := &User{}

	err = row.Scan(&tm.Id, &tm.Address)
	if err != nil {
		t.Fatal(err)
	}
	//assert.Equal(t, "changed", tm.Address)
	fmt.Println(tm.Address.Val.(Addresses).City) // Val是any类型，要做类型断言
	fmt.Println(tm.Address.Val.(Addresses).Province)

	//fmt.Println(tm.Address.Val.City)
	//fmt.Println(tm.Address.Val.Province)
}

type Addresses struct {
	Province string
	City     string
}

type User struct {
	Id      int64
	Address JsonColumn[Addresses]
}

type JsonColumn[T any] struct {
	Val   any
	Valid bool // 标记是不是NULL
}

func (j *JsonColumn[T]) Scan(src any) error {
	if src == nil {
		return nil
	}
	bs := src.([]byte)
	if len(bs) == 0 {
		return nil
	}
	err := json.Unmarshal(bs, &j.Val)
	j.Valid = err != nil
	return err
}

func (j JsonColumn[T]) Value() (driver.Value, error) {
	return json.Marshal(j.Val)
}

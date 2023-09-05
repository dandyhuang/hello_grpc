package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aj3423/aproto"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/reflect/protoreflect"
	"os"
	"reflect"
)

// 根据类型名称获取json数据
func findTypeByName(str string) (su protoreflect.ProtoMessage) {
	messageType1 := proto.MessageType(str)

	// 获取message实例
	rsp := reflect.New(messageType1.Elem())
	req := reflect.New(messageType1.Elem()).Interface().(protoreflect.ProtoMessage)
	fmt.Println("req:=====", req)
	descriptor := proto.MessageV2(req).ProtoReflect().Descriptor()
	fd := descriptor.Fields()
	for i := 0; i < fd.Len(); i++ {
		field := fd.Get(i)
		if field.Kind().String() == "uint64" {
			if field.IsList() {
				var num []uint64
				u := proto.Uint64(0)
				u2 := append(num, *u)
				rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(u2))
			} else {
				rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(proto.Uint64(0)))
			}

		} else if field.Kind().String() == "bool" {
			rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(proto.Bool(false)))
		} else if field.Kind().String() == "int32" {
			if field.IsList() {
				var num []int32
				u := proto.Int32(0)
				u2 := append(num, *u)
				rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(u2))
			} else {
				rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(proto.Int32(0)))
			}
		} else if field.Kind().String() == "int64" {
			if field.IsList() {
				var num []int64
				u := proto.Int64(0)
				u2 := append(num, *u)
				rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(u2))
			} else {
				rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(proto.Int64(0)))
			}
		} else if field.Kind().String() == "uint32" {
			if field.IsList() {
				var num []uint32
				u := proto.Uint32(0)
				u2 := append(num, *u)
				rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(u2))
			} else {
				rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(proto.Uint32(0)))
			}
		} else if field.Kind().String() == "message" {
			//if field.IsList() {
			//	t := rsp.Elem().FieldByName(string(field.Name())).Type().String()
			//	s := strings.Split(t, ".")
			//	pm2 := findTypeByName("hiproto." + s[1])
			//	if s[1] == "Party" {
			//		var u2 []*hiproto.Party
			//		p, _ := pm2.(*hiproto.Party)
			//		u2 = append(u2, p)
			//		rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(u2))
			//	} else if s[1] == "PartyUser" {
			//		var u2 []*hiproto.PartyUser
			//		p, _ := pm2.(*hiproto.PartyUser)
			//		u2 = append(u2, p)
			//		rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(u2))
			//	}
			//} else {
			//	t := rsp.Elem().FieldByName(string(field.Name())).Type().String()
			//	s := strings.Split(t, ".")
			//	// fmt.Println("s1", s[1])
			//	// fmt.Println("field.Name()", field.Name())
			//	pm2 := findTypeByName("hiproto." + s[1])
			//	rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(pm2))
			//}
		} else if field.Kind().String() == "bytes" {
			var num []uint8
			var u uint8 = 0
			u2 := append(num, u)
			rsp.Elem().FieldByName(string(field.Name())).Set(reflect.ValueOf(u2))
		}
		//fmt.Printf("Field %d: Name=%s, JSONName=%s, Number=%d, Type=%s\n", i, field.Name(), field.JSONName(), field.Number(), field.Kind())

	}
	pm := rsp.Interface().(protoreflect.ProtoMessage)
	//if str == "hiproto" {
	//	req := reflect.New(messageType1.Elem()).Interface().(*hiproto.GetMyPartyRequest)
	//	fmt.Print("req", req)
	//}
	return pm
}

// 根据类型名称获取解析详情数据
//func findTypeByDetial(str string, list []protocol.CmdIdDetial) (result []protocol.CmdIdDetial) {
//	messageType1 := proto.MessageType(str)
//	messageType2 := proto.MessageType(str)
//	// 获取message实例
//	rsp := reflect.New(messageType2.Elem())
//	req := reflect.New(messageType1.Elem()).Interface().(protoreflect.ProtoMessage)
//	descriptor := proto.MessageV2(req).ProtoReflect().Descriptor()
//	fd := descriptor.Fields()
//	for i := 0; i < fd.Len(); i++ {
//		var cmdId protocol.CmdIdDetial
//		field := fd.Get(i)
//		if field.HasOptionalKeyword() {
//			cmdId.Constraint = constants.Constraint_Optional
//		} else {
//			cmdId.Constraint = constants.Constraint_Repeated
//		}
//		cmdId.Field = field.JSONName()
//		cmdId.Type = fmt.Sprintf("%s", field.Kind())
//		if field.Kind().String() == "uint64" {
//			if field.IsList() {
//
//			} else {
//				cmdId.Default = "0"
//			}
//
//		} else if field.Kind().String() == "bool" {
//			cmdId.Default = "false"
//		} else if field.Kind().String() == "uint32" {
//			cmdId.Default = "0"
//		} else if field.Kind().String() == "message" {
//			t := rsp.Elem().FieldByName(changString(field.Name())).Type().String()
//			s := strings.Split(t, ".")
//			//fmt.Println("s1", s[1])
//			list = findTypeByDetial("hiproto."+s[1], list)
//
//		}
//		list = append(list, cmdId)
//		//fmt.Printf("Field %d: Name=%s, JSONName=%s, Number=%d, Type=%s\n", i, field.Name(), field.JSONName(), field.Number(), field.Kind())
//		if str == "hiproto" {
//			req := reflect.New(messageType1.Elem()).Interface().(*hiproto.GetMyPartyRequest)
//			fmt.Print("req", req)
//		}
//	}
//	return list
//}

func main() {
	var key, addr string
	flag.StringVar(&addr, "addr", "addrss", "配置文件")
	flag.StringVar(&key, "key", "key", "配置文件")
	var address []string
	flag.Parse()
	address = append(address, addr)
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: address,
		//MaxRedirects: c.config.MaxRetries,
		//ReadOnly:     c.config.ReadOnly,
		//Password:     c.config.Password,
		//MaxRetries:   c.config.MaxRetries,
		//DialTimeout:  c.config.DialTimeout,
		//ReadTimeout:  c.config.ReadTimeout,
		//WriteTimeout: c.config.WriteTimeout,
		//PoolSize:     c.config.PoolSize,
		//MinIdleConns: c.config.MinIdleConns,
		//IdleTimeout:  c.config.IdleTimeout,
	})
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error("redis ping", err)
		os.Exit(1)

	}
	v, _ := rdb.Get(ctx, key).Bytes()
	out, err := aproto.TryDumpEx(v, &aproto.ConsoleRenderer{})
	if err != nil {
		panic(err)
	}

	fmt.Println("value:", out)
}

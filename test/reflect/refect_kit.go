package refelect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)
type any interface {
}
func IterateFields(val any,  t *testing.T) {
	res, err := iterateFields(val, t)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range res {
		fmt.Println(k, v)
	}
}
func iterateFields(val any,  t *testing.T) (map[string]any, error) {
	if val == nil {
		return nil, errors.New("不能nil")
	}
	typ := reflect.TypeOf(val)
	refVal := reflect.ValueOf(val)

	if refVal.Kind() == reflect.Ptr {
		refVal = refVal.Elem()
	}
	if refVal.Kind() == reflect.Struct {
		numFiled := refVal.NumField()
		t.Log("filed:", numFiled, "name:", typ.Name(), " kind", refVal.Kind())
		//	t.Log("filed:", numFiled, " kind", refVal.Field(0))
		res := make(map[string]any, numFiled)
		for i := 0; i < numFiled; i++ {
			res[typ.Field(i).Name] = refVal.Field(i).Interface()
		}
		return res, nil
	}
	return nil, nil
}

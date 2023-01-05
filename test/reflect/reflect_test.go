package refelect

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func testPanic(path string) {
	if path == "" {
		panic("error")
	}
}

func TestReflectPainc(t *testing.T) {
	assert.Panicsf(t, func() {
		testPanic("1")
	}, "get")
	typ := reflect.TypeOf(&User{
		Name: "ton",
	})
	if typ.Kind() == reflect.Struct {
		n := typ.NumField()
		fmt.Println(n)
	} else if typ.Kind() == reflect.Ptr {
		fmt.Println("指针？？？")
	}
}

type User struct {
	Name string
}

func Test_iterateFields(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		want    map[string]any
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name:    "nil",
			val:     nil,
			wantErr: errors.New("不能nil"),
		},
		{
			name: "user",
			val: User{
				Name: "tom",
			},
			want: map[string]any{
				"Name": "tom",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("val:", tt.val)
			got, err := iterateFields(tt.val, t)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_Public_iterateFields(t *testing.T) {
	tests := []struct {
		Name    string
		Val     any
		Want    map[string]any
		WantErr error
	}{
		// TODO: Add test cases.
		{
			Name:    "nil",
			Val:     nil,
			WantErr: errors.New("不能nil"),
		},
		{
			Name: "user",
			Val: User{
				Name: "tom",
			},
			Want: map[string]any{
				"Name": "tom",
			},
			WantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Log("val:", tt.Val)
			got, err := iterateFields(tt.Val, t)
			assert.Equal(t, tt.WantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.Want, got)
		})
	}
}

package utils

import (
	"github.com/hhjpin/goutils/errors"
	"reflect"
)

//参数s只能是struct或[]struct类型
func ConvertStruct2Map(s interface{}) ([]map[string]interface{}, error) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		inLen := v.Len()
		res := make([]map[string]interface{}, inLen)
		for i := 0; i < inLen; i++ {
			v2 := v.Index(i)
			if v2.Kind() != reflect.Struct {
				return nil, errors.NewFormat(9, "子元素转换类型只能是struct")
			}
			t2 := v2.Type()
			numField := v2.NumField()
			res2 := make(map[string]interface{}, numField)
			for j := 0; j < numField; j++ {
				field := t2.Field(j)
				tag := field.Tag.Get("json")
				if tag == "" {
					return nil, errors.NewFormat(9, "字段"+field.Name+"缺少json的tag")
				}
				res2[tag] = v2.Field(j).Interface()
			}
			res[i] = res2
		}
		return res, nil
	} else if t.Kind() == reflect.Struct {
		numField := v.NumField()
		res := make(map[string]interface{}, numField)
		for i := 0; i < numField; i++ {
			field := t.Field(i)
			tag := field.Tag.Get("json")
			if tag == "" {
				return nil, errors.NewFormat(9, "字段"+field.Name+"缺少json的tag")
			}
			res[tag] = v.Field(i).Interface()
		}
		return []map[string]interface{}{res}, nil
	} else {
		return nil, errors.NewFormat(9, "转换类型只能是struct或[]struct")
	}
}

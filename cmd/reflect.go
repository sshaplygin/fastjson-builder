package main

import (
	"reflect"
	"strings"

	"github.com/valyala/fastjson"
)

const (
	jsonTagName = "json"
	skipField   = "-"
)

func GetReflectUserVal(arena *fastjson.Arena, data interface{}) *fastjson.Value {
	if data == nil {
		return arena.NewNull()
	}

	rValue := reflect.ValueOf(data)
	if rValue.Kind() != reflect.Array && rValue.Kind() != reflect.Slice &&
		rValue.Kind() != reflect.Pointer && rValue.Kind() != reflect.Struct &&
		rValue.Kind() != reflect.Map {
		panic("invalid input data type")
	}

	rType := reflect.TypeOf(data)
	if rType.Kind() == reflect.Pointer {
		rValue = rValue.Elem()
		rType = rValue.Type()
	}

	if rType.Kind() == reflect.Struct {
		return walk(arena, rType, rValue)
	}

	if rValue.Kind() != reflect.Array && rValue.Kind() != reflect.Slice &&
		rValue.Kind() != reflect.Map {
		panic("not supported data type")
	}

	return arena.NewNull()
}

func walk(arena *fastjson.Arena, rType reflect.Type, rValue reflect.Value) *fastjson.Value {
	obj := arena.NewObject()

	for i := 0; i < rType.NumField(); i++ {
		fieldType := rType.Field(i)

		tagValues := strings.Split(fieldType.Tag.Get(jsonTagName), ",")
		if fieldType.Anonymous || has(tagValues, skipField) {
			continue
		}

		fieldName := fieldType.Name
		if len(tagValues) > 0 && tagValues[0] != "" {
			fieldName = tagValues[0]
		}

		fieldValue := rValue.Field(i)

		val := arena.NewNull()
		switch fieldType.Type.Kind() {
		case reflect.Bool:
			val = newBoolVal(arena, fieldValue)
		case reflect.String:
			val = arena.NewString(fieldValue.String())
		case reflect.Slice, reflect.Array:
			val = arena.NewNull()

			if fieldValue.Len() != 0 {
				newArr := arena.NewArray()
				for i := 0; i < fieldValue.Len(); i++ {
					arrValue := fieldValue.Index(i)

					// var val *fastjson.Value
					switch arrValue.Kind() {
					// todo: support all type
					case reflect.String:
						val = arena.NewString(arrValue.String())
					case reflect.Bool:
						val = newBoolVal(arena, arrValue)
					case reflect.Int:
						val = arena.NewNumberInt(int(arrValue.Int()))
					default:
						panic("not implemented array value")
					}

					newArr.SetArrayItem(i, val)
				}

				val = newArr
			}
		case reflect.Struct:
			val = walk(arena, fieldValue.Type(), fieldValue)
		case reflect.Pointer:
			if !fieldValue.IsNil() {
				val = walk(arena, fieldValue.Elem().Type(), fieldValue.Elem())
			}
		default:
			panic("not implemented struct field value")
		}

		obj.Set(fieldName, val)
	}

	return obj
}

func newBoolVal(arena *fastjson.Arena, rValue reflect.Value) *fastjson.Value {
	if rValue.Bool() {
		return arena.NewTrue()
	}

	return arena.NewFalse()
}

func has[T comparable](vals []T, val T) bool {
	for _, v := range vals {
		if v == val {
			return true
		}
	}

	return false
}

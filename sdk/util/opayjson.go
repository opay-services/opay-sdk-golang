package util

/*
*auth by lijian
 */

import (
	"errors"
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"reflect"
	"strconv"
	"strings"
)

type strcutValueDesc struct {
	value     reflect.Value
	tag       reflect.StructTag
	fieldName string
}

func paresArrayValue(v reflect.Value) string {
	var ret string
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret += strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ret += strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		ret += fmt.Sprintf("%f", v.Float())
	case reflect.Bool:
		if v.Bool() {
			ret += "true"
		} else {
			ret += "false"
		}
	case reflect.String:
		ret += "\""
		ret += v.String()
		ret += "\""
	case reflect.Struct:
		ret += parseStruct(v.Interface())
	case reflect.Array, reflect.Slice:
		ret += parseArray(v)
	case reflect.Interface:
		return paresArrayValue(reflect.ValueOf(v))
	}
	return ret
}

func parseArray(value reflect.Value) string {
	var ret string
	ret += "["
	len := value.Len()
	for i := 0; i < len; i++ {
		v := value.Index(i)
		ret += paresArrayValue(v)
		if i != len-1 {
			ret += ","
		}
	}
	ret += "]"
	return ret
}

//if output empty return false
func parseStructValue(jsonValue strcutValueDesc) string {
	var ret string
	if jsonValue.value.Kind() == reflect.Ptr {
		if jsonValue.value.IsNil() {
			tag := jsonValue.tag.Get("json")
			if strings.Contains(tag, "omitempty") {

			} else {
				ret += fmt.Sprintf("\"%s\":", jsonValue.fieldName)
				ret += "null"
			}
			return ret
		} else {
			jsonValue.value = jsonValue.value.Elem()
		}
	}

	if (jsonValue.value.Kind() == reflect.Interface) {
		jsonValue.value = reflect.ValueOf(jsonValue.value.Interface())
		return parseStructValue(jsonValue)
	}

	value := jsonValue.value
	ret += fmt.Sprintf("\"%s\":", jsonValue.fieldName)
	switch value.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret += strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ret += strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		ret += fmt.Sprintf("%f", value.Float())
	case reflect.String:
		ret += "\""
		ret += value.String()
		ret += "\""
	case reflect.Bool:
		if value.Bool() {
			ret += "true"
		} else {
			ret += "false"
		}
	case reflect.Map:
		ret += parseHashMap(value.Interface())
	case reflect.Struct:
		ret += parseStruct(value.Interface())
	case reflect.Array, reflect.Slice:
		ret += parseArray(value)
	}
	return ret
}

func parseHashMap(hashMap interface{}) string {
	var ret string
	if reflect.ValueOf(hashMap).Kind() != reflect.Map {
		return ret
	}

	m := treemap.NewWithStringComparator()
	hashKeys := reflect.ValueOf(hashMap).MapKeys()
	for _, key := range hashKeys {
		v := reflect.ValueOf(hashMap).MapIndex(key)
		m.Put(fmt.Sprintf("%v", key), v.Interface())
	}

	ret += "{"
	keys := m.Keys()
	l := len(keys)
	for i := 0; i < l; i++ {
		ret += "\""
		ret += keys[i].(string)
		ret += "\""
		ret += ":"
		v, _ := m.Get(keys[i])
		value := reflect.ValueOf(v)

		switch value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ret += strconv.FormatInt(value.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			ret += strconv.FormatUint(value.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			ret += fmt.Sprintf("%f", value.Float())
		case reflect.String:
			ret += "\""
			ret += value.String()
			ret += "\""
		case reflect.Bool:
			if value.Bool() {
				ret += "true"
			} else {
				ret += "false"
			}
		case reflect.Struct:
			ret += parseStruct(value.Interface())
		case reflect.Array, reflect.Slice:
			ret += parseArray(value)
		}

		if i != l -1  {
			ret += ","
		}
	}

	ret += "}"
	return ret
}

func parseStruct(input interface{}) string {
	var ret string
	m := treemap.NewWithStringComparator()
	s := reflect.TypeOf(input)
	v := reflect.ValueOf(input)
	for idx := 0; idx < s.NumField(); idx++ {
		st := s.Field(idx)
		sv := v.FieldByName(st.Name)
		desc := strcutValueDesc{tag: st.Tag, value: sv}
		jsonTag := st.Tag.Get("json")
		if jsonTag != "" {
			i := strings.Index(jsonTag, ",")
			if i == -1 {
				desc.fieldName = jsonTag
			} else {
				desc.fieldName = jsonTag[:i]
			}
		} else {
			desc.fieldName = st.Name
		}
		m.Put(desc.fieldName, desc)
	}

	ret += "{"
	keys := m.Keys()

	firstFalg := false
	for i := 0; i < len(keys); i++ {
		value, _ := m.Get(keys[i])
		r := parseStructValue(value.(strcutValueDesc))
		if len(r) != 0 {
			if firstFalg {
				ret += ","
				ret += r
				firstFalg = true
			} else {
				ret += r
				firstFalg = true
			}
		}
	}
	ret += "}"
	return ret
}

func OpayJsonMarshal(s interface{}) (str string, err error) {
	switch reflect.TypeOf(s).Kind() {
	case reflect.Ptr:
		if reflect.ValueOf(s).IsNil() {
			err = errors.New("ptr is nil")
			return
		}
		switch reflect.TypeOf(s).Elem().Kind() {
		case reflect.Struct:
			return parseStruct(reflect.ValueOf(s).Elem().Interface()), nil
		case reflect.Map:
			return parseHashMap(reflect.ValueOf(s).Elem().Interface()), nil
		default:
			err = errors.New("input must struct or struct ptr")
			return
		}
	case reflect.Struct:
		return parseStruct(s), nil

	case reflect.Map:
		return parseHashMap(s), nil
	default:
		err = errors.New("input must struct or struct ptr or map")
		return
	}
}

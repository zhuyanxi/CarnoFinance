package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/zhuyanxi/CarnoFinance/domain"
)

func SetStructFieldValuesByJsonTags(structObj *domain.StockDailyPrice, jsonTags []string, values []any) error {
	structValue := reflect.ValueOf(structObj).Elem()
	for i, tagName := range jsonTags {
		fieldName := getFieldNameByJsonTag(structObj, tagName)
		field := structValue.FieldByName(fieldName)
		if !field.IsValid() {
			return fmt.Errorf("field %s with tag %s not found", fieldName, tagName)
		}
		// data field in this app only have float64 and string type
		switch field.Kind() {
		case reflect.Float64:
			fv, ok := values[i].(float64)
			if !ok {
				return fmt.Errorf("invalid values %v", values[i])
			}
			field.SetFloat(fv)
		case reflect.String:
			fv, ok := values[i].(string)
			if !ok {
				return fmt.Errorf("invalid values %v", values[i])
			}
			field.SetString(fv)
		default:
			return fmt.Errorf("type not support: %s", field.Kind())
		}
	}

	return nil
}

func getFieldNameByJsonTag(structObj any, tag string) string {
	structValue := reflect.TypeOf(structObj).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		// json:"ts_code,omitempty"
		jTag := field.Tag.Get("json")
		if strings.Split(jTag, ",")[0] == tag {
			return field.Name
		}
	}
	return ""
}

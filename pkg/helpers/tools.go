// Package helpers 存放辅助方法
package helpers

import (
	"fmt"
	mathRand "math/rand"
	"reflect"
	"time"
)

func Error2String(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// RandomNumber 生成长度为 length 随机数字字符串
func RandomNumber(length int) string {
	rand := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	letters := "01234567890123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// FirstElement 安全地获取 args[0]，避免 panic: runtime error: index out of range
func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

// RandomCharString 生成长度为 length 的随机字符串
//
// Note:
//
//	没有字符l
func RandomCharString(length int) string {
	rand := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	letters := "abcdefghijkmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomString 生成长度为 length 的随机字符串:数字+字符
//
// Note:
//
//	没有字符l
func RandomString(length int) string {
	rand := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	letters := "abcdefghijkmno01234567890123456789pqrstuvwxyzABCDEF01234567890123456789GHIJKLMNOPQRSTUVWXYZ01234567890123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

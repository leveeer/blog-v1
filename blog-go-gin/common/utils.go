package common

import (
	"errors"
	"fmt"
	"reflect"
)

// SimpleCopyFields
//参数传递时，第二个参数使用指针还是实例请自行斟酌，第一个参数必须是指针，涉及的字段必须是对外的
//需要注意的是，该拷贝方法为浅拷贝，换句话说，如果说对象内嵌套有其他的引用类型如Slice,Map等，
//用此方法完成拷贝后，源对象中的引用类型属性内容发生了改变，该对象对应的属性中内容也会改变。
func SimpleCopyFields(dst, src interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针，.Elem()类似于*ptr的操作返回指针指向的地址反射类型
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	FieldNums := dstType.NumField()

	for i := 0; i < FieldNums; i++ {
		// 属性
		field := dstType.Field(i)
		// 待填充属性值
		fieldValue := srcValue.FieldByName(field.Name) // 无效，说明src没有这个属性 || 属性同名但类型不同
		if !fieldValue.IsValid() || field.Type != fieldValue.Type() {
			continue
		}
		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(fieldValue)
		}
	}

	return nil
}

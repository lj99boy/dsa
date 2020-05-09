// 用反射实现泛型
package reflect

import (
	"reflect"
)

type Animal struct {
	v interface{}
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (a Animal) GetCatName() string {
	return a.v.(Cat).name
}

func (a Animal) GetDogName() string {
	return a.v.(Dog).name
}

func (a Animal) GetAnimalName() string {
	str := ""
	getValue := reflect.ValueOf(a)
	getType := reflect.TypeOf(a)

	for i := 0; i < getValue.NumField(); i++ {
		field := getType.Field(i)
		fieldVal := getValue.Field(i).Interface()
		if field.Name == "name"{
			str = fieldVal.(string)
		}
	}
	return str
}

package main

/*
使用反射来遍历结构体的字段， 调用结构体的方法， 并获取结构体标签的值
*/


import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name   string `json:"monsterName"`
	Age    int    `json:"monsterAge"`
	Gender string `json:"monsterGender"`
	Score  int    `json:"monsterScore"`
}

func (m *Monster) Print() {
	fmt.Printf("姓名: %s  年龄: %d, 性别: %s, 分数: %d\n", m.Name, m.Age, m.Gender, m.Score)
}

func (m *Monster) Set(name string, age int, gender string, score int) (M *Monster) {
	m.Name = name
	m.Age = age
	m.Gender = gender
	m.Score = score

	return m
}

func (m *Monster) Sum(n1, n2 int) int {
	return n1 + n2
}

func New(name string, age int, gender string, score int) *Monster {
	return &Monster{
		Name:   name,
		Age:    age,
		Gender: gender,
		Score:  score,
	}
}

func ReflectTest(b interface{}) {
	t := reflect.TypeOf(b)

	//if t.Kind() == reflect.Ptr {
	//	t = t.Elem()
	//}
	v := reflect.ValueOf(b)

	//if v.Kind() == reflect.Ptr {
	//	v = v.Elem()
	//}

	// 修改指定字段的值，也可以使用v.Elem().Field(0).SetString("zhenping1")修改
	v.Elem().FieldByName("Name").SetString("zhenping1")
	fmt.Println(v.Elem().FieldByName("Name")) // 执行结果: zhenping1

	FieldNum := t.Elem().NumField()
	for i := 0; i < FieldNum; i++ {
		// 获取字段的名称
		fmt.Printf("v的字段名称: %v\n", t.Elem().Field(i).Name) // v的字段名称: Name
		// 获取字段的tag
		fmt.Printf("v的字段Json Tag: %v\n", t.Elem().Field(i).Tag.Get("json")) // v的字段Json Tag: monsterName
		// 获取字段的值
		fmt.Printf("t的值信息%v\n", v.Elem().Field(i).Interface()) // t的值信息zhenping
	}

	MethodNum := v.NumMethod()
	fmt.Printf("m对象的方法数量: %d\n", MethodNum) // m对象的方法数量: 2

	// 调用结构体的方法 0表示第0个方法，call(nil) : 表示调用并执行，nil表示参数为空
	v.Method(0).Call(nil)

	// 调用结构体有参数的方法
	// 调用Sum方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(30))
	params = append(params, reflect.ValueOf(20))
	res := v.MethodByName("Sum").Call(params)
	fmt.Println(res[0].Int()) // 执行结果 50

	// 调用set方法

	var params1 []reflect.Value
	params1 = append(params1, reflect.ValueOf("zhenping"))
	params1 = append(params1, reflect.ValueOf(20))
	params1 = append(params1, reflect.ValueOf("male"))
	params1 = append(params1, reflect.ValueOf(99))

	res = v.Method(1).Call(params1) // 此处的1表示, 结构体的第2个方法，方案按ASCII排序
	fmt.Println(res[0])             // 执行结果: &{zhenping 20 male 99}, 返回结果是reflect.value类型。0表示当前返回结果只有一个，如果有多个，使用下标即可取。

}

func main() {
	m := New("zhenping", 18, "male", 100)
	ReflectTest(m)
	fmt.Println(m) // &{zhenping 20 male 99}
}

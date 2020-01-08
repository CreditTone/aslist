package aslist

import (
	"fmt"
	"testing"
)

type A struct {
	Name string
}

type B struct {
	Age int
}

func (self *B) Compare(a, b interface{}) bool {
	obj_a := a.(*B)
	obj_b := b.(*B)
	if obj_a.Age > obj_b.Age {
		return true
	}
	return false
}

func TestAsListSortWithCompareFunc(t *testing.T) {
	asList := NewAsList()
	asList.Push(A{Name: "我的名字好长"})
	asList.Push(A{Name: "我名长"})
	asList.Push(A{Name: "我的名字好长阿"})
	asList.Push(A{Name: "我名字长"})
	asList.Push(A{Name: "我的名字长"})

	t.Log("第一次遍历")
	asList.Range(func(index int, item interface{}) {
		t.Log(index, item)
	})
	asList.SortWithCompareFunc(func(a, b interface{}) bool {
		struct_a := a.(A)
		struct_b := b.(A)
		if len(struct_a.Name) > len(struct_b.Name) {
			return true
		}
		return false
	})
	t.Log("排序后遍历")
	asList.Range(func(index int, item interface{}) {
		t.Log(index, item)
	})
}

//测试使用指针结构
func TestStructPoint(t *testing.T) {
	asList := NewAsList()
	asList.Push(&A{Name: "我的名字好长"})
	asList.Push(&A{Name: "我名长"})
	asList.Push(&A{Name: "我的名字好长阿"})
	asList.Push(&A{Name: "我名字长"})
	asList.Push(&A{Name: "我的名字长"})

	t.Log("第一次遍历")
	asList.Range(func(index int, item interface{}) {
		t.Log(index, item)
	})
	asList.SortWithCompareFunc(func(a, b interface{}) bool {
		struct_a := a.(*A)
		struct_b := b.(*A)
		if len(struct_a.Name) > len(struct_b.Name) {
			return true
		}
		return false
	})
	t.Log("排序后遍历")
	asList.Range(func(index int, item interface{}) {
		t.Log(index, item)
	})
}

//测试自动排序
func TestAutoSort(t *testing.T) {
	asList := NewAsList()
	asList.Push(&B{Age: 121})
	asList.Push(&B{Age: 120})
	asList.Push(&B{Age: 23})
	asList.Push(&B{Age: 150})
	asList.Push(&B{Age: 69})
	t.Log("第一次遍历")
	asList.Range(func(index int, item interface{}) {
		t.Log(index, item)
	})
	asList.Sort()
	t.Log("排序后遍历")
	asList.Range(func(index int, item interface{}) {
		t.Log(index, item)
	})
}

//测试队列
func TestQueue(t *testing.T) {
	asList := NewAsList()
	asList.Push(&B{Age: 121})
	asList.Push(&B{Age: 120})
	asList.Push(&B{Age: 23})
	asList.Push(&B{Age: 150})
	asList.Push(&B{Age: 69})
	t.Log("第一次遍历")
	asList.Range(func(index int, item interface{}) {
		t.Log(index, item)
	})
	asList.Sort()
	t.Log("排序后遍历")
	asList.Range(func(index int, item interface{}) {
		t.Log(index, item)
	})
	for i := 1; i < 10; i++ {
		t.Log(fmt.Sprintf("第%d次pop:", i), asList.Pop())
	}
}

func TestMarshalJson(t *testing.T) {
	asList := NewAsList()
	asList.Push(&B{Age: 121})
	asList.Push(&B{Age: 120})
	asList.Push(&B{Age: 23})
	asList.Push(&B{Age: 150})
	asList.Push(&B{Age: 69})
	t.Log("第一次MarshalJson")
	t.Log(string(asList.MarshalJson()))
	asList.Sort()
	t.Log("排序后MarshalJson")
	t.Log(string(asList.MarshalJson()))
	t.Log("RightPop:", asList.RightPop())
	t.Log("LeftPop:", asList.LeftPop())
	t.Log(string(asList.MarshalJson()))
}

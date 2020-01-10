package aslist

import (
	"encoding/json"
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
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
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
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
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
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
	asList.SortWithCompareFunc(func(a, b interface{}) bool {
		struct_a := a.(*A)
		struct_b := b.(*A)
		if len(struct_a.Name) > len(struct_b.Name) {
			return true
		}
		return false //如果要中断遍历，请返回true
	})
	t.Log("排序后遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
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
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
	asList.Sort()
	t.Log("排序后遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
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
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
	asList.Sort()
	t.Log("排序后遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
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

func TestUnMarshalJson(t *testing.T) {
	list := []interface{}{}
	list = append(list, &B{Age: 121}, &B{Age: 120}, &B{Age: 23}, &B{Age: 150}, &B{Age: 69})
	listJson, _ := json.Marshal(list)
	//加载json
	asList := NewAsList()
	asList.Push(&B{Age: 89})
	//true表示需要追加 unSerialize反序列化的函数
	err := asList.UnmarshalJson(listJson, true, func(itemData []byte) interface{} {
		item := new(B)
		json.Unmarshal(itemData, item)
		return item
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("UnmarshalJson后遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
}

func TestClear(t *testing.T) {
	asList := NewAsList()
	asList.Push(&B{Age: 121})
	asList.Push(&B{Age: 120})
	asList.Push(&B{Age: 23})
	asList.Push(&B{Age: 150})
	asList.Push(&B{Age: 69})
	t.Log("第一次遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
	asList.ClearTargets(func(index int, item interface{}) bool {
		struct_a := item.(*B)
		//把小于Age小于100的全部清除掉
		if struct_a.Age < 100 {
			return true
		}
		return false
	})
	t.Log("ClearTargets后遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
	asList.Clear() //清除所有元素
	t.Log("Clear后元素个数为", asList.Length())
}

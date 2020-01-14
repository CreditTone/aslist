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

//定义一个BInterface来实现指针和结构体的多态
type BInterface interface {
	GetAge() int
}

//实现BInterface定义的方法，注意(self B)不能是(self *B)
func (self B) GetAge() int {
	return self.Age
}

//func (self *B) Compare(a, b interface{}) bool {
//	obj_a := a.(*B)
//	obj_b := b.(*B)
//	if obj_a.Age > obj_b.Age {
//		return true
//	}
//	return false
//}

//使用多态接口实现类型的统一转换
func (self *B) Compare(a, b interface{}) bool {
	obj_a := a.(BInterface)
	obj_b := b.(BInterface)
	if obj_a.GetAge() > obj_b.GetAge() {
		return true
	}
	return false
}

func TestSmartGanerateUniqueId(t *testing.T) {
	b1 := B{Age: 1}
	b2 := &B{Age: 1}
	b3 := B{Age: 2}
	b4 := B{Age: 1}
	t.Log("SmartGanerateUniqueId,不忽略指针和结构体类型生成唯一id")
	//不忽略指针和结构体类型生成唯一id
	t.Log("b1", SmartGanerateUniqueId(b1))
	t.Log("b2", SmartGanerateUniqueId(b2))
	t.Log("b3", SmartGanerateUniqueId(b3))
	t.Log("b4", SmartGanerateUniqueId(b4))

	t.Log("SmartGanerateUniqueIdWithIgnorePoint,忽略指针和结构体类型生成唯一id")
	//忽略指针和结构体类型生成唯一id
	t.Log("b1", SmartGanerateUniqueIdWithIgnorePoint(b1))
	t.Log("b2", SmartGanerateUniqueIdWithIgnorePoint(b2))
	t.Log("b3", SmartGanerateUniqueIdWithIgnorePoint(b3))
	t.Log("b4", SmartGanerateUniqueIdWithIgnorePoint(b4))
}

func TestGanerateUniqueId(t *testing.T) {
	asList := NewAsList()
	//设置GanerateUniqueId函数，asList将会做唯一性校验
	asList.GanerateUniqueId = func(i interface{}) string {
		bi := i.(BInterface)
		//假设以Age生成唯一id
		return fmt.Sprintf("%d", bi.GetAge())
	}
	asList.Push(&B{Age: 121})
	asList.Push(&B{Age: 120})
	asList.Push(&B{Age: 23})
	asList.Push(&B{Age: 150})
	asList.Push(&B{Age: 69})
	t.Log("测试唯一'性'功能😄，第一次遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
	asList.Push(&B{Age: 123})
	asList.Push(&B{Age: 120}) //重复
	asList.Push(&B{Age: 23})  //重复
	asList.Push(&B{Age: 150}) //重复
	asList.Push(&B{Age: 96})
	t.Log("测试唯一'性'功能😄，第二次遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
}

func TestGanerateUniqueIdWithSmartGanerateUniqueId(t *testing.T) {
	asList := NewAsList()
	//设置GanerateUniqueId函数，asList将会做唯一性校验
	asList.GanerateUniqueId = SmartGanerateUniqueId //aslist.SmartGanerateUniqueId，这里我放的全是指针类型所以不必用SmartGanerateUniqueIdWithIgnorePoint。
	asList.Push(&B{Age: 121})
	asList.Push(&B{Age: 120})
	asList.Push(&B{Age: 23})
	asList.Push(&B{Age: 150})
	asList.Push(&B{Age: 69})
	t.Log("测试唯一'性'功能😄，第一次遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
	asList.Push(&B{Age: 123})
	asList.Push(&B{Age: 120}) //重复
	asList.Push(&B{Age: 23})  //重复
	asList.Push(&B{Age: 150}) //重复
	asList.Push(&B{Age: 96})
	t.Log("测试唯一'性'功能😄，第二次遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
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

func TestPolymorphic(t *testing.T) {
	asList := NewAsList()
	asList.Push(&B{Age: 121})
	asList.Push(B{Age: 120})
	asList.Push(&B{Age: 23})
	asList.Push(B{Age: 150})
	asList.Push(B{Age: 69})
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

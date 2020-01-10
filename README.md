# aslist

**aslist(A Sorted List)是golang语言实现的线程安全可排序的list。便捷的调用方式，使开发者快速入门使用。**

**为什么要设计aslist？**
   答：如果你是从java转golang开发，你就会发现golang中对数组(array)、切片(slice)的封装比较原生。这对开发者来说过去java日常对List方便操作放在golang中是非常痛苦的。如排序(sort)、栈(stack)操作、先进先出(FIFO)、左进右出(LIRI)......本人深耕java多年深刻体会到你的痛苦，所以借鉴java对list体验设计的思路封装了这个轻量级的aslist。使你在golang中能找回java的感觉。

## 快速开始

	go get -u github.com/CreditTone/aslist

### 导入
```golang
import(
	"github.com/CreditTone/aslist"
)
```

### 定义结构
```golang
type A struct {
	Name string
}
```

### 创建aslist
```golang
	asList := NewAsList()
```

### push元素
```golang
	asList.Push(A{Name: "我的名字好长"})
	asList.Push(A{Name: "我名长"})
	asList.Push(A{Name: "我的名字好长阿"})
	asList.Push(A{Name: "我名字长"})
	asList.Push(A{Name: "我的名字长"})
```

### 使用SortWithCompareFunc传入自定义函数排序
```golang
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
			return true //需要a 排在 b前面返回true
		}
		return false
	})
	t.Log("排序后遍历")
	asList.Range(func(index int, item interface{}) bool {
    		t.Log(index, item)
    		return false //如果要中断遍历，请返回true
    	})
```
输出结果：
```text
=== RUN   TestAsListSortWithCompareFunc
---  PASS: TestAsListSortWithCompareFunc (0.00s)
    aslist_test.go:33: 第一次遍历
    aslist_test.go:35: 0 {我的名字好长}
    aslist_test.go:35: 1 {我名长}
    aslist_test.go:35: 2 {我的名字好长阿}
    aslist_test.go:35: 3 {我名字长}
    aslist_test.go:35: 4 {我的名字长}
    aslist_test.go:45: 排序后遍历
    aslist_test.go:47: 0 {我的名字好长阿}
    aslist_test.go:47: 1 {我的名字好长}
    aslist_test.go:47: 2 {我的名字长}
    aslist_test.go:47: 3 {我名字长}
    aslist_test.go:47: 4 {我名长}
PASS
```

### 使用指针结构做元素也可以
```golang
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
			return true //需要a 排在 b前面返回true
		}
		return false
	})
	t.Log("排序后遍历")
	asList.Range(func(index int, item interface{}) {
		t.Log(index, item)
	})
```
输出结果：
```text
=== RUN   TestStructPoint
--- PASS: TestStructPoint (0.00s)
    aslist_test.go:60: 第一次遍历
    aslist_test.go:62: 0 &{我的名字好长}
    aslist_test.go:62: 1 &{我名长}
    aslist_test.go:62: 2 &{我的名字好长阿}
    aslist_test.go:62: 3 &{我名字长}
    aslist_test.go:62: 4 &{我的名字长}
    aslist_test.go:72: 排序后遍历
    aslist_test.go:74: 0 &{我的名字好长阿}
    aslist_test.go:74: 1 &{我的名字好长}
    aslist_test.go:74: 2 &{我的名字长}
    aslist_test.go:74: 3 &{我名字长}
    aslist_test.go:74: 4 &{我名长}
PASS
```

### 数据结构实现Compare(a, b interface{}) bool方法可直接调用Sort排序
```golang
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
```
调用Sort方法排序
```golang
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
```
看到这里有没有java ArrayList.sort的感脚.输出如下
```text
=== RUN   TestAutoSort
--- PASS: TestAutoSort (0.00s)
    aslist_test.go:86: 第一次遍历
    aslist_test.go:88: 0 &{121}
    aslist_test.go:88: 1 &{120}
    aslist_test.go:88: 2 &{23}
    aslist_test.go:88: 3 &{150}
    aslist_test.go:88: 4 &{69}
    aslist_test.go:91: 排序后遍历
    aslist_test.go:93: 0 &{150}
    aslist_test.go:93: 1 &{121}
    aslist_test.go:93: 2 &{120}
    aslist_test.go:93: 3 &{69}
    aslist_test.go:93: 4 &{23}
PASS
```

### 队列操作
```golang
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
```
输出结果如下：
```text
=== RUN   TestQueue
--- PASS: TestQueue (0.00s)
    aslist_test.go:105: 第一次遍历
    aslist_test.go:107: 0 &{121}
    aslist_test.go:107: 1 &{120}
    aslist_test.go:107: 2 &{23}
    aslist_test.go:107: 3 &{150}
    aslist_test.go:107: 4 &{69}
    aslist_test.go:110: 排序后遍历
    aslist_test.go:112: 0 &{150}
    aslist_test.go:112: 1 &{121}
    aslist_test.go:112: 2 &{120}
    aslist_test.go:112: 3 &{69}
    aslist_test.go:112: 4 &{23}
    aslist_test.go:115: 第1次pop: &{150}
    aslist_test.go:115: 第2次pop: &{121}
    aslist_test.go:115: 第3次pop: &{120}
    aslist_test.go:115: 第4次pop: &{69}
    aslist_test.go:115: 第5次pop: &{23}
    aslist_test.go:115: 第6次pop: <nil>
    aslist_test.go:115: 第7次pop: <nil>
    aslist_test.go:115: 第8次pop: <nil>
    aslist_test.go:115: 第9次pop: <nil>
PASS
```

### 序列化为json
```golang
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
```
输出结果如下：
```text
=== RUN   TestMarshalJson
--- PASS: TestMarshalJson (0.00s)
    aslist_test.go:126: 第一次MarshalJson
    aslist_test.go:127: [{"Age":121},{"Age":120},{"Age":23},{"Age":150},{"Age":69}]
    aslist_test.go:129: 排序后MarshalJson
    aslist_test.go:130: [{"Age":150},{"Age":121},{"Age":120},{"Age":69},{"Age":23}]
    aslist_test.go:131: RightPop: &{23}
    aslist_test.go:132: LeftPop: &{150}
    aslist_test.go:133: [{"Age":121},{"Age":120},{"Age":69}]
PASS
```

### 反序列化json
```golang
	list := []interface{}{}
        list = append(list, &B{Age: 121}, &B{Age: 120}, &B{Age: 23}, &B{Age: 150}, &B{Age: 69})
    	listJson, _ := json.Marshal(list)
    	//加载json
    	asList := NewAsList()
    	asList.UnmarshalJson(listJson)
    	t.Log("UnmarshalJson后遍历")
    	asList.Range(func(index int, item interface{}) bool {
    		t.Log(index, item)
    		return false //如果要中断遍历，请返回true
    	})
```
输出结果如下：
```text
=== RUN   TestUnMarshalJson
--- PASS: TestUnMarshalJson (0.00s)
    aslist_test.go:152: UnmarshalJson后遍历
    aslist_test.go:154: 0 map[Age:121]
    aslist_test.go:154: 1 map[Age:120]
    aslist_test.go:154: 2 map[Age:23]
    aslist_test.go:154: 3 map[Age:150]
    aslist_test.go:154: 4 map[Age:69]
PASS
```

### ClearTargets和Clear
```golang
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
```
输出结果如下：
```text
=== RUN   TestClear
--- PASS: TestClear (0.00s)
    aslist_test.go:166: 第一次遍历
    aslist_test.go:168: 0 &{121}
    aslist_test.go:168: 1 &{120}
    aslist_test.go:168: 2 &{23}
    aslist_test.go:168: 3 &{150}
    aslist_test.go:168: 4 &{69}
    aslist_test.go:179: ClearTargets后遍历
    aslist_test.go:181: 0 &{121}
    aslist_test.go:181: 1 &{120}
    aslist_test.go:181: 2 &{150}
    aslist_test.go:185: Clear后元素个数为 0
PASS
```
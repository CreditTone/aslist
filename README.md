# aslist

**aslist(A Sorted List)是golang语言实现的线程安全可排序的list。便捷的调用方式，使开发者快速入门使用。**

**前言**
>近来有喷子喷我，说golang有排序框架何苦要自己造轮子。我只想说中国的拿来主义思想就是导致今天中国为何终始没有自主研发的操作系统，没有自己的芯片。你跟我说这句话就像问"已经有了Java为什么还要学Golang、为什么男人有了老婆还想有小三"。因为需求嘛。各种各样的需求，有人在乎性能，有人在乎便捷，有人在乎二者都在乎。你要是觉得我造的轮子不好，你自己就去造一个更好的。不要在这里说一些显得你很浅薄的话。aslist就是给在乎便捷、不想了解底层算法的人使用。不喜欢你可以走。


**为什么要设计aslist？**
>如果你是从java转golang开发，你就会发现golang中对数组(array)、切片(slice)的封装比较原生。这对开发者来说过去java日常对List方便操作放在golang中是非常痛苦的。如排序(sort)、栈(stack)操作、先进先出(FIFO)、左进右出(LIRI)......本人深耕java多年深刻体会到你的痛苦，所以借鉴java对list体验设计的思路封装了这个轻量级的aslist。使你在golang中能找回java的感觉。

**为什么Range、ClearTargets、Pop不能像Java一样使用泛型？**
>没有办法，当前golang的泛型还没有像Java一样强大。先忍忍，估计马上就要出来了。

## 快速开始

	go get -u github.com/CreditTone/aslist
	
## 如果你使用gomod管理依赖

	go get -u github.com/CreditTone/aslist@master

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

### 1.创建aslist
```golang
	asList := NewAsList()
```

### 2.push元素
```golang
	asList.Push(A{Name: "我的名字好长"})
	asList.Push(A{Name: "我名长"})
	asList.Push(A{Name: "我的名字好长阿"})
	asList.Push(A{Name: "我名字长"})
	asList.Push(A{Name: "我的名字长"})
```

### 3.使用SortWithCompareFunc传入自定义函数排序
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

### 4.使用指针结构做元素也可以
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

### 5.数据结构实现Compare(a, b interface{}) bool方法可直接调用Sort排序
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

### 6.当你的AsList当中指针和结构体同时存在怎么办？这里教你一个技巧巧妙的解决这个问题
```golang
//重新定义结构
type B struct {
	Age int
}

//定义一个BInterface来实现指针和结构体的多态
type BInterface interface {
	GetAge() int
}

//实现BInterface定义的方法，注意必须是(self B)不能是(self *B)。否则不能做到多态调用
func (self B) GetAge() int {
	return self.Age
}


//使用多态接口实现类型的统一转换
func (self *B) Compare(a, b interface{}) bool {
	obj_a := a.(BInterface)
	obj_b := b.(BInterface)
	if obj_a.GetAge() > obj_b.GetAge() {
		return true
	}
	return false
}
```
#### 6.1测试多态设计思路
```golang
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
```
输出结果如下：
```text
=== RUN   TestPolymorphic
--- PASS: TestPolymorphic (0.00s)
    aslist_test.go:134: 第一次遍历
    aslist_test.go:136: 0 &{121}
    aslist_test.go:136: 1 {120}
    aslist_test.go:136: 2 &{23}
    aslist_test.go:136: 3 {150}
    aslist_test.go:136: 4 {69}
    aslist_test.go:140: 排序后遍历
    aslist_test.go:142: 0 {150}
    aslist_test.go:142: 1 &{121}
    aslist_test.go:142: 2 {120}
    aslist_test.go:142: 3 {69}
    aslist_test.go:142: 4 &{23}
PASS
```

### 7.实现唯一性校验（Set）功能
```golang
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
```
输出结果如下：
```text
=== RUN   TestGanerateUniqueId
--- PASS: TestGanerateUniqueId (0.00s)
    aslist_test.go:79: 测试唯一'性'功能😄，第一次遍历
    aslist_test.go:81: 0 &{121}
    aslist_test.go:81: 1 &{120}
    aslist_test.go:81: 2 &{23}
    aslist_test.go:81: 3 &{150}
    aslist_test.go:81: 4 &{69}
    aslist_test.go:89: 测试唯一'性'功能😄，第二次遍历
    aslist_test.go:91: 0 &{121}
    aslist_test.go:91: 1 &{120}
    aslist_test.go:91: 2 &{23}
    aslist_test.go:91: 3 &{150}
    aslist_test.go:91: 4 &{69}
    aslist_test.go:91: 5 &{123}
    aslist_test.go:91: 6 &{96}
PASS
```
#### 7.1也许你的结构体过于复杂，并且你不擅长写生成唯一id的函数。没有关系，aslist提供了两个通用的函数供你选择
```golang
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
```
输出结果如下：
```text
=== RUN   TestSmartGanerateUniqueId
--- PASS: TestSmartGanerateUniqueId (0.00s)
    aslist_test.go:51: SmartGanerateUniqueId,不忽略指针和结构体类型生成唯一id
    aslist_test.go:53: b1 40cf310cccaa86d9e0acc86a4a5e1fe3
    aslist_test.go:54: b2 e65e33466c73b3c7ef6692759eb0c61b
    aslist_test.go:55: b3 a3628aa02f4eda73de8106404d497476
    aslist_test.go:56: b4 40cf310cccaa86d9e0acc86a4a5e1fe3
    aslist_test.go:58: SmartGanerateUniqueIdWithIgnorePoint,忽略指针和结构体类型生成唯一id
    aslist_test.go:60: b1 40cf310cccaa86d9e0acc86a4a5e1fe3
    aslist_test.go:61: b2 40cf310cccaa86d9e0acc86a4a5e1fe3
    aslist_test.go:62: b3 a3628aa02f4eda73de8106404d497476
    aslist_test.go:63: b4 40cf310cccaa86d9e0acc86a4a5e1fe3
PASS
```

#### 7.2测试aslist.SmartGanerateUniqueId
```golang
	asList := NewAsList()
	//设置GanerateUniqueId函数，asList将会做唯一性校验
	asList.GanerateUniqueId = SmartGanerateUniqueId //aslist.SmartGanerateUniqueId，这里我放的全是指针类型所以不必用SmartGanerateUniqueIdWithIgnorePoint。
	asList.Push(&B{Age: 121})
	asList.Push(&B{Age: 120})
	asList.Push(&B{Age: 23})
	asList.Push(&B{Age: 150})
	asList.Push(&B{Age: 69})
	t.Log("测试唯一性功能，第一次遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
	asList.Push(&B{Age: 123})
	asList.Push(&B{Age: 120}) //重复
	asList.Push(&B{Age: 23})  //重复
	asList.Push(&B{Age: 150}) //重复
	asList.Push(&B{Age: 96})
	t.Log("测试唯一性功能，第二次遍历")
	asList.Range(func(index int, item interface{}) bool {
		t.Log(index, item)
		return false //如果要中断遍历，请返回true
	})
```
输出结果如下：
```text
=== RUN   TestGanerateUniqueIdWithSmartGanerateUniqueId
--- PASS: TestGanerateUniqueIdWithSmartGanerateUniqueId (0.00s)
    aslist_test.go:105: 测试唯一性功能，第一次遍历
    aslist_test.go:107: 0 &{121}
    aslist_test.go:107: 1 &{120}
    aslist_test.go:107: 2 &{23}
    aslist_test.go:107: 3 &{150}
    aslist_test.go:107: 4 &{69}
    aslist_test.go:115: 测试唯一性功能，第二次遍历
    aslist_test.go:117: 0 &{121}
    aslist_test.go:117: 1 &{120}
    aslist_test.go:117: 2 &{23}
    aslist_test.go:117: 3 &{150}
    aslist_test.go:117: 4 &{69}
    aslist_test.go:117: 5 &{123}
    aslist_test.go:117: 6 &{96}
PASS
```

### 8.队列操作
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

### 9.序列化为json
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

### 10.反序列化json
```golang
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
```
输出结果如下：
```text
=== RUN   TestUnMarshalJson
--- PASS: TestUnMarshalJson (0.00s)
    aslist_test.go:163: UnmarshalJson后遍历
    aslist_test.go:165: 0 &{89}
    aslist_test.go:165: 1 &{121}
    aslist_test.go:165: 2 &{120}
    aslist_test.go:165: 3 &{23}
    aslist_test.go:165: 4 &{150}
    aslist_test.go:165: 5 &{69}
PASS
```

### 11.ClearTargets和Clear
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
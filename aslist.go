package aslist

import (
	"encoding/json"
	"errors"
	log "github.com/CreditTone/colorfulog"
	"reflect"
	"sync"
)

type AsList struct {
	locker *sync.Mutex
	list   []interface{}
}

func NewAsList() *AsList {
	return &AsList{
		&sync.Mutex{},
		[]interface{}{},
	}
}

func (self *AsList) MarshalJson() []byte {
	self.locker.Lock()
	defer self.locker.Unlock()
	data, _ := json.Marshal(self.list)
	return data
}

func (self *AsList) UnmarshalJson(data []byte) {
	self.locker.Lock()
	defer self.locker.Unlock()
	self.list = []interface{}{}
	json.Unmarshal(data, &self.list)
}

func (self *AsList) Push(obj interface{}) {
	self.RightPush(obj)
}

func (self *AsList) Pop() interface{} {
	return self.LeftPop()
}

func (self *AsList) RightPush(obj interface{}) {
	self.locker.Lock()
	defer self.locker.Unlock()
	self.list = append(self.list, obj)
}

func (self *AsList) LeftPush(obj interface{}) {
	self.locker.Lock()
	defer self.locker.Unlock()
	newList := []interface{}{obj}
	self.list = append(newList, self.list...)
}

func (self *AsList) RightPop() interface{} {
	self.locker.Lock()
	defer self.locker.Unlock()
	l := len(self.list)
	if l > 0 {
		lastIndex := l - 1
		lastItem := self.list[lastIndex]
		if l == 1 {
			self.list = []interface{}{}
		} else {
			self.list = self.list[:lastIndex]
		}

		return lastItem
	}
	return nil
}

func (self *AsList) LeftPop() interface{} {
	self.locker.Lock()
	defer self.locker.Unlock()
	l := len(self.list)
	if l > 0 {
		firstIndex := 0
		firstItem := self.list[firstIndex]
		if l == 1 {
			self.list = []interface{}{}
		} else {
			self.list = self.list[1:]
		}
		return firstItem
	}
	return nil
}

func (self *AsList) Range(rangeFun func(index int, item interface{})) {
	self.locker.Lock()
	defer self.locker.Unlock()
	for i, item := range self.list {
		rangeFun(i, item)
	}
}

func (self *AsList) checkSortCondition() (Comparator, error) {
	firstType := reflect.TypeOf(self.list[0])
	var comparator Comparator
	var ok bool
	if comparator, ok = self.list[0].(Comparator); !ok {
		return nil, errors.New("类型" + firstType.String() + "必须实现aslist.Comparator接口，才可以使用Sort，或者你也可以使用SortWithCompareFunc动态传入自定义比较函数。")
	}
	for i, item := range self.list {
		if i == 0 {
			continue
		}
		itemType := reflect.TypeOf(item)
		if firstType != itemType {
			return nil, errors.New("Sort方法不支持不同类型 1." + firstType.String() + " 2." + itemType.String() + " 进行排序")
		}
	}
	return comparator, nil
}

func (self *AsList) Sort() {
	if len(self.list) <= 1 {
		return
	}
	if comparator, err := self.checkSortCondition(); err != nil {
		log.Error(err.Error())
	} else {
		self.SortWithCompareFunc(comparator.Compare)
	}
}

func (self *AsList) SortWithCompareFunc(compare func(a, b interface{}) bool) {
	self.locker.Lock()
	defer self.locker.Unlock()
	//冒泡排序
	loopTimes := 0
	happendSwop := false
	newslen := len(self.list)
	for i := 0; i < newslen; i++ {
		happendSwop = false
		for j := newslen - 1; j > i; j-- {
			loopTimes++
			//大分数，置前
			backedItem := self.list[j]  //后项
			frontItem := self.list[j-1] //前项
			if compare(backedItem, frontItem) {
				//交换位置
				happendSwop = true
				self.list[j-1] = backedItem
				self.list[j] = frontItem
			}
		}
		//没有发生交换位置说明已经是顺序了
		if happendSwop == false {
			break
		}
	}
}

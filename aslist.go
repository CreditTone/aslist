package aslist

import (
	"encoding/json"
	"errors"
	log "github.com/CreditTone/colorfulog"
	"gopkg.in/fatih/set.v0"
	"reflect"
	"sync"
)

type AsList struct {
	locker           *sync.Mutex
	list             []interface{}
	GanerateUniqueId func(interface{}) string
	UniqueSet        set.Interface
}

func NewAsList() *AsList {
	return &AsList{
		&sync.Mutex{},
		[]interface{}{},
		nil,
		set.New(set.ThreadSafe),
	}
}

func NewAsListWithGanerateUniqueIdFunc(ganerateUniqueId func(interface{}) string) *AsList {
	return &AsList{
		&sync.Mutex{},
		[]interface{}{},
		ganerateUniqueId,
		set.New(set.ThreadSafe),
	}
}

func (self *AsList) Has(obj interface{}) bool {
	if self.GanerateUniqueId != nil {
		return self.UniqueSet.Has(self.GanerateUniqueId(obj))
	} else {
		ret := false
		self.Range(func(index int, item interface{}) bool {
			if obj == item {
				ret = true
				return true
			}
			return false
		})
		return ret
	}
	return false
}

func (self *AsList) Length() int {
	self.locker.Lock()
	defer self.locker.Unlock()
	return len(self.list)
}

func (self *AsList) MarshalJson() []byte {
	self.locker.Lock()
	defer self.locker.Unlock()
	data, _ := json.Marshal(self.list)
	return data
}

func (self *AsList) UnmarshalJson(data []byte, isAppend bool, unSerialize func(itemData []byte) interface{}) error {
	var err error
	tempList := []interface{}{}
	err = json.Unmarshal(data, &tempList)
	if err != nil {
		return err
	}
	list := []interface{}{}
	for _, item := range tempList {
		itemData, _ := json.Marshal(item)
		val := unSerialize(itemData)
		if val != nil {
			list = append(list, val)
		}
	}
	if isAppend {
		for _, appendItem := range list {
			self.Push(appendItem)
		}
	} else {
		self.locker.Lock()
		defer self.locker.Unlock()
		if self.GanerateUniqueId != nil {
			self.UniqueSet.Clear()
		}
		self.list = []interface{}{}
		for _, item := range list {
			if self.GanerateUniqueId != nil {
				id := self.GanerateUniqueId(item)
				if !self.UniqueSet.Has(id) {
					self.UniqueSet.Add(id)
					self.list = append(self.list, item)
				}
			} else {
				self.list = append(self.list, item)
			}
		}
	}
	return nil
}

func (self *AsList) Push(obj interface{}) {
	self.RightPush(obj)
}

func (self *AsList) Pop() interface{} {
	return self.LeftPop()
}

func (self *AsList) RightPush(obj interface{}) {
	if obj != nil {
		self.locker.Lock()
		defer self.locker.Unlock()
		if self.GanerateUniqueId != nil {
			id := self.GanerateUniqueId(obj)
			if !self.UniqueSet.Has(id) {
				self.UniqueSet.Add(id)
				self.list = append(self.list, obj)
			}
		} else {
			self.list = append(self.list, obj)
		}
	}

}

func (self *AsList) LeftPush(obj interface{}) {
	if obj != nil {
		self.locker.Lock()
		defer self.locker.Unlock()
		newList := []interface{}{obj}
		if self.GanerateUniqueId != nil {
			id := self.GanerateUniqueId(obj)
			if !self.UniqueSet.Has(id) {
				self.UniqueSet.Add(id)
				self.list = append(newList, self.list...)
			}
		} else {
			self.list = append(newList, self.list...)
		}
	}
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
		if self.GanerateUniqueId != nil {
			self.UniqueSet.Remove(self.GanerateUniqueId(lastItem))
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
		if self.GanerateUniqueId != nil {
			self.UniqueSet.Remove(self.GanerateUniqueId(firstItem))
		}
		return firstItem
	}
	return nil
}

func (self *AsList) Range(rangeFunc func(index int, item interface{}) bool) {
	self.locker.Lock()
	defer self.locker.Unlock()
	for i, item := range self.list {
		if rangeFunc(i, item) {
			break
		}
	}
}

func (self *AsList) Clear() {
	self.locker.Lock()
	defer self.locker.Unlock()
	self.list = []interface{}{}
	self.UniqueSet.Clear()
}

func (self *AsList) ClearTargets(clearTargetsFunc func(index int, item interface{}) bool) {
	self.locker.Lock()
	defer self.locker.Unlock()
	nlist := []interface{}{}
	for i, item := range self.list {
		if !clearTargetsFunc(i, item) {
			nlist = append(nlist, item)
		} else {
			if self.GanerateUniqueId != nil {
				self.UniqueSet.Remove(self.GanerateUniqueId(item))
			}
		}
	}
	self.list = nlist
}

func (self *AsList) checkSortCondition() (Comparator, error) {
	firstType := reflect.TypeOf(self.list[0])
	var comparator Comparator
	var ok bool
	if comparator, ok = self.list[0].(Comparator); !ok {
		return nil, errors.New("类型" + firstType.String() + "必须实现aslist.Comparator接口，才可以使用Sort，或者你也可以使用SortWithCompareFunc动态传入自定义比较函数。")
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
	loopTimes := 0
	happendSwop := false
	newslen := len(self.list)
	for i := 0; i < newslen; i++ {
		happendSwop = false
		for j := newslen - 1; j > i; j-- {
			loopTimes++
			backedItem := self.list[j]
			frontItem := self.list[j-1]
			if compare(backedItem, frontItem) {
				happendSwop = true
				self.list[j-1] = backedItem
				self.list[j] = frontItem
			}
		}
		if happendSwop == false {
			break
		}
	}
}

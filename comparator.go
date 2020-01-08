package aslist

type Comparator interface {
	//比较函数，如果希望a排前面则返回true，否则返回false
	Compare(a, b interface{}) bool
}

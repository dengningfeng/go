package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c:=[...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a==c) 长度不同的数组进行比较会编译错误
	t.Log(a == d) //数字长度相同并且元素相同则会返回true

}

func TestBitClear(t *testing.T) {
	a := 7            //0111
	a = a &^ Readable //按位清零
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	a = a &^ Writable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

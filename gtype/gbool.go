package gtype

import "sync/atomic"

type Bool struct {
	val int32
}

func NewBool(value ...bool) *Bool {
	t := &Bool{}
	if len(value) > 0 {
		if value[0] {
			t.val = 1
		} else {
			t.val = 0
		}
	}
	return t
}

func (t *Bool) Clone() *Bool {
	return NewBool(t.Val())
}

// 并发安全设置变量值，返回之前的旧值
func (t *Bool) Set(value bool) (old bool) {
	if value {
		old = atomic.SwapInt32(&t.val, 1) == 1
	} else {
		old = atomic.SwapInt32(&t.val, 0) == 1
	}
	return
}

func (t *Bool) Val() bool {
	return atomic.LoadInt32(&t.val) > 0
}

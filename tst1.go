//
// ternary search tree
//

package tst

import "unicode/utf8"

type TernaryNode1_t struct {
	eq_kid * TernaryNode1_t
	hi_kid * TernaryNode1_t
	lo_kid * TernaryNode1_t
	key rune
	value string	// prefix terminator
}

type TernaryTree1_t struct {
	root * TernaryNode1_t
}

type Cursor1_t struct {
	cur ** TernaryNode1_t
}

func (self * TernaryTree1_t) Add(str string, value string) {
	cur := &self.root
	var last ** TernaryNode1_t
	for _, key := range str {
		for *cur != nil && key != (*cur).key {
			if key < (*cur).key {
				cur = &(*cur).lo_kid
			} else {
				cur = &(*cur).hi_kid
			}
		}
		if *cur == nil {
			*cur = &TernaryNode1_t{key: key}
		}
		last = cur
		cur = &(*cur).eq_kid
	}
	if last != nil {
		(*last).value = value
	}
}

func (self * TernaryTree1_t) Cursor() (c * Cursor1_t) {
	c = &Cursor1_t{}
	c.cur = &self.root
	return
}

func (self * Cursor1_t) Fetch(key rune) (value string, next bool) {
	for *self.cur != nil && key != (*self.cur).key {
		if key < (*self.cur).key {
			self.cur = &(*self.cur).lo_kid
		} else {
			self.cur = &(*self.cur).hi_kid
		}
	}
	if *self.cur == nil {
		return value, false
	}
	if len((*self.cur).value) > 0 {
		value = (*self.cur).value
	}
	self.cur = &(*self.cur).eq_kid
	return value, *self.cur != nil
}

func (self * TernaryTree1_t) Search(str string) (int, int, string, bool) {
	var found string
	var value string
	var next bool
	last := 0
	c := self.Cursor()
	for n, key := range str {
		value, next = c.Fetch(key)
		if len(value) > 0 {
			found = value
			last = n + utf8.RuneLen(key)
		}
		if next == false {
			return last, n, found, false
		}
	}
	return last, len(str), found, next
}

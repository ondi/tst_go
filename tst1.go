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

func (self * TernaryTree1_t) Cursor() (c * Cursor1_t) {
	c = &Cursor1_t{}
	c.cur = &self.root
	return
}

func (self * TernaryTree1_t) Add(str string, value string) {
	c := self.Cursor()
	var last ** TernaryNode1_t
	for _, key := range str {
		for *c.cur != nil && key != (*c.cur).key {
			if key < (*c.cur).key {
				c.cur = &(*c.cur).lo_kid
			} else {
				c.cur = &(*c.cur).hi_kid
			}
		}
		if *c.cur == nil {
			*c.cur = &TernaryNode1_t{key: key}
		}
		last = c.cur
		c.cur = &(*c.cur).eq_kid
	}
	if last != nil {
		(*last).value = value
	}
}

func (self * TernaryTree1_t) Next(c * Cursor1_t, key rune) (value string, next bool) {
	for *c.cur != nil && key != (*c.cur).key {
		if key < (*c.cur).key {
			c.cur = &(*c.cur).lo_kid
		} else {
			c.cur = &(*c.cur).hi_kid
		}
	}
	if *c.cur == nil {
		return value, false
	}
	if len((*c.cur).value) > 0 {
		value = (*c.cur).value
	}
	c.cur = &(*c.cur).eq_kid
	return value, *c.cur != nil
}

func (self * TernaryTree1_t) Search(str string) (int, int, string, bool) {
	last := 0
	c := self.Cursor()
	var found string
	for n, key := range str {
		value, next := self.Next(c, key)
		if len(value) > 0 {
			found = value
			_, size := utf8.DecodeRuneInString(str[n:])
			last = n + size
		}
		if next == false {
			return last, n, found, false
		}
	}
	return last, len(str), found, c.cur != nil
}

func (self * TernaryTree1_t) DEPRECATED_Search1(str string) (int, int, string, bool) {
	last := 0
	c := self.Cursor()
	var value string
	for n, key := range str {
		for *c.cur != nil && key != (*c.cur).key {
			if key < (*c.cur).key {
				c.cur = &(*c.cur).lo_kid
			} else {
				c.cur = &(*c.cur).hi_kid
			}
		}
		if *c.cur == nil {
			return last, n, value, false
		}
		if len((*c.cur).value) > 0 {
			value = (*c.cur).value
			_, size := utf8.DecodeRuneInString(str[n:])
			last = n + size
		}
		c.cur = &(*c.cur).eq_kid
	}
	return last, len(str), value, *c.cur != nil
}

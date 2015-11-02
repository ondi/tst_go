//
// ternary search tree
//

package tst

import "unicode/utf8"

type TernaryNode2_t struct {
	hi_kid int
	eq_kid int
	lo_kid int
	key rune
	value string	// prefix terminator
}

type TernaryTree2_t struct {
	root []TernaryNode2_t
}

type Cursor2_t struct {
	cur int
}

func (self * TernaryTree2_t) Cursor() (c * Cursor2_t) {
	c = &Cursor2_t{}
	if len(self.root) == 0 {
		c.cur = -1
	}
	return
}

func (self * TernaryTree2_t) Add(str string, value string) {
	last := -1
	c := self.Cursor()
	for _, key := range str {
		for c.cur != -1 && key != self.root[c.cur].key {
			last = c.cur
			if key < self.root[c.cur].key {
				c.cur = self.root[c.cur].lo_kid
				if c.cur == -1 {
					self.root[last].lo_kid = len(self.root)
				}
			} else {
				c.cur = self.root[c.cur].hi_kid
				if c.cur == -1 {
					self.root[last].hi_kid = len(self.root)
				}
			}
		}
		if c.cur == -1 {
			c.cur = len(self.root)
			if last != -1 && self.root[last].eq_kid == -1 {
				self.root[last].eq_kid = c.cur
			}
			self.root = append(self.root, TernaryNode2_t{key: key, eq_kid: -1, hi_kid: -1, lo_kid: -1})
		}
		last = c.cur
		c.cur = self.root[c.cur].eq_kid
	}
	if last != -1 {
		self.root[last].value = value
	}
}

func (self * TernaryTree2_t) Next(c * Cursor2_t, key rune) (value string, next bool) {
	for c.cur != -1 && key != self.root[c.cur].key {
		if key < self.root[c.cur].key {
			c.cur = self.root[c.cur].lo_kid
		} else {
			c.cur = self.root[c.cur].hi_kid
		}
	}
	if c.cur == -1 {
		return value, false
	}
	if len(self.root[c.cur].value) > 0 {
		value = self.root[c.cur].value
	}
	c.cur = self.root[c.cur].eq_kid
	return value, c.cur != -1
}

func (self * TernaryTree2_t) Search(str string) (int, int, string, bool) {
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
	return last, len(str), found, c.cur != -1
}

func (self * TernaryTree2_t) DEPRECATED_Search1(str string) (int, int, string, bool) {
	last := 0
	c := self.Cursor()
	var value string
	for n, key := range str {
		for c.cur != -1 && key != self.root[c.cur].key {
			if key < self.root[c.cur].key {
				c.cur = self.root[c.cur].lo_kid
			} else {
				c.cur = self.root[c.cur].hi_kid
			}
		}
		if c.cur == -1 {
			return last, n, value, false
		}
		if len(self.root[c.cur].value) > 0 {
			value = self.root[c.cur].value
			_, size := utf8.DecodeRuneInString(str[n:])
			last = n + size
		}
		c.cur = self.root[c.cur].eq_kid
	}
	return last, len(str), value, c.cur != -1
}

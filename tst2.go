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

func (self * TernaryTree2_t) Add(str string, value string) {
	var key rune
	var cur int
	var last int
	if len(self.root) == 0 {
		cur = -1
		last = -1
	}
	for _, key = range str {
		for cur != -1 && key != self.root[cur].key {
			last = cur
			if key < self.root[cur].key {
				cur = self.root[cur].lo_kid
				if cur == -1 {
					self.root[last].lo_kid = len(self.root)
				}
			} else {
				cur = self.root[cur].hi_kid
				if cur == -1 {
					self.root[last].hi_kid = len(self.root)
				}
			}
		}
		if cur == -1 {
			cur = len(self.root)
			if last != -1 && self.root[last].eq_kid == -1 {
				self.root[last].eq_kid = cur
			}
			self.root = append(self.root, TernaryNode2_t{key: key, eq_kid: -1, hi_kid: -1, lo_kid: -1})
		}
		last = cur
		cur = self.root[cur].eq_kid
	}
	if last != -1 {
		self.root[last].value = value
	}
}

func (self * TernaryTree2_t) Search(str string) (int, int, string, bool) {
	var key rune
	var value string
	var n int
	var last int
	var cur int
	if len(self.root) == 0 {
		cur = -1
	}
	for n, key = range str {
		for cur != -1 && key != self.root[cur].key {
			if key < self.root[cur].key {
				cur = self.root[cur].lo_kid
			} else {
				cur = self.root[cur].hi_kid
			}
		}
		if cur == -1 {
			return last, n, value, false
		}
		if len(self.root[cur].value) > 0 {
			value = self.root[cur].value
			_, size := utf8.DecodeRuneInString(str[n:])
			last = n + size
		}
		cur = self.root[cur].eq_kid
	}
	return last, len(str), value, cur != -1
}

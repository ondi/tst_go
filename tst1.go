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

func (self * TernaryTree1_t) Add(str string, value string) {
	var key rune
	cur := &self.root
	var last ** TernaryNode1_t
	for _, key = range str {
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

func (self * TernaryTree1_t) Search(str string) (int, int, string, bool) {
	var n int
	var last int
	var key rune
	var value string
	cur := self.root
	for n, key = range str {
		for cur != nil && key != cur.key {
			if key < cur.key {
				cur = cur.lo_kid
			} else {
				cur = cur.hi_kid
			}
		}
		if cur == nil {
			return last, n, value, false
		}
		if len(cur.value) > 0 {
			value = cur.value
			_, size := utf8.DecodeRuneInString(str[n:])
			last = n + size
		}
		cur = cur.eq_kid
	}
	return last, len(str), value, cur != nil
}

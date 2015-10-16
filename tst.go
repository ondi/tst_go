//
// ternary search tree
//

package tst

import "unicode/utf8"

type TernaryNode_t struct {
	eq_kid * TernaryNode_t
	hi_kid * TernaryNode_t
	lo_kid * TernaryNode_t
	key rune
	value string
}

type TernaryTree_t struct {
	root * TernaryNode_t
}

func (self * TernaryTree_t) Add(str string, value string) {
	var val rune
	cur := &self.root
	var last ** TernaryNode_t
	for _, val = range str {
		for *cur != nil && val != (*cur).key {
			if val < (*cur).key {
				cur = &(*cur).lo_kid
			} else {
				cur = &(*cur).hi_kid
			}
		}
		if *cur == nil {
			*cur = &TernaryNode_t{}
			(*cur).key = val
		}
		last = cur
		cur = &(*cur).eq_kid
	}
	if last != nil {
		(*last).value = value
	}
}

func (self * TernaryTree_t) Search(str string) (int, bool, string) {
	cur := self.root
	var n int
	var val rune
	var value string
	for n, val = range str {
		if cur == nil {
			return n, false, value
		}
		for val != cur.key {
			if val < cur.key {
				cur = cur.lo_kid
			} else {
				cur = cur.hi_kid
			}
			if cur == nil {
				return n, false, value
			}
		}
		value = cur.value
		cur = cur.eq_kid
	}
	_, size := utf8.DecodeRuneInString(str[n:])
	return n + size, cur != nil, value
}

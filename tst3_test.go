//
//
//

package tst

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Tst3_01(t *testing.T) {
	temp := NewTree3[string]()

	temp.Add("/", "/")
	temp.Add("/debug/kube", "/debug/kube")
	temp.Add("/debug", "/debug")
	temp.Add("/pprof", "/pprof")
	temp.Add("/metrics", "/metrics")

	var ok bool
	var value string

	value, ok = temp.Search("")
	assert.Assert(t, ok == false)

	value, ok = temp.Search("v1/test")
	assert.Assert(t, ok == false)

	value, ok = temp.Search("/v1/test")
	assert.Assert(t, value == "/")

	value, ok = temp.Search("/debu")
	assert.Assert(t, value == "/", value)

	value, ok = temp.Search("/debug/test")
	assert.Assert(t, value == "/debug")

	value, ok = temp.Search("/debug/size")
	assert.Assert(t, value == "/debug")

	value, ok = temp.Search("/pprof/heap")
	assert.Assert(t, value == "/pprof")

	value, ok = temp.Search("/pprof/profile")
	assert.Assert(t, value == "/pprof")

	value, ok = temp.Search("/metrics/sql")
	assert.Assert(t, value == "/metrics")

	value, ok = temp.Search("/metrics/page")
	assert.Assert(t, value == "/metrics")
}

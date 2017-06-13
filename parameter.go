package gaap

import (
	"strconv"
)

// Parameters save url parameters.
type Parameters map[string]string

// Get gets url parameter
//  ex. routing : /foo/{bar}
//      url     : http://{host}/foo/1
//
//  bar := ctx.Parameters.Get("bar")  // bar = "1"
func (p Parameters) Get(key string) string {
	return p[key]
}

// GetAsInt gets url parameter as int value
func (p Parameters) GetAsInt(key string) (int64, error) {
	val := p.Get(key)
	return strconv.ParseInt(val, 0, 64)
}

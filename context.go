package gaap

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// V is utility type to create json response.
//  ex. gaap.JSON(200, gaap.V{"id": 1})
type V map[string]interface{}

// Context is wrapper request, response and appengine context
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Ctx            context.Context
	Params         Parameters
}

// NewContext creates new context struct
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	ctx := appengine.NewContext(r)
	return &Context{
		ResponseWriter: w,
		Request:        r,
		Ctx:            ctx,
		Params:         mux.Vars(r),
	}
}

// ParseJSONBody parses request json to given argument
//  usage.
//
//   struct Foo {
//     Foo string `json:"foo"`
//     Bar string `json:"bar"`
//   }
//   ...
//   payload := new(Foo)
//   err := ctx.ParseJSONBody(payload)
//   // Handlings
func (c *Context) ParseJSONBody(v interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}

// JSON is util method to response as json
func (c *Context) JSON(status int, v interface{}) {
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.ResponseWriter.WriteHeader(status)
	if err := json.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
		log.Errorf(c.Ctx, "[gaap.Context#JSON()] %v", err)
		c.InternalServerError(err.Error())
	}
}

// InternalServerError is util method to response server error
func (c *Context) InternalServerError(msg string) {
	http.Error(c.ResponseWriter, msg, http.StatusInternalServerError)
}

// BadRequest is util method to response badrequest
func (c *Context) BadRequest(msg string) {
	http.Error(c.ResponseWriter, msg, http.StatusBadRequest)
}

// NotFound is util method to respose notfound
func (c *Context) NotFound(msg string) {
	http.Error(c.ResponseWriter, msg, http.StatusNotFound)
}

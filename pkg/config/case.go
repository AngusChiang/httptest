package config

import (
	"strings"

	"github.com/wklken/httptest/pkg/tpl"
)

// TODO: use reflect instead hard code

type BasicAuth struct {
	Username string
	Password string
}

func (b *BasicAuth) Empty() bool {
	return b.Username == "" && b.Password == ""
}

type Request struct {
	Method string
	URL    string
	Body   string

	Header    map[string]string
	Cookie    string
	BasicAuth BasicAuth `mapstructure:"basic_auth"`
}

const TplBrace = "{{"

func (r *Request) Render(ctx map[string]interface{}) {
	if strings.Contains(r.Method, TplBrace) {
		r.Method = tpl.Render(r.Method, ctx)
	}

	if strings.Contains(r.URL, TplBrace) {
		r.URL = tpl.Render(r.URL, ctx)
	}

	if strings.Contains(r.Body, TplBrace) {
		r.Body = tpl.Render(r.Body, ctx)
	}

	for k, v := range r.Header {
		if strings.Contains(v, TplBrace) {
			r.Header[k] = tpl.Render(v, ctx)
		}
	}
}

type Assert struct {
	Status string

	StatusCode      int
	StatusCodeIn    []int `mapstructure:"statusCode_in"`
	StatusCodeNotIn []int `mapstructure:"statusCode_not_in"`
	StatusCodeLt    int   `mapstructure:"statusCode_lt"`
	StatusCodeLte   int   `mapstructure:"statusCode_lte"`
	StatusCodeGt    int   `mapstructure:"statusCode_gt"`
	StatusCodeGte   int   `mapstructure:"statusCode_gte"`

	ContentLength    int64
	ContentLengthLt  int64 `mapstructure:"contentLength_lt"`
	ContentLengthLte int64 `mapstructure:"contentLength_lte"`
	ContentLengthGt  int64 `mapstructure:"contentLength_gt"`
	ContentLengthGte int64 `mapstructure:"contentLength_gte"`

	ContentType string

	// TODO
	StatusIn         []string `mapstructure:"status_in"`
	StatusNotIn      []string `mapstructure:"status_not_in"`
	ContentTypeIn    []string `mapstructure:"contentType_in"`
	ContentTypeNotIn []string `mapstructure:"contentType_not_in"`

	// latency
	LatencyLt  int64 `mapstructure:"latency_lt"`
	LatencyLte int64 `mapstructure:"latency_lte"`
	LatencyGt  int64 `mapstructure:"latency_gt"`
	LatencyGte int64 `mapstructure:"latency_gte"`

	Body string

	BodyContains      string `mapstructure:"body_contains"`
	BodyNotContains   string `mapstructure:"body_not_contains"`
	BodyStartsWith    string `mapstructure:"body_startswith"`
	BodyEndsWith      string `mapstructure:"body_endswith"`
	BodyNotStartsWith string `mapstructure:"body_not_startswith"`
	BodyNotEndsWith   string `mapstructure:"body_not_endswith"`

	Json   []AssertJson `mapstructure:"json"`
	Header map[string]interface{}
}

func (a *Assert) Render(ctx map[string]interface{}) {
	a.Status = tpl.Render(a.Status, ctx)

	a.ContentType = tpl.Render(a.ContentType, ctx)

	if len(a.StatusIn) > 0 {
		n := make([]string, 0, len(a.StatusIn))
		for _, s := range a.StatusIn {
			n = append(n, tpl.Render(s, ctx))
		}
		a.StatusIn = n
	}
	if len(a.StatusNotIn) > 0 {
		n := make([]string, 0, len(a.StatusNotIn))
		for _, s := range a.StatusNotIn {
			n = append(n, tpl.Render(s, ctx))
		}
		a.StatusNotIn = n
	}
	if len(a.ContentTypeIn) > 0 {
		n := make([]string, 0, len(a.ContentTypeIn))
		for _, s := range a.ContentTypeIn {
			n = append(n, tpl.Render(s, ctx))
		}
		a.ContentTypeIn = n
	}
	if len(a.ContentTypeNotIn) > 0 {
		n := make([]string, 0, len(a.ContentTypeNotIn))
		for _, s := range a.ContentTypeNotIn {
			n = append(n, tpl.Render(s, ctx))
		}
		a.ContentTypeNotIn = n
	}

	a.Body = tpl.Render(a.Body, ctx)

	a.BodyContains = tpl.Render(a.BodyContains, ctx)
	a.BodyNotContains = tpl.Render(a.BodyNotContains, ctx)
	a.BodyStartsWith = tpl.Render(a.BodyStartsWith, ctx)
	a.BodyEndsWith = tpl.Render(a.BodyEndsWith, ctx)
	a.BodyNotStartsWith = tpl.Render(a.BodyNotStartsWith, ctx)
	a.BodyNotEndsWith = tpl.Render(a.BodyNotEndsWith, ctx)

	for _, j := range a.Json {
		j.Render(ctx)
	}

}

type AssertJson struct {
	Path  string
	Value interface{}
}

func (a *AssertJson) Render(ctx map[string]interface{}) {
	if strings.Contains(a.Path, TplBrace) {
		a.Path = tpl.Render(a.Path, ctx)
	}
}

type Hook struct {
	SaveCookie string `mapstructure:"save_cookie"`
}

type Case struct {
	Title       string
	Description string

	Request Request
	Assert  Assert
	Hook    Hook
}

func (c *Case) Render(ctx map[string]interface{}) {
	if strings.Contains(c.Title, TplBrace) {
		c.Title = tpl.Render(c.Title, ctx)
	}
	if strings.Contains(c.Description, TplBrace) {
		c.Description = tpl.Render(c.Description, ctx)

	}

	c.Request.Render(ctx)
	c.Assert.Render(ctx)
}

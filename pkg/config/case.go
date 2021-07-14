package config

type Request struct {
	Method string
	URL    string
	Body   string

	Header map[string]string
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

	Json []AssertJson `mapstructure:"json"`
}

type Case struct {
	Title       string
	Description string

	Request Request
	Assert  Assert
}

type AssertJson struct {
	Path  string
	Value interface{}
}

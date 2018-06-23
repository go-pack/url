package url

import (
	"strings"
	url2 "net/url"
)

type Builder struct {
	scheme   string
	port     string
	host     string
	path     string
	fragment string
	query    url2.Values
}

func (b *Builder) Scheme(scheme string) *Builder {
	b.scheme = scheme
	return b
}
func (b *Builder) Port(port string) *Builder {
	b.port = port
	return b
}
func (b *Builder) Path(path string) *Builder {
	b.path = path
	return b
}
func (b *Builder) Query(query map[string][]string) *Builder {
	b.query = query
	return b
}
func (b *Builder) AddQuery(query map[string][]string) *Builder {
	for k, v := range query {
		b.query[k] = v
	}
	return b
}
func (b *Builder) PathVariable(pathVariable map[string]string) *Builder {

	for k, v := range pathVariable {
		b.path = strings.Replace(b.path, k, v, 1)
	}
	return b

}
func (b *Builder) Fragment(fragment string) *Builder {
	b.fragment = fragment
	return b
}

//Init
func (b *Builder) Init(url string) bool {
	urlInfo, err := url2.ParseRequestURI(url)
	if err != nil {
		return false
	} else {
		b.scheme = urlInfo.Scheme
		b.port = urlInfo.Port()
		b.path = urlInfo.Path
		b.query = urlInfo.Query()
	}
	return true
}
func (b *Builder) ToString() string {
	var url []string
	i := len(b.scheme) <= 0
	if i {
		b.scheme = "http://"
	} else {
		b.scheme += "//"
	}
	url = append(url, b.scheme)

	h := len(b.host) <= 0
	if h {
		b.host = "127.0.0.1"
	}
	url = append(url, b.host)

	p := len(b.path) > 0
	if p {
		url = append(url, b.path)
	}

	pt := len(b.port) > 0
	if pt {
		url = append(url, b.port)
	}

	var queryString string
	if len(b.query) > 0 {
		queryString += "?" + b.query.Encode()
	}

	url = append(url, queryString)

	pf := len(b.fragment) > 0
	if pf {
		url = append(url, "#"+b.fragment)
	}
	return strings.Join(url, "")
}

func NewBuilder() *Builder {
	return &Builder{}
}

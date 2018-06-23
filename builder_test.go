package url

import (
	"testing"
	"fmt"
)

func TestBuilder_ToString(t *testing.T) {
	urlBuilder := NewBuilder()
	mx := make(map[string][]string)
	mx["uuid"] = []string{"123456","777"}
	path := urlBuilder.Path("/vhot/api").Query(mx).Fragment("xx").ToString()
	fmt.Printf("%s",path)
}
func TestBuilder_Init(t *testing.T) {
	urlBuilder := NewBuilder()
	urlBuilder.Init("http://127.0.0.1/vhot/api?uuid=123456&uuid=777")

	mx := make(map[string][]string)
	mx["ccid"] = []string{"123456","777"}
	path := urlBuilder.AddQuery(mx).ToString()
	fmt.Printf("%s",path)

}

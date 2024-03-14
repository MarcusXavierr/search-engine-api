package search_test

import (
	"io"
	"strings"
	"testing"

	"github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch/search"
)

const (
	blogSearchJson = `{"size":0,"query":{"multi_match":{"query":"cringe","fields":["title^3","content"]}},"aggs":{"unique_titles":{"terms":{"field":"title.keyword","size":10},"aggs":null},"aggs":{"top_hit_for_title":{"top_hits":{"size":1}}}}}`
)

func TestMountBlogSearchQuery(t *testing.T) {
	got := new(strings.Builder)
	result := search.MountBlogSearchQuery("cringe", 10, true)
	_, err := io.Copy(got, result.GetQuery())

	if err != nil {
		t.Error(err)
	}

	want := blogSearchJson

	if got.String() != want {
		t.Errorf("expected %s\n but got %s", want, got)
	}

}

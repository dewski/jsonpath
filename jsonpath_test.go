package jsonpath

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplePathExamples(t *testing.T) {
	jp := NewReader(map[string]interface{}{
		"action": "user.login",
		"issue": map[string]interface{}{
			"title": "My issue",
		},
		"users": []map[string]interface{}{
			{
				"username": "dewski",
			},
		},
	})

	assert.Equal(t, "user.login", jp.Path(".action"))
	assert.Equal(t, "My issue", jp.Path(".issue.title"))
	assert.Equal(t, "dewski", jp.Path(".users[0].username"))
}

func TestSimplePaths(t *testing.T) {
	jp := NewReader(map[string]interface{}{
		"action": "user.login",
		"issue": map[string]interface{}{
			"title": "My issue",
		},
		"users": []map[string]interface{}{
			{
				"username": "dewski",
			},
		},
	})

	expected := []string{
		".action",
		".issue.title",
		".users[0].username",
	}
	actual := jp.Paths()
	sort.Strings(expected)
	sort.Strings(actual)

	assert.Equal(t, expected, actual)
}

func TestPathWithSpaces(t *testing.T) {
	jp := NewReader(map[string]interface{}{
		"event type": "user.login",
		"some issues": map[string]interface{}{
			"issue title": "My issue",
		},
		"all my users": []map[string]interface{}{
			{
				"user name": "dewski",
			},
		},
	})

	assert.Equal(t, "user.login", jp.Path(`."event type"`))
	assert.Equal(t, "My issue", jp.Path(`."some issues"."issue title"`))
	assert.Equal(t, "dewski", jp.Path(`."all my users"[0]."user name"`))
}

type custom string

func TestPathPreservesOriginalType(t *testing.T) {
	jp := NewReader(map[string]interface{}{
		"action": custom("user.login"),
	})

	assert.Equal(t, "jsonpath.custom", reflect.TypeOf(jp.Path(`.action`)).String())
}

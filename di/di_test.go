package di

import (
	"fmt"
	"github.com/fatih/structtag"
	"reflect"
	"sort"
	"testing"
)

func TestTagParse(t *testing.T) {
	s := tag{}
	tp := reflect.TypeOf(s)
	field, _ := tp.FieldByName("a")
	fmt.Println(field.Tag.Get("json"))
	tags, _ := structtag.Parse(string(field.Tag))
	// iterate over all tags
	for _, t := range tags.Tags() {
		fmt.Printf("tag: %+v\n", t)
	}
	jsonTag, _ := tags.Get("json")
	jsonTag.Name = "foo_bar"
	jsonTag.Options = nil
	tags.Set(jsonTag)
	tags.Set(&structtag.Tag{
		Key:     "hcl",
		Name:    "foo",
		Options: []string{"squash"},
	})
	fmt.Println(tags) // Output: json:"foo_bar" xml:"foo" hcl:"foo,squash"
	// sort tags according to keys
	sort.Sort(tags)
	fmt.Println(tags) // Output: hcl:"foo,squash" json:"foo_bar" xml:"foo"
}

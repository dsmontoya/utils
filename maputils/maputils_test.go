package maputils

import (
	"reflect"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInterfaceValue(t *testing.T) {
	Convey("Given a map[string]string", t, func() {
		m := map[string]string{
			"a": "a",
			"b": "b",
		}
		StringToInterface(m)
	})
}

func TestCopy(t *testing.T) {
	type post struct {
		Title string
		Date  time.Time
	}
	type user struct {
		Name  string
		Posts []*post
	}
	source := map[string]interface{}{
		"string": "text",
		"int":    1,
		"array":  []int{1, 2, 3, 4},
		"map": map[string]interface{}{
			"bool":   true,
			"struct": struct{}{},
		},
		"user": &user{
			Name: "john",
			Posts: []*post{{
				Title: "A good title",
				Date:  time.Now(),
			}},
		},
	}
	destination := map[string]interface{}{}
	t.Run("copy", func(t *testing.T) {
		Copy(source, destination)

		if !reflect.DeepEqual(source, destination) {
			t.Error("source and destination should be equal")
		}

		destination["int"] = 2
		destination["map"].(map[string]interface{})["bool"] = false

		if source["int"] == 2 {
			t.Error("source[\"int\"] = 2, want 1")
		}
		if !source["map"].(map[string]interface{})["bool"].(bool) {
			t.Error("source[\"map\"][\"bool\"] = false, want true")
		}
	})
}

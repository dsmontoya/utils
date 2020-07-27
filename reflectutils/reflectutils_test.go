package reflectutils

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testStruct struct{}

func TestDeepValue(t *testing.T) {
	Convey("Given a pointer to struct value", t, func() {
		s := &testStruct{}
		v := reflect.ValueOf(s)
		Convey("When the deep value is checked", func() {
			dv := DeepValue(v)

			Convey("The deep value should be a struct", func() {
				k := dv.Kind()
				So(k, ShouldEqual, reflect.Struct)
			})
		})
	})

	Convey("Given an interface containing a pointer struct value", t, func() {
		var iface interface{}
		s := &testStruct{}
		iface = s
		v := reflect.ValueOf(iface)
		Convey("When the deep value is checked", func() {
			dv := DeepValue(v)

			Convey("The deep value should be a struct", func() {
				k := dv.Kind()
				So(k, ShouldEqual, reflect.Struct)
			})
		})
	})

	Convey("Given a nil value", t, func() {
		var s *testStruct
		v := reflect.ValueOf(s)
		Convey("When the deep value is checked", func() {
			dv := DeepValue(v)

			Convey("The deep value should be a string", func() {
				k := dv.Kind()
				So(k, ShouldEqual, reflect.Invalid)
			})
		})
	})

	Convey("Given a string value", t, func() {
		s := ""
		v := reflect.ValueOf(s)
		Convey("When the deep value is checked", func() {
			dv := DeepValue(v)

			Convey("The deep value should be a string", func() {
				k := dv.Kind()
				So(k, ShouldEqual, reflect.String)
			})
		})
	})
}

func TestSetSlice(t *testing.T) {
	type args struct {
		slice interface{}
		index int
		x     interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"[]int without pointer", args{[]int{1, 2, 3}, 0, 5}, []int{5, 2, 3}},
		{"[]int with pointer", args{&[]int{1, 2, 3}, 0, 5}, &[]int{5, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetSlice(tt.args.slice, tt.args.index, tt.args.x)
			if !reflect.DeepEqual(tt.args.slice, tt.want) {
				t.Errorf("slice = %v, want %v", tt.args.slice, tt.want)
			}
		})
	}
}

func TestEach(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name  string
		args  args
		limit int
	}{
		{"break after slice[1]", args{[]int{1, 2, 3, 4}}, 1},
	}
	for _, tt := range tests {
		count := 0
		f := func(i int, item reflect.Value) bool {
			if i == tt.limit {
				return false
			}
			count++
			return true
		}
		t.Run(tt.name, func(t *testing.T) {
			Each(tt.args.slice, f)
			if count != tt.limit {
				t.Errorf("count = %v, want %v", count, tt.limit)
			}
		})
	}
}

func TestSetField(t *testing.T) {
	numberPointer := 5
	type strct struct {
		Number    int
		NumberPtr *int
	}
	type args struct {
		container interface{}
		name      string
		value     interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"set Number", args{&strct{Number: 1}, "Number", 5}, &strct{Number: 5}},
		{"set NumberPtr", args{&strct{NumberPtr: new(int)}, "NumberPtr", 5}, &strct{NumberPtr: &numberPointer}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetField(tt.args.container, tt.args.name, tt.args.value)
			if !reflect.DeepEqual(tt.args.container, tt.want) {
				t.Errorf("container = %v, want %v", tt.args.container, tt.want)
			}
		})
	}
}

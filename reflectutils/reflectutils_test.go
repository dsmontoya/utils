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
		{"break after slice[len(slice)-1]", args{[]int{1, 2, 3, 4}}, 3},
	}
	for _, tt := range tests {
		f := func(i int, item reflect.Value) bool {
			if i+1 == tt.limit {
				return false
			}
			return true
		}
		t.Run(tt.name, func(t *testing.T) {
			count := Each(tt.args.slice, f)
			if count != tt.limit {
				t.Errorf("count = %v, want %v", count, tt.limit)
			}
		})
	}
}

func TestSetField(t *testing.T) {
	numberPointer := 5
	type strct struct {
		Number     int
		NumberPtr  *int
		unexported int
	}
	type args struct {
		container interface{}
		name      string
		value     interface{}
	}
	tests := []struct {
		name      string
		args      args
		want      bool
		wantValue interface{}
	}{
		{"Number", args{&strct{Number: 1}, "Number", 5}, true, &strct{Number: 5}},
		{"NumberPtr", args{&strct{NumberPtr: new(int)}, "NumberPtr", 5}, true, &strct{NumberPtr: &numberPointer}},
		{"no ptr", args{strct{NumberPtr: new(int)}, "NumberPtr", 5}, false, strct{NumberPtr: new(int)}},
		{"Unknown", args{&strct{Number: 1}, "Unkown", 5}, false, &strct{Number: 1}},
		{"Unexported", args{&strct{Number: 1}, "unexported", 5}, false, &strct{Number: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetField(tt.args.container, tt.args.name, tt.args.value); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.container, tt.wantValue) {
				t.Errorf("container = %v, want %v", tt.args.container, tt.wantValue)
			}
		})
	}
}

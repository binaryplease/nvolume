package nvolume

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		size []int
	}
	tests := []struct {
		name string
		args args
		want NVolume
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.size...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNVolume_Set(t *testing.T) {
	type fields struct {
		Dimensions int
		Shape      []int
	}
	type args struct {
		val NElement
		dim []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nvol := &NVolume{
				Dimensions: tt.fields.Dimensions,
				Shape:      tt.fields.Shape,
			}
			nvol.Set(tt.args.val, tt.args.dim...)
		})
	}
}

func TestNVolume_Get(t *testing.T) {
	type fields struct {
		Dimensions int
		Shape      []int
	}
	type args struct {
		dim []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   NElement
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nvol := NVolume{
				Dimensions: tt.fields.Dimensions,
				Shape:      tt.fields.Shape,
			}
			if got := nvol.Get(tt.args.dim...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NVolume.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dimToPos(t *testing.T) {
	tests := []struct {
		name string
		vol  NVolume
		dims []int
		want int
	}{
		{"2 2 2", New(3, 3, 3), []int{2, 2, 2}, 26},
		{"0 0 0", New(3, 3, 3), []int{0, 0, 0}, 0},
		{"0 0 1", New(3, 3, 3), []int{0, 0, 1}, 1},
		{"0 0 2", New(3, 3, 3), []int{0, 0, 2}, 2},
		{"0 0 3", New(3, 3, 3), []int{0, 0, 3}, 3},
		{"0 0 4", New(3, 3, 3), []int{0, 0, 4}, 4},
		{"0 0 0", New(3, 3, 3), []int{0, 0, 0}, 4},
		{"1 0 0", New(3, 3, 3), []int{1, 0, 0}, 4},
		{"2 0 0", New(3, 3, 3), []int{2, 0, 0}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.vol.dimToPos(tt.dims); got != tt.want {
				t.Errorf("dimToPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

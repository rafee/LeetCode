/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_radomizedConstructor(t *testing.T) {
	tests := []struct {
		name string
		want RandomizedSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := radomizedConstructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("radomizedConstructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedSet_Insert(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *RandomizedSet
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Insert(tt.args.val); got != tt.want {
				t.Errorf("RandomizedSet.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedSet_Remove(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *RandomizedSet
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Remove(tt.args.val); got != tt.want {
				t.Errorf("RandomizedSet.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedSet_GetRandom(t *testing.T) {
	tests := []struct {
		name string
		this *RandomizedSet
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.GetRandom(); got != tt.want {
				t.Errorf("RandomizedSet.GetRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

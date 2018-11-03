package Solution

import (
	"reflect"
	"testing"
)

func TestMajorityElement1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCacse 1", []int{3, 2, 3}, 3},
		{"TestCacse 1", []int{4, 4, 4, 3, 3}, 4},
		{"TestCacse 2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := majorityElement1(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestMajorityElement2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCacse 1", []int{3, 2, 3}, 3},
		{"TestCacse 2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := majorityElement2(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestMajorityElement3(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCacse 1", []int{3, 2, 3}, 3},
		{"TestCacse 2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := majorityElement3(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestMajorityElement4(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCacse 1", []int{3, 2, 3}, 3},
		{"TestCacse 2", []int{3, 3, 4}, 3},
		{"TestCacse 3", []int{6, 6, 6, 7, 7}, 6},
		{"TestCacse 4", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"TestCacse 5", []int{233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333}, 233},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := majorityElement4(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestMajorityElement5(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCacse 1", []int{3, 2, 3}, 3},
		{"TestCacse 2", []int{3, 3, 4}, 3},
		{"TestCacse 3", []int{6, 6, 6, 7, 7}, 6},
		{"TestCacse 4", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"TestCacse 5", []int{233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333}, 233},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := majorityElement5(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution4(b *testing.B) {
	b.StopTimer()

	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCacse 1", []int{3, 2, 3}, 3},
		{"TestCacse 2", []int{3, 3, 4}, 3},
		{"TestCacse 3", []int{6, 6, 6, 7, 7}, 6},
		{"TestCacse 4", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"TestCacse 5", []int{233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333, 2333}, 233},
	}
	b.StartTimer()

	b.Run("majorityElement1", func(b *testing.B) {
		for _, c := range cases {
			got := majorityElement1(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				b.Errorf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		}
	})
	b.Run("majorityElement2", func(b *testing.B) {
		for _, c := range cases {
			got := majorityElement2(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				b.Errorf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		}
	})
	b.Run("majorityElement3", func(b *testing.B) {
		for _, c := range cases {
			got := majorityElement3(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				b.Errorf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		}
	})
	b.Run("majorityElement4", func(b *testing.B) {
		for _, c := range cases {
			got := majorityElement4(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				b.Errorf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		}
	})
	b.Run("majorityElement5", func(b *testing.B) {
		for _, c := range cases {
			got := majorityElement5(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				b.Errorf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		}
	})
}

//	使用案列
func ExampleSolution() {

}

package main

import "testing"

var tests = []struct {
	FuncNames []string
	Vals      []int
	expected  []interface{}
	result    []interface{}
}{
	{
		[]string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
		[]int{0, -2, 0, -3, 0, 0, 0, 0},
		[]interface{}{nil, nil, nil, nil, -3, nil, 0, -2},
		[]interface{}{},
	},
	{
		[]string{"MinStack", "push", "push", "top", "getMin", "pop", "getMin", "top"},
		[]int{0, 1, 2, 0, 0, 0, 0, 0},
		[]interface{}{nil, nil, nil, 2, 1, nil, 1, 1},
		[]interface{}{},
	},
	{
		[]string{"MinStack", "push", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin", "pop", "getMin"},
		[]int{0, 2, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0},
		[]interface{}{nil, nil, nil, nil, nil, 0, nil, 0, nil, 0, nil, 2},
		[]interface{}{},
	},
}

func TestMinstack(t *testing.T) {

	var minStack MinStack

	for _, v := range tests {

		for x, y := range v.FuncNames {
			switch y {
			case "MinStack":
				minStack = Constructor()
				v.result = append(v.result, nil)
			case "push":
				minStack.Push(v.Vals[x])
				v.result = append(v.result, nil)
			case "pop":
				minStack.Pop()
				v.result = append(v.result, nil)
			case "getMin":
				res := minStack.GetMin()
				v.result = append(v.result, res)
			case "top":
				top := minStack.Top()
				v.result = append(v.result, top)

			}
		}

		for key, val := range v.expected {
			if v.result[key] != val {
				t.Errorf("expected %v got %v", val, v.result[key])
			}
		}

	}

}

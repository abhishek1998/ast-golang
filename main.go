package main

import (
	"fmt"
	"math"
)

type node struct {
	value  *int
	symbol *string
	left   *node
	right  *node
}

func (n *node) evaluate() float64 {
	res := 0.0

	if n == nil {
		return res
	}

	if n.value != nil {
		res = float64(*n.value)
	} else {
		if *n.symbol == "+" {
			//get the result
			res += n.left.evaluate()
			res += n.right.evaluate()
			return res
		} else if *n.symbol == "-" {
			//get the result
			left := n.left.evaluate()
			right := n.right.evaluate()
			res = math.Abs(float64(left - right))
			return res
		}

	}

	return res
}

func print(n *node) {
	if n == nil {
		return
	}
	print(n.left)
	if n.value != nil {
		fmt.Printf("%d ", *n.value)
	} else {
		fmt.Printf("%s ", *n.symbol)
	}
	print(n.right)

}

func main() {
	var plus_sign *string
	plus_sign = new(string)
	*plus_sign = "+"

	var minus_sign *string
	minus_sign = new(string)
	*minus_sign = "-"

	var num_1 *int
	num_1 = new(int)
	*num_1 = 10

	var num_2 *int
	num_2 = new(int)
	*num_2 = 30

	var num_3 *int
	num_3 = new(int)
	*num_3 = 10

	root := &node{
		value:  nil,
		symbol: plus_sign,
		left: &node{
			value:  nil,
			symbol: plus_sign,
			left: &node{
				value:  num_1,
				symbol: nil,
				left:   nil,
				right:  nil,
			},
			right: &node{
				value:  num_2,
				symbol: nil,
				left:   nil,
				right:  nil,
			},
		},
		right: &node{
			value:  nil,
			symbol: minus_sign,
			left: &node{
				value:  num_1,
				symbol: nil,
				left:   nil,
				right:  nil,
			},
			right: &node{
				value:  num_2,
				symbol: nil,
				left:   nil,
				right:  nil,
			},
		},
	}

	print(root)
	fmt.Println()
	fmt.Println(root.evaluate())
}

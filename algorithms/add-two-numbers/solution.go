package add_two_numbers

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val1 := big.NewInt(0)
	val2 := big.NewInt(0)
	val1, _ = val1.SetString(reverseString(readNode(l1)), 10)
	val2, _ = val2.SetString(reverseString(readNode(l2)), 10)
	result := big.NewInt(0)
	result.Add(val1, val2)
	resultStr := reverseString(result.String())

	return makeNode(resultStr)
}

func readNode(l *ListNode) string {
	val := strconv.Itoa(l.Val)
	if l.Next == nil {
		return val
	}
	valStr := val + readNode(l.Next)
	return valStr
}

func makeNode(str string) *ListNode {
	var next *ListNode
	if len(str) > 0 {
		var err error
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		next = makeNode(str[1:])
	}
	if len(str) > 0 {
		val, _ := strconv.Atoi(str[0:1])
		return &ListNode{
			Val:  val,
			Next: next,
		}
	} else {
		return nil
	}
}

func reverseString(str string) string {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += str[i : i+1]
	}
	return reversed
}

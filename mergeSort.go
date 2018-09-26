package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	stepOne := head
	stepTwo := stepOne
	for stepTwo.Next != nil && stepTwo.Next.Next != nil {
		stepOne = stepOne.Next
		stepTwo = stepTwo.Next.Next
	}
	secondHalf := stepOne.Next
	stepOne.Next = nil
	head = SortList(head)
	secondHalf = SortList(secondHalf)
	return merge(head, secondHalf)
}

func merge(first, second *ListNode) (newHead *ListNode) {
	newHead = new(ListNode)
	temp := newHead
	for first != nil && second != nil {
		if first.Val < second.Val {
			temp.Next = first
			first = first.Next
		} else {
			temp.Next = second
			second = second.Next
		}
		temp = temp.Next
	}
	if first != nil {
		temp.Next = first
	}
	if second != nil {
		temp.Next = second
	}
	return newHead.Next
}

func getSlice(head *ListNode) []int {
	sli := make([]int, 0)
	for head != nil {
		sli = append(sli, head.Val)
		head = head.Next
	}
	return sli
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func makeList(arr []int) *ListNode {
	head := new(ListNode)
	temp := head
	var prev *ListNode
	for _, ele := range arr {
		temp.Val = ele
		temp.Next = new(ListNode)
		prev = temp
		temp = temp.Next
	}
	prev.Next = nil
	return head
}

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	strs := strings.Split(input, " ")
	nums := make([]int, 0)
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		nums = append(nums, num)
	}
	list := makeList(nums)
	printList(SortList(list))
}

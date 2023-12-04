package main

import "fmt"

/*
LINK: https://practice.geeksforgeeks.org/problems/subarray-with-given-sum-1587115621/1


Description:
Given an unsorted array A of size N that contains only positive integers, find a continuous sub-array that adds to a given number S and return the left and right index(1-based indexing) of that subarray.

In case of multiple subarrays, return the subarray indexes which come first on moving from left to right.

Note:- You have to return an ArrayList consisting of two elements left and right. In case no such subarray exists return an array consisting of element -1.

Example 1:

Input:
N = 5, S = 12
A[] = {1,2,3,7,5}
Output: 2 4
Explanation: The sum of elements
from 2nd position to 4th position
is 12.

Example 2:

Input:
N = 10, S = 15
A[] = {1,2,3,4,5,6,7,8,9,10}
Output: 1 5
Explanation: The sum of elements
from 1st position to 5th position
is 15.
*/

func main() {

	//Input:
	//	N = 5, S = 12
	//	A[] = {1,2,3,7,5}
	//Output: 2 4

	input := []int{1, 2, 3, 7, 5}
	s := 12

	firstExample(input, s)
	//complexity o(n²)
	slidingWindow(input, s)
	//complexity o(n)

}

func firstExample(input []int, s int) {
	var sum int
	for i := 0; i < len(input); i++ { // N
		sum = 0
		for f := i; f < len(input); f++ { // N
			sum += input[f]
			if sum == s {
				fmt.Printf("%d %d\n", i+1, f+1)
				return
			}

		}

	}
	fmt.Printf("-1")
}

func slidingWindow(input []int, s int) {
	var sum int
	inicio := 0

	for finish := 0; finish < len(input); finish++ { //N
		sum += input[finish]

		for sum > s && inicio < finish {
			sum -= input[inicio]
			inicio++
		}

		if sum == s {
			fmt.Printf("%d %d\n", inicio+1, finish+1)
			return
		}

	}
	fmt.Printf("-1\n")
	return

}

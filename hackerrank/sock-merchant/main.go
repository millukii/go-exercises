/*
John works at a clothing store. He has a large pile of socks that he must pair by color for sale.
Given an array of integers representing the color of each sock, determine how many pairs of socks
with matching colors there are.
For example, there are  socks with colors . There is one pair of color  and one of color .
There are three odd socks left, one of each color. The number of pairs is .

Function Description

Complete the sockMerchant function in the editor below. It must return an integer representing the number of
matching pairs of socks that are available.

sockMerchant has the following parameter(s):

n: the number of socks in the pile
ar: the colors of each sock

Output Format

Return the total number of matching pairs of socks that John can sell.

Sample Input

9
10 20 20 10 10 30 50 10 20
Sample Output

3
*/
package main

import (
	"fmt"
)

func main() {
	var n int
	//read form input number of elements
	fmt.Scan(&n)

	socks := make([]int, 101)

	for i := 0; i < n; i++ {
		var t int
		//read every number
		fmt.Scan(&t)
		//sum in t key
		socks[t]++
	}

	count := 0

	for _, v := range socks {
		count += v / 2
	}

	fmt.Println(count)
}

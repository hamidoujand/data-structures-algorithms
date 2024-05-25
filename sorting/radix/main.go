package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
Radix Sort: its not like others, those sorting algorithms are called "comparison sorting algorithms"
only can be applied to slice or array of NUMBERS, used to sort "binary" so if you can covert something to binary
you can use this on it. it works based on the number of "digits" inside of any number


Big O(n * k) "k" is the largest digits count inside of a number if the number is really large and have a lot of digits
we "k" grows to that. so most of the times can be ignored but in large numbers it has impact

*/

func main() {
	items := []int{23, 345, 5467, 12, 2345, 9852}
	fmt.Println(radixSort(items))
}

func radixSort(items []int) []int {
	maxDigits := mostDigits(items)
	for k := 0; k < maxDigits; k++ {
		buckets := make([][]int, 10)

		//we loop for each number and put them into buckets
		for i := 0; i < len(items); i++ {
			//we take each number at index[i] and ask for the digit at index k of that number
			digit := getDigit(items[i], k)
			// fmt.Printf("number %d at index %d has %d\n", items[i], k, digit)
			//use the digit to get the corresponding bucked from the bucket list
			bucket := buckets[digit]
			//then we push the number itself into that bucked
			bucket = append(bucket, items[i])
			//then we put it back
			buckets[digit] = bucket
		}
		//we need to put all of those number now into a new slice and reassign that slice to be "items"
		//again so we reorder them again and again till "k < maxDigits"
		results := make([]int, 0, 10)
		for _, bucket := range buckets {
			results = append(results, bucket...)
		}
		//now we re-assign it back to "items" variable
		items = results
	}
	return items
}

func digitCount(num int) int {
	str := strconv.Itoa(int(math.Abs(float64(num))))
	return len(str)
}

func mostDigits(nums []int) int {
	var maxDigits float64
	for _, num := range nums {
		maxDigits = math.Max(float64(digitCount(num)), maxDigits)
	}
	return int(maxDigits)
}

func getDigit(num int, place int) int {
	return int(math.Abs(float64(num))/math.Pow(10, float64(place))) % 10
}

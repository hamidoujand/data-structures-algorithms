package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	jobs := []Job{
		{
			deadline: 1,
			payment:  1,
		},
		{
			deadline: 2,
			payment:  2,
		},
		{
			deadline: 2,
			payment:  2,
		},
		{
			deadline: 7,
			payment:  1,
		},
		{
			deadline: 4,
			payment:  3,
		},
		{
			deadline: 4,
			payment:  5,
		},
		{
			deadline: 3,
			payment:  1,
		},
	}
	fmt.Println(optimalFreelancing(jobs))
}

type Job struct {
	deadline int
	payment  int
}

// O(n log n) Time | O(1) Space
func optimalFreelancing(jobs []Job) int {
	const weekdays = 7
	profit := 0
	timeline := [weekdays]bool{}
	//sort the jobs based on payment in descending order
	slices.SortFunc(jobs, func(j1 Job, j2 Job) int {
		return j2.payment - j1.payment
	})

	for _, job := range jobs {

		maxTimeWeHaveToFinishIt := int(math.Min(float64(job.deadline), weekdays)) - 1 // -1 for "idx out of bound"
		//now we need to put this job at the timeline based on the maxTimeWeHaveToFinishIt
		//so loop from that "maxTime" and see if in those idx inside of timeline we have a empty day or not
		//if yes then "true" that spot and add the payment to the profit
		for idx := maxTimeWeHaveToFinishIt; idx >= 0; idx-- {

			if !timeline[idx] {
				timeline[idx] = true
				profit += job.payment
				//break then
				break
			}
		}
	}

	return profit
}

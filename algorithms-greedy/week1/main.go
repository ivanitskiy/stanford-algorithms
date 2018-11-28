package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type job struct {
	weight int
	length int
}

var filePath string

var schedulesJobsDiffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Sum of weighted completion times of the resulting scheduled jobs in decreasing order of the difference (weight - length)",
	Long: `This file describes a set of jobs with positive and integral weights and lengths. It has the format

	[number_of_jobs]
	
	[job_1_weight] [job_1_length]
	
	[job_2_weight] [job_2_length]
	
	...
	
	For example, the third line of the file is "74 59", indicating that the second job has weight 74 and length 59.
	
	You should NOT assume that edge weights or lengths are distinct.
	
	Your task in this problem is to run the greedy algorithm that schedules jobs in decreasing order of the difference (weight - length). Recall from lecture that this algorithm is not always optimal. IMPORTANT: if two jobs have equal difference (weight - length), you should schedule the job with higher weight first. Beware: if you break ties in a different way, you are likely to get the wrong answer. You should report the sum of weighted completion times of the resulting schedule --- a positive integer --- in the box below.
	
	ADVICE: If you get the wrong answer, try out some small test cases to debug your algorithm (and post your test cases to the discussion forum).`,
	Run: func(cmd *cobra.Command, args []string) {
		jobs := readJobs()
		sort.Slice(jobs, func(i, j int) bool {
			diffi := jobs[i].weight - jobs[i].length
			diffj := jobs[j].weight - jobs[j].length
			if diffi == diffj {
				return jobs[i].weight > jobs[j].weight
			} else {
				return diffi > diffj
			}
		})

		// fmt.Println("Jobs:", jobs)

		var retval, lcount int64
		for _, v := range jobs {
			lcount = lcount + int64(v.length)
			retval = retval + int64(v.weight)*lcount
		}
		fmt.Println("By diff:", retval)
		// By diff: 69119377652
	},
}

var schedulesJobsRatioCmd = &cobra.Command{
	Use:   "ratio",
	Short: "Sum of weighted completion times of the resulting scheduled jobs in decreasing order of the ratio (weight/length)",
	Long: `For this problem, use the same data set as in the previous problem.

	Your task now is to run the greedy algorithm that schedules jobs (optimally) in decreasing order of the ratio (weight/length). In this algorithm, it does not matter how you break ties. You should report the sum of weighted completion times of the resulting schedule --- a positive integer --- in the box below.`,
	Run: func(cmd *cobra.Command, args []string) {
		jobs := readJobs()
		sort.Slice(jobs, func(i, j int) bool {
			return float64(jobs[i].weight)/float64(jobs[i].length) > float64(jobs[j].weight)/float64(jobs[j].length)
		})
		// fmt.Println("Jobs:", jobs)
		var retval, lcount int64

		for _, v := range jobs {
			lcount = lcount + int64(v.length)
			retval = retval + int64(v.weight)*lcount
		}
		fmt.Println("By ratio:", retval)
		// By ratio: 67311454237

	},
}

var mstCmd = &cobra.Command{
	Use:   "mst ",
	Short: "Runs Prim's minimum spanning tree",
	Long: `This file describes an undirected graph with integer edge costs. It has the format

	[number_of_nodes] [number_of_edges]
	
	[one_node_of_edge_1] [other_node_of_edge_1] [edge_1_cost]
	
	[one_node_of_edge_2] [other_node_of_edge_2] [edge_2_cost]
	
	...
	
	For example, the third line of the file is "2 3 -8874", indicating that there is an edge connecting vertex #2 and vertex #3 that has cost -8874.
	
	You should NOT assume that edge costs are positive, nor should you assume that they are distinct.
	
	Your task is to run Prim's minimum spanning tree algorithm on this graph. You should report the overall cost of a minimum spanning tree --- an integer, which may or may not be negative --- in the box below.
	
	IMPLEMENTATION NOTES: This graph is small enough that the straightforward O(mn) time implementation of Prim's algorithm should work fine. OPTIONAL: For those of you seeking an additional challenge, try implementing a heap-based version. The simpler approach, which should already give you a healthy speed-up, is to maintain relevant edges in a heap (with keys = edge costs). The superior approach stores the unprocessed vertices in the heap, as described in lecture. Note this requires a heap that supports deletions, and you'll probably need to maintain some kind of mapping between vertices and their positions in the heap.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, mstCmd!", filePath)
	},
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readJobs() []job {
	f, err := os.Open(filePath)
	checkError(err)
	defer f.Close()
	scanner := bufio.NewScanner(bufio.NewReader(f))
	scanner.Scan()
	var jobCount int
	jobCount, err = strconv.Atoi(scanner.Text())
	checkError(err)
	i := 0
	array := make([]job, jobCount)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		weight, err := strconv.Atoi(nums[0])
		checkError(err)
		length, err := strconv.Atoi(nums[1])
		array[i].weight = weight
		array[i].length = length
		i++
	}
	return array

}

func readMST() {
	f, err := os.Open(filePath)
	checkError(err)
	defer f.Close()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "algorithms-greedy-week1",
		Short: "Calculates Programming Assignment #1",
		Long:  "algorithms-greedy",
	}
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "File path (required)")
	rootCmd.MarkPersistentFlagRequired("file")

	rootCmd.AddCommand(schedulesJobsDiffCmd, schedulesJobsRatioCmd, mstCmd)
	rootCmd.Execute()

	// jobs := readJobs()
	// ratio: 190444405675
	// diff:  188636667966
	// ===================
	// 		 1807737709
	// ////////////////////
	// sort.Slice(jobs, func(i, j int) bool {
	// 	return float64(jobs[i].weight)/float64(jobs[i].length) < float64(jobs[j].weight)/float64(jobs[j].length)
	// })

	// lcount = 0
	// retval = 0
	// for _, v := range jobs {
	// 	lcount = lcount + v.length
	// 	retval = retval + int64(v.weight*lcount)
	// }
	// fmt.Println("By weight:", jobs, "\n", retval)

}

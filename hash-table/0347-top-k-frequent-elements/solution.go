package solution

import "sort"

func topKFrequent1(nums []int, k int) []int {
	numFreqMap := make(map[int]int)
	for _, num := range nums {
		numFreqMap[num]++
	}

	freqNumMap := make(map[int][]int, len(numFreqMap))
	for num, freq := range numFreqMap {
		freqNumMap[freq] = append(freqNumMap[freq], num)
	}
	freqList := make([]int, 0, len(freqNumMap))
	for freq := range freqNumMap {
		freqList = append(freqList, freq)
	}
	sort.Ints(freqList)

	res := make([]int, 0)
	for i := 0; i < len(freqList); i++ {
		res = append(res, freqNumMap[freqList[len(freqList)-i-1]]...)
		if len(res) >= k {
			break
		}
	}

	return res[:k]
}

func topKFrequent2(nums []int, k int) []int {
	numFreqMap := make(map[int]int)
	for _, num := range nums {
		numFreqMap[num]++
	}

	res := make([]int, 0, len(numFreqMap))
	for num := range numFreqMap {
		res = append(res, num)
	}

	sort.SliceStable(res, func(i, j int) bool {
		return numFreqMap[res[i]] > numFreqMap[res[j]]
	})

	return res[:k]
}

func topKFrequent3(nums []int, k int) []int {
	numFreqMap := make(map[int]int)
	for _, num := range nums {
		numFreqMap[num]++
	}

	bucket := make([][]int, len(nums)+1)
	for num, freq := range numFreqMap {
		bucket[freq] = append(bucket[freq], num)
	}

	res := make([]int, 0, len(nums))
	for freq := len(bucket) - 1; freq >= 0 && len(res) < k; freq-- {
		if len(bucket[freq]) != 0 {
			res = append(res, bucket[freq]...)
		}
	}

	return res
}

func topKFrequent4(nums []int, k int) []int {
	numFreqMap := make(map[int]int)
	for _, num := range nums {
		numFreqMap[num]++
	}

	freqNumMap := make(map[int][]int, len(numFreqMap))
	for num, freq := range numFreqMap {
		freqNumMap[freq] = append(freqNumMap[freq], num)
	}

	res := make([]int, 0, len(nums))
	for freq := len(nums); len(res) != k; freq-- {
		for _, num := range freqNumMap[freq] {
			if len(res) != k {
				res = append(res, num)
			}
		}
	}

	return res
}

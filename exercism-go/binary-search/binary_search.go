package binarysearch

func SearchInts(list []int, key int) int {
	start := 0
	end := len(list) - 1
	mid := (start + end) / 2
	for start <= end {
		curr := list[mid]
		if curr > key {
			end = mid - 1
			mid = (start + end) / 2
		} else if curr < key {
			start = mid + 1
			mid = (start + end) / 2
		} else {
			return mid
		}
	}
	return -1
}

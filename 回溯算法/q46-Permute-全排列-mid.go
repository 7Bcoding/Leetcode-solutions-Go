package 回溯算法

var arr []int
var result [][]int
var numsMap = map[int]int{}

func permute(nums []int) [][]int {
	result = [][]int{}
	backtrack(nums)
	return result
}

func backtrack(nums []int) [][]int {
	if len(arr) == len(nums) {
		// 为什么加入解集时，要将数组内容拷贝到一个新的数组里，再加入解集？
		// 因为该 path 变量存的是地址引用，结束当前递归时，将它加入 res 后，
		//该算法还要进入别的递归分支继续搜索，还要继续将这个 path 传给别的递归调用，
		//它所指向的内存空间还要继续被操作，所以 res 中的 path 的内容会被改变，这就不对。
		// 所以要弄一份当前的拷贝，放入 res，这样后续对 path 的操作，就不会影响已经放入 res 的内容。
		temp := make([]int, len(nums))
		copy(temp, arr)
		result = append(result, temp)
		return result
	}
	for i := 0; i < len(nums); i++ {
		if numsMap[nums[i]] == 0 {
			numsMap[nums[i]]++
			arr = append(arr, nums[i])
			backtrack(nums)
			arr = arr[:len(arr)-1]
			numsMap[nums[i]]--
		}
	}
	return result
}

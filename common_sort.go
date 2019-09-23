package main

import "fmt"

//根据文章：https://juejin.im/post/5a30c3a7518825569539a319

//冒泡排序
func bubble_sort(array []int) error {
	for i, _ := range array {
		for j := len(array) - 1; j > i; j-- { //从右往左把最小的值放到i的位置。
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	return nil
}

//插入排序
func insert_sort(array []int) error {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- { //每次把i位置的值插入到i前面的序列中正确的位置
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	return nil
}

//快速排序
func quick_sort(array []int, start, end int) error {
	if start >= end {
		return nil
	}
	left := start
	right := end
	pivot_pos := start //(start+end)/2
	pivot := array[pivot_pos]
	for left < right { //left和right最后会碰头
		for array[right] >= pivot && left < right {
			right--
		} //因为先从右边往左边找，所以最后的那个值是小于等于pivot的
		for array[left] <= pivot && left < right {
			left++
		}
		array[left], array[right] = array[right], array[left]
	}
	array[left], array[pivot_pos] = array[pivot_pos], array[left] //因为停下来的值小于等于pivot，所以将pivot和left的值交换
	//递归子序列
	quick_sort(array, start, left-1)
	quick_sort(array, left+1, end)

	return nil
}

// 归并排序
func mergeSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    m := len(arr)/ 2
    return mergeTwo(mergeSort(arr[:m]), mergeSort(arr[m:]))
}

func mergeTwo(left, right []int) []int{
    out := make([]int, 0, len(left) + len(right))
    li, ri := 0, 0
    for li < len(left) && ri < len(right) {
        if left[li] < right[ri] {
            out = append(out, left[li])
            li ++
        }else {
            out = append(out, right[ri])
            ri ++
        }
    }
    if li < len(left){
        out = append(out, left[li:]...)
    }else if ri < len(right) {
        out = append(out, right[ri:]...)
    }
    return out
}

//调整堆为最大堆
func adjust_heap(array []int, pos, length int) {
	for i := pos*2 + 1; i < length; i = 2*pos + 1 { //交换了子节点和父节点后可能会影响下面的子堆，所以还要去调整子堆
		if i+1 < length && array[i] < array[i+1] {
			i++
		}
		if array[pos] < array[i] {
			array[pos], array[i] = array[i], array[pos]
			pos = i
		} else {
			break
		}
	}
}

//堆排序
func heap_sort(array []int) {
	arr_len := len(array)
	for i := arr_len/2 - 1; i >= 0; i-- { //从最下面的非叶子节点开始向上调整堆
		adjust_heap(array, i, arr_len)
	}
	for i := arr_len - 1; i > 0; i-- { //将堆顶元素放到数组末尾，然后不断调整为最大堆，最后就形成了升序的数组
		array[0], array[i] = array[i], array[0]
		adjust_heap(array, 0, i)
	}
}

//计数排序
//时间复杂度O(n+k)，n是数组大小，k是数组里最大数，时间复杂度很牛逼了，只是基数排序局限于整数排序,并且k不是很大的时候效果更好
//下面这个例子只适合正整数排序

func count_sort(array []int) {
	max := 0
	arr_len := len(array)
	for _,i := range array {
		if i > max {
			max = i
		}
	}
	count_arr := make([]int, max+1)
	for _,i := range array {
		count_arr[i]++
	}
	for i := 1; i < max+1; i++ {//获取不同的数在有序序列中的位置
		count_arr[i] = count_arr[i] + count_arr[i-1]
	}
	sorted := make([]int, arr_len)
	for i := range array {
		count_arr[array[i]]--
		sorted[count_arr[array[i]]] = array[i]
	}
	copy(array, sorted)

}

func main() {
	var to_sort = []int{9, 4, 5, 6, 345, 2, 77, 3, 56, 8, 11, 335}
	fmt.Println(to_sort)
	//bubble_sort(to_sort)
	//insert_sort(to_sort)
	//quick_sort(to_sort, 0, len(to_sort)-1)
	//heap_sort(to_sort)
	count_sort(to_sort)
	fmt.Println(to_sort)
}

package main

import "fmt"

// 希尔排序通过将比较的全部元素分为几个区域来提升插入排序的性能。
// 这样可以让一个元素可以一次性地朝最终位置前进一大步。
// 然后算法再取越来越小的步长进行排序，
// 算法的最后一步就是普通的插入排序，但是到了这步，需排序的数据几乎是已排好的了（此时插入排序较快）。

// 假设有一个很小的数据在一个已按升序排好序的数组的末端。
// 如果用复杂度为O(n2)的排序（冒泡排序或插入排序），
// 可能会进行n次的比较和交换才能将该数据移至正确位置。
// 而希尔排序会用较大的步长移动数据，所以小数据只需进行少数比较和交换即可到正确位置。

// 一个更好理解的希尔排序实现：
// 将数组列在一个表中并对列排序（用插入排序）。
// 重复这过程，不过每次用更长的列来进行。
// 最后整个表就只有一列了。
// 将数组转换至表是为了更好地理解这算法，
// 算法本身仅仅对原数组进行排序（通过增加索引的步长，例如是用i += step_size而不是i++）。
func main() {
	arrs := [][]int{
		{6, 5, 4, 3, 1, 2},
		{22, 34, 3, 40, 18, 4},
		{},
		{1, 2, 3},
	}
	for _, arr := range arrs {
		shellSort(arr)
		fmt.Println(arr)
	}
}

func shellSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	step := len(arr) / 2 // 初始步长
	for ; step > 0; step /= 2 {
		// 进行插入排序
		for i := step; i < len(arr); i += step {
			value := arr[i] // 准备要插入的数字
			j := i - step   // 左边有序数组的长度
			for ; j >= 0; j -= step {
				if value < arr[j] {
					arr[j+step] = arr[j] // 右移
				} else {
					break
				}
			}
			arr[j+step] = value
		}
	}
}

package leet_0004

import "math"

/**
* There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]
The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5

O(log(min(m,n)))
参考博客：http://blog.csdn.net/chen_xinjia/article/details/69258706

index:  0	1	2	3	4	5
		  	L1   R1
num1:	3 	5 |	8 	9    	      4  cut1：左有几个元素
num2:	1	2 	7 | 10  11  12     6  cut2：左有几个元素
			   L2  R2

num3:  1 2 3 5 7 | 8 9 10 11 12

num3 -> num1 + num2 -> num1

L1 <= R2
L2 <= R1


(cutL, cutR) cut1在这个区间取值
L1 > R2 cut1 <<  (cutL, cut1 - 1)
L2 > R1 cut2 >>  (cut1 + 1, cutR)
index: 0	1	2	3	4	5
		  L1   R1
num1:	3 	5 |	8 	9   |	      4  cut1：左有几个元素
num2:	1	2 	7 | 10  11  12     6  cut2：左有几个元素
			   L2  R2
num3:  1 2 3 5 7 | 8 9 10 11 12
int cut1 = 2;
int cut2 = 3;
int cutL = 0;
int cutR = 4;
time : O(log(min(m,n)))
space : O(1)
*/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	// 确保nums1比nums2短，复杂度：log(min(m,n))
	if n1 > n2 {
		return FindMedianSortedArrays(nums2, nums1)
	}

	n, cut1, cut2, l1, r1 := n1+n2, 0, 0, 0, n1
	L1, R1, L2, R2 := 0, 0, 0, 0
	for cut1 <= n1 { // 跟传统的二分不一样
		// 计算cut1和cut2的长度
		cut1 = l1 + (r1-l1)/2
		cut2 = n/2 - cut1
		// 计算L1,R1,L2,R2，注意边界情况
		if cut1 == 0 {
			L1 = math.MinInt32
		} else {
			L1 = nums1[cut1-1]
		}
		if cut1 == n1 {
			R1 = math.MaxInt32
		} else {
			R1 = nums1[cut1]
		}
		if cut2 == 0 {
			L2 = math.MinInt32
		} else {
			L2 = nums2[cut2-1]
		}
		if cut2 == n2 {
			R2 = math.MaxInt32
		} else {
			R2 = nums2[cut2]
		}
		// 根据上面的分析，当L1<=R2 and L2<=R1时，
		// 满足退出的条件，否则要移动cut1和cut2
		if L1 > R2 { // L1太大，cut1左移
			r1 = cut1 - 1
		} else if L2 > R1 { // R1太小，cut1右移
			l1 = cut1 + 1
		} else { // 满足退出条件
			break
		}
	}
	// 总长度为偶数
	if n%2 == 0 {
		return float64(math.Max(float64(L1), float64(L2))+
			math.Min(float64(R1), float64(R2))) / 2
	}
	// 总长度为奇数
	return float64(math.Min(float64(R1), float64(R2)))
}

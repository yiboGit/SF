package main

import (
	"sort"
)

func main() {

}

//  Definition for an interval:
type Interval struct {
	Start int
	End   int
}

// 思路 就是说 既然每个元素是区间，那么我先按照区间的start有小到大排序
// 这样 在我循环区间（Interval）数组的时候 start是按照由小到大进行循环的，这个很重要
// 这时候 我只需要 循环的元素与结果集中最后一个区间元素 （原index:0的元素） 进行比较
// 如果 循环元素start > 结果集中最后一个元素end，那么将循环元素按照顺序放入结果集
// 如果 循环元素start <= 结果集中最后一个元素end；并且循环元素end <= 结果集中最后一个元素end,那么将结果集中元素的end 变为 循环元素end 其余情况则不变

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	result := []Interval{intervals[0]}

	for _, inter := range intervals {
		// 由小到大已排好序
		// start < end 说明有交集
		if inter.Start <= result[len(result)-1].End {
			// inter.end > result(len-1).end 说明需要改变 结果集中的end 扩大范围
			//           <        说明inter区间已被包含，不需要任何改变
			if inter.End > result[len(result)-1].End {
				result[len(result)-1].End = inter.End
			}
		} else {
			//无交集 顺序加入，此后循环的元素 只可能和这个新放入的有交集
			result = append(result, inter)
		}
	}
	return result
}

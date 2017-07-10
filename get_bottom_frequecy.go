/*************************************************************************************************
在横线上填写数字，使之符合要求。
要求如下：对应的数字下填入的数，代表上面的数在下面出现的次数，比如3下面是1，代表3要在下面出现一次。
正确答案是：0 1 2 3 4 5 6 7 8 9
         6 2 1 0 0 0 1 0 0 0
思路：因为第二行的数字是第一行的数在下面出现的次数，下面10个格子，总共10次。。。所以第2排数字之和为10。
首先从0入手，先填9，肯定不可能，9下面要是1，只剩8个位填0，不够填8，8下面要填1，1要至少填2，
后面不用再想，因为已经剩下7个位置，不够填0……如此类推。到0下面填6的时候就得到我上面的答案了。。
其实可以推出这个题目的两个关键条件：
1、第2排数字之和为10。
2、两排数字上下相乘之和也是10！
满足这两个条件的就是答案，下面来编写程序实现！
**************************************************************************************************/

package main

import (
	"fmt"
)

type NumberTB struct {
	Num     []int
	Len     int
	Success bool
}

func NewNumberTB(len int) *NumberTB {
	return &NumberTB{
		Num:     make([]int, len),
		Len:     len,
		Success: false,
	}
}

func (this *NumberTB) GetBottom() {
	for !this.Success {
		this.setNextBottom()
	}
}

func (this *NumberTB) setNextBottom() {
	reb := true
	for i := 0; i < this.Len; i++ {
		frequecy := this.getFrequecy(i)
		if this.Num[i] != frequecy {
			this.Num[i] = frequecy
			reb = false
		}
	}

	this.Success = reb
}

func (this *NumberTB) getFrequecy(num int) int {
	count := 0
	for i := 0; i < this.Len; i++ {
		if this.Num[i] == num {
			count++
		}
	}

	return count
}

func main() {
	numtb := NewNumberTB(10)
	numtb.GetBottom()
	fmt.Println(numtb.Num)
}

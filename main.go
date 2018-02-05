package main

import (
	"fmt"
)

func main()  {
c:=make(chan int,4)
	cr:=make(chan int,4)
c<-1
c<-8
c<-76
c<-(-123)
cr=sort(c,4)
	for i:=0;i<4 ;i++  {
		fmt.Println(<-cr)

	}
}

func sort(m chan int,size int) chan int{
	if size <= 1 {
		return m
	}
	left_size:=0
	right_size:=0
	if size%2==0 {
		left_size=size/2
		right_size=size/2
	}else {
		left_size=int(size/2)
		right_size=int(size/2)+1
	}
left_channel:=make(chan int,left_size)
right_channel:=make(chan int,right_size)

		for i:=0;i<left_size;i++{
			left_channel<-<-m
		}
		for i:=0;i<right_size;i++{
			right_channel<-<-m
		}
left:=sort(left_channel,left_size)
right:=sort(right_channel,right_size)
	arr := merge(left, right,left_size,right_size)
	return arr
}

func merge(A, B chan int,size_left int,size_right int) chan int {
	arr := make(chan int, size_right+size_left)
	size_queuearr:=0
	// index j for A, k for B
	j, k := 0, 0

	for i := 0; i < size_left+size_right; i++ {
		// fix for index out of range without using sentinel
		if j >= size_left {
			setItem(arr,i,size_right+size_left,getItem(B,k,size_right),size_queuearr)
			size_queuearr++
			k++
			continue
		} else if k >= size_right {
			setItem(arr,i,size_right+size_left,getItem(A,j,size_left),size_queuearr)
			size_queuearr++
			j++
			continue
		}
		// default loop condition

		if getItem(A,j,size_left) > getItem(B,k,size_right) {
			setItem(arr,i,size_right+size_left,getItem(B,k,size_right),size_queuearr)
			size_queuearr++
			k++
		} else {


			setItem(arr,i,size_right+size_left,getItem(A,j,size_left),size_queuearr)
			size_queuearr++
			j++
		}
	}

	return arr
}
func getItem(m chan int,item int,size int)int  {
	replace:=make(chan int,size)
ret:=0
	for i:=0;i<size;i++ {
		replace<-<-m
	}
	for i:=0;i<size;i++ {

		if i==item {
			ret=<-replace
			m<-ret
			continue
		}
		m<-<-replace
	}
	return ret
}
func setItem(m chan int,item int,size int,input int,fill int)  {
	replace:=make(chan int,size)
	for i:=0;i<item;i++ {
replace<-<-m
	}

		replace<-input

	for i:=item;i<fill;i++ {
		replace<-<-m
	}
	for i:=0;i<=fill;i++ {
		m<-<-replace
	}


}
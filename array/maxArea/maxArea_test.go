package main

import "testing"

func TestHelloWorld(t *testing.T) {
	height := []int{1,8,6,2,5,4,8,3,7}
	s:= maxArea(height)
}
//暴力解法
func maxArea1(height []int) int {
    var s int
    for i:=0 ; i<len(height);i++ {
        for j:=i+1 ; j<len(height);j++{
            h := 0
            if height[i] >= height[j]{
                h = height[j]
            }else{
                h = height[i]
            }
            w := j-i
            cur := h*w
            if s < cur {
                s = cur
            }
        }
    }
    return s
}


//双指针，夹逼法
func maxArea2(height []int) int {
    var h,s int
    i:=0
    j:=len(height)-1

    for i<j {
        if height[i] >= height[j] {
            h = height[j]
            j--
        }else{
            h = height[i]
            i++
        }
        cur := h*(j-i+1)
        if cur > s {
            s = cur
        }
    }
    return s
}
}

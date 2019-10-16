package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}
//循环
func climbStairs(n int) int {
    if n <=  2{
        return n
    }
	ai,a1,a2 :=0,1,2
    for i:=3 ;i<=n ;i++{
        // 写成一行是错误的 ai,a1,a2 = a1+a2,a2,ai
        ai,a1 = a1+a2,a2
		a2 = ai
    }
    return ai
}

//递归缓存
func climbStairs1(n int) int {
    cache := make(map[int]int,n)
    return climbStairsWithCache(n, cache)
}

func climbStairsWithCache(n int , cache map[int]int ) int {
    if n <=  2{
        return n
    }
    if value,ok := cache[n]; ok {
        return value
    }
    cache[n] = climbStairsWithCache(n-1,cache) + climbStairsWithCache(n-2,cache)
    return cache[n]
}

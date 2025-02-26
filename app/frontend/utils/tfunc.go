package utils

import (
	"fmt"
	"strconv"
)

func IndexMod(idx int, m int) int {
	return idx % m
}

func RangeSlice(s, l int) []int {
	r := make([]int, l)
	for i := 0; i < l; i++ {
		r[i] = s + i
	}
	return r
}

func ShowPages(curPage, totalPage interface{}) []int {

	var ci int
	var ti int
	var ok bool
	if ci, ok = curPage.(int); !ok {
		ci, _ = strconv.Atoi(curPage.(string))
	}

	if ti, ok = totalPage.(int); !ok {
		ti, _ = strconv.Atoi(totalPage.(string))
	}
	pageSlice := make([]int, min(ti, 5))
	if ti == 1 {
		return []int{1}
	}

	if ti < 5 {
		for i := 0; i < ti; i++ {
			pageSlice[i] = i + 1
		}
		return pageSlice
	}

	var startPage = ci - min(ci-1, max(4+ci-ti, 2))

	for i := 0; i < 5; i++ {
		pageSlice[i] = startPage + i
	}

	return pageSlice
}

func SubOne(n int) int {
	return n - 1
}

func AddOne(n int) int {
	return n + 1
}

func InSlice[T int | string](ele T, collection []T) bool {
	for _, e := range collection {
		if e == ele {
			return true
		}
	}
	return false
}

func RowIdx(n int, rowLen int) int {
	return n/rowLen + 1
}

func ColIdx(n int, colLen int) int {
	return n%colLen + 1
}

func CalcPrice(p string, q string) string {
	price, err := strconv.ParseFloat(p, 64)
	if err != nil {
		return err.Error()
	}
	quantity, err := strconv.ParseInt(q, 10, 32)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%.2f", price*float64(quantity))
}

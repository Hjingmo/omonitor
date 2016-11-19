package utils

import (
	"fmt"
	"strconv"
)

func UpdateCompareInt64(oldIds, newIds []int64) (addIds, delIds []int64) {
	for _, o := range oldIds {
		ock := CheckIdInt64(newIds, o)
		if !ock {
			delIds = append(delIds, o)
		}
	}
	for _, n := range newIds {
		nck := CheckIdInt64(oldIds, n)
		if !nck {
			addIds = append(addIds, n)
		}
	}
	return
}

func CheckIdInt(ids []int, cId int) bool {
	for _, i := range ids {
		if i == cId {
			return true
		}
	}
	return false
}

func CheckIdInt64(ids []int64, cId int64) bool {
	for _, i := range ids {
		if i == cId {
			return true
		}
	}
	return false
}

func UpdateCompareInt(oldIds []int, newIds []int) (addIds, delIds []int) {
	for _, o := range oldIds {
		ock := CheckIdInt(newIds, o)
		if !ock {
			delIds = append(delIds, o)
		}
	}
	for _, n := range newIds {
		nck := CheckIdInt(oldIds, n)
		if !nck {
			addIds = append(addIds, n)
		}
	}
	return
}

func StringToInt(stringIds []string) []int {
	var result []int
	for _, i := range stringIds {
		intI, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, intI)
	}
	return result
}

func StringToInt64(stringIds []string) []int64 {
	var result []int64
	for _, i := range stringIds {
		intI, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, intI)
	}
	return result
}

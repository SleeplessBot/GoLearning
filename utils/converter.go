package utils

import (
	"strconv"
	"strings"
)

func Uint2Str(value uint) string {
	return strconv.FormatInt(int64(value), 10)
}

func Int2Str(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

func Int642Str(value int64) string {
	return strconv.FormatInt(value, 10)
}

func Csv2Strs(csv string) []string {
	return strings.Split(csv, ",")
}

func Strs2Csv(strVals []string) string {
	return strings.Join(strVals, ",")
}

func Csv2Int64s(csv string) ([]int64, error) {
	strVals := strings.Split(csv, ",")
	var intVals []int64
	for _, strVal := range strVals {
		intVal, err := strconv.ParseInt(strVal, 10, 64)
		if err != nil {
			return nil, err
		}
		intVals = append(intVals, intVal)
	}
	return intVals, nil
}

func Csv2Uints(csv string) ([]uint, error) {
	strVals := strings.Split(csv, ",")
	var intVals []uint
	for _, strVal := range strVals {
		intVal, err := strconv.ParseInt(strVal, 10, 64)
		if err != nil {
			return nil, err
		}
		intVals = append(intVals, uint(intVal))
	}
	return intVals, nil
}

func Int64s2Csv(intVals []int64) string {
	var strVals []string
	for _, intVal := range intVals {
		strVals = append(strVals, Int642Str(intVal))
	}
	return strings.Join(strVals, ",")
}

func Str2Int(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

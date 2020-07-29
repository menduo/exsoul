// Created by @menduo @ 2020/6/19
package exsoul

import (
	"strconv"
	"strings"
	"time"
)

// isEmptyStrList 是不是一个空的字符串列表，如果列表为空，或者每个值 trim 后都是空字符串。
func isEmptyStrList(row []string) bool {
	var isEmpty = true
	if len(row) == 0 {
		return true
	}

	for _, s := range row {
		if strings.Trim(s, " ") != "" {
			isEmpty = false // 只要有一行不为空，就认为它是不是一个 空 行
			return isEmpty
		}
	}

	return isEmpty
}

// getStrFromArray 从数组中获取值
func getStrFromArray(row []string, idx int) string {
	if len(row) < idx+1 {
		return ""
	}
	return row[idx]
}

// validateMonth 验证 month 字符串
func validateMonth(s string, layouts ...string) (time.Time, error) {
	var layout = "2006-01"
	if len(layouts) > 0 {
		layout = layouts[0]
	}

	t, err := time.ParseInLocation(layout, s, time.Local)
	return t, err
}

// validateDate 验证日期字符串
func validateDate(s string, layouts ...string) (time.Time, error) {
	var layout = "2006-01-02"
	if len(layouts) > 0 {
		layout = layouts[0]
	}

	t, err := time.ParseInLocation(layout, s, time.Local)
	return t, err
}

func toInt(val string) (rval int, err error) {
	val = strings.TrimSpace(val)
	rval, err = strconv.Atoi(val)
	return rval, err
}

func toInt32(val string) (rval int32, err error) {
	val = strings.TrimSpace(val)
	tempVal, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return
	}
	rval = int32(tempVal)
	return rval, err
}

func toInt64(val string) (rval int64, err error) {
	val = strings.TrimSpace(val)
	rval, err = strconv.ParseInt(val, 10, 64)
	return rval, err
}

func toFloat32(pval string) (rval float32, err error) {
	pval = strings.TrimSpace(pval)
	rval64, err := strconv.ParseFloat(pval, 32)
	if err != nil {
		return
	}
	rval = float32(rval64)
	return rval, err
}

func toFloat64(pval string) (rval float64, err error) {
	pval = strings.TrimSpace(pval)
	rval, err = strconv.ParseFloat(pval, 64)
	return rval, err
}

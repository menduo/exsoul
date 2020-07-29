// Created by @menduo @ 2020/6/19
package exsoul

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type Row struct {
	colList []string
}

func (r *Row) GetCol(colIdx int) string {
	return getStrFromArray(r.colList, colIdx)
}

func (r *Row) IsEmptyRow() bool {
	return isEmptyStrList(r.colList)
}

// GetColAsInt col->str->int32
func (r *Row) GetColAsStr(colIdx int, dftVal ...string) (string, error) {
	var err error
	var val string

	if len(r.colList) < colIdx+1 {
		if len(dftVal) > 0 {
			val = dftVal[0]
			return val, nil
		}
		err = ErrorOutOfRange
		return "", err
	}
	val = r.GetCol(colIdx)
	return val, err
}

// GetColAsInt col->str->int32
func (r *Row) GetColAsInt(colIdx int, dftVal ...int) (int, error) {
	val, err := toInt(r.GetCol(colIdx))
	if err != nil {
		if len(dftVal) > 0 {
			val = dftVal[0]
		}
	}
	return val, err
}

// GetColAsInt32 col->str->int32
func (r *Row) GetColAsInt32(colIdx int, dftVal ...int32) (int32, error) {
	val, err := toInt32(r.GetCol(colIdx))
	if err != nil {
		if len(dftVal) > 0 {
			val = dftVal[0]
		}
	}
	return val, err
}

// GetColAsInt64 col到int64
func (r *Row) GetColAsInt64(colIdx int, dftVal ...int64) (int64, error) {
	val, err := toInt64(r.GetCol(colIdx))
	if err != nil {
		if len(dftVal) > 0 {
			val = dftVal[0]
		}
	}
	return val, err
}

// GetColAsFloat 到 float, float64
func (r *Row) GetColAsFloat(colIdx int, dftVal ...float64) (float64, error) {
	val, err := toFloat64(r.GetCol(colIdx))
	if err != nil {
		if len(dftVal) > 0 {
			val = dftVal[0]
		}
	}
	return val, err
}

// GetColAsFloat32 转为 flot32
func (r *Row) GetColAsFloat32(colIdx int, dftVal ...float32) (float32, error) {
	val, err := toFloat32(r.GetCol(colIdx))
	if err != nil {
		if len(dftVal) > 0 {
			val = dftVal[0]
		}
	}
	return val, err
}

// GetColAsFloat64 转为 flot64
func (r *Row) GetColAsFloat64(colIdx int, dftVal ...float64) (float64, error) {
	val, err := toFloat64(r.GetCol(colIdx))
	if err != nil {
		if len(dftVal) > 0 {
			val = dftVal[0]
		}
	}
	return val, err
}

// GetColAsYuanToFen 从 元 转成 分。把 cell 的值转为 float64 ，再 乘以 100?
func (r *Row) GetColAsYuanToFen(colIdx int, dftVal ...int64) (int64, error) {
	return r.GetColAsYuanToFenWithDigit(colIdx, 2, dftVal...)
}

// GetColAsYuanToFenWithDigit 从 元 转成 分。把 cell 的值转为 float64 ，再 乘以 100?
func (r *Row) GetColAsYuanToFenWithDigit(colIdx, digit int, dftVal ...int64) (int64, error) {
	var value int64
	tval, err := toFloat64(r.GetCol(colIdx))
	if err != nil {
		if len(dftVal) > 0 {
			value = dftVal[0]
		}
	}

	digit2, _ := decimal.NewFromString(fmt.Sprintf("1%0*d", digit, 0))
	value = decimal.NewFromFloat(tval).Mul(digit2).IntPart()
	return value, err
}

// GetColAsFenToYuan 从 分转成 元。把 cell 的值转为 int64，再除以 100?
func (r *Row) GetColAsFenToYuan(colIdx int, dftVal ...float64) (float64, error) {
	return r.GetColAsFenToYuanWithDigit(colIdx, 2, dftVal...)
}

// GetColAsFenToYuanWithDigit 从 分转成 元。把 cell 的值转为 int64，再除以 100?
func (r *Row) GetColAsFenToYuanWithDigit(colIdx, digit int, dftVal ...float64) (float64, error) {
	var value float64
	tval, err := toInt64(r.GetCol(colIdx))
	if err != nil {
		if len(dftVal) > 0 {
			value = dftVal[0]
		}
	}

	digit2, _ := toInt64(fmt.Sprintf("1%0*d", digit, 0))
	value, _ = decimal.NewFromInt(tval).Div(decimal.NewFromInt(digit2)).Float64()
	return value, err
}

// IsColValidMonthStr 是否为有效 月份 字符串
func (r *Row) IsColValidMonthStr(colIdx int, layouts ...string) (bool, error) {
	if _, err := validateMonth(r.GetCol(colIdx), layouts...); err != nil {
		return false, err
	}
	return true, nil
}

// IsColValidDateStr 是否为有效 日期 字符串
func (r *Row) IsColValidDateStr(colIdx int, layouts ...string) (bool, error) {
	if _, err := validateDate(r.GetCol(colIdx), layouts...); err != nil {
		return false, err
	}
	return true, nil
}

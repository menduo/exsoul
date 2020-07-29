// Created by @menduo @ 2020/7/14
package exsoul

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
)

// NewFile 为写文件而生！
func NewFile() *Exsoul {
	esfile := &Exsoul{
		eObj:      excelize.NewFile(),
		sheetRows: nil,
	}
	return esfile
}

// SetHeader 设置表头
func (e *Exsoul) SetHeader(values *[]interface{}) error {
	sheet := e.eObj.GetSheetName(0)
	return e.SetRowByColForSheet(sheet, 1, 1, values)
}

// SetHeaderForSheet 设置指定 sheet 的表头
func (e *Exsoul) SetHeaderForSheet(sheet string, values *[]interface{}) error {
	return e.SetRowByColForSheet(sheet, 1, 1, values)
}

// SetHeader 设置表头
func (e *Exsoul) SetHeaderByCol(col int, values *[]interface{}) error {
	sheet := e.eObj.GetSheetName(0)
	return e.SetRowByColForSheet(sheet, col, 1, values)
}

// SetHeaderForSheet 设置指定 sheet 的表头
func (e *Exsoul) SetHeaderByColForSheet(sheet string, col int, values *[]interface{}) error {
	return e.SetRowByColForSheet(sheet, col, 1, values)
}

// SetRow 为默认sheet 写入行，lineNum 为行号（不是索引下标），values 为切片指针
func (e *Exsoul) SetRow(lineNum int, values *[]interface{}) error {
	sheet := e.eObj.GetSheetName(0)
	return e.SetRowByColForSheet(sheet, 1, lineNum, values)
}

// SetRowForSheet 为指定sheet，写入行，lineNum 为行号（不是索引下标），values 为切片指针
func (e *Exsoul) SetRowForSheet(sheet string, lineNum int, values *[]interface{}) error {
	return e.SetRowByColForSheet(sheet, 1, lineNum, values)
}

// SetRowByCol 为默认sheet 写入行，lineNum 为行号（不是索引下标），col为每行从第多少列开始，values 为切片指针
func (e *Exsoul) SetRowByCol(col, lineNum int, values *[]interface{}) error {
	sheet := e.eObj.GetSheetName(0)
	return e.SetRowByColForSheet(sheet, col, lineNum, values)
}

// SetRowByColForSheet 为指定sheet，写入行，lineNum 为行号（不是索引下标），col为每行从第多少列开始，values 为切片指针
func (e *Exsoul) SetRowByColForSheet(sheet string, col, lineNum int, values *[]interface{}) error {
	axis, err := excelize.CoordinatesToCellName(col, lineNum)
	if err != nil {
		return err
	}
	err = e.eObj.SetSheetRow(sheet, axis, values)
	return err
}

// SetRowByList 为默认 sheet 批量写入行，从第1行开始
func (e *Exsoul) SetRowByList(vlist []*[]interface{}) error {
	sheet := e.eObj.GetSheetName(0)
	return e.SetRowByListForSheet(sheet, vlist)
}

// SetRowByListForSheet 为指定sheet批量写入行
func (e *Exsoul) SetRowByListForSheet(sheet string, vlist []*[]interface{}) error {
	if len(vlist) == 0 {
		return nil
	}

	for i, v := range vlist {
		lineNum := i + 1
		axis, _ := excelize.CoordinatesToCellName(1, lineNum)
		err := e.eObj.SetSheetRow(sheet, axis, v)
		if err != nil {
			return newError("[exs]error while set value for `%s`: `%v`", axis, v)
		}
	}
	return nil
}

// SaveAs 保存到 filepath 文件中
func (e *Exsoul) SaveAs(filepath string) error {
	return e.eObj.SaveAs(filepath)
}

// SaveToWriter 保存至 writer 接口
func (e *Exsoul) SaveToWriter(writer io.Writer) error {
	return e.eObj.Write(writer)
}

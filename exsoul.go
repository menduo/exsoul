// Created by @menduo @ 2020/6/19
package exsoul

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
	"io/ioutil"
	"os"
)

// Exsoul
type Exsoul struct {
	eObj      *excelize.File //
	sheetRows map[int][]*Row //
}

func (e *Exsoul) GetExcelObj() *excelize.File {
	return e.eObj
}

// NewFromFile 从文件创建
func NewFromFile(filepath string) (*Exsoul, error) {
	fobj, err := os.OpenFile(filepath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, newParamErrorWithError("Invalid file object", err)
	}
	defer fobj.Close()
	return NewFromReader(fobj)
}

// NewFromReader 从 reader 接口创建
func NewFromReader(reader io.Reader) (*Exsoul, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil || f == nil {
		return nil, newParamErrorWithError("Invalid excel object", err)
	}
	return newExcelService(f), nil
}

// newExcelService 内部方法
func newExcelService(eFile *excelize.File) *Exsoul {
	return &Exsoul{
		eObj:      eFile,
		sheetRows: make(map[int][]*Row, len(eFile.GetSheetList())),
	}
}

// LoadRowsFromSheetIndex 按 sheet 的索引读取 所有行
func (e *Exsoul) LoadRowsFromFirstSheet() []*Row {
	index := 0
	firstSheetName := e.eObj.GetSheetName(index)
	rawRowList, _ := e.eObj.GetRows(firstSheetName)

	if vList, has := e.sheetRows[index]; has {
		return vList
	}

	// 也可以使用 excelize 的 Rows 方法以及对应的。但感觉比较复杂一些?
	rowObjList := make([]*Row, 0)
	for _, rowOne := range rawRowList {
		if isEmptyStrList(rowOne) {
			continue
		}
		rowObjList = append(rowObjList, &Row{rowOne})
	}

	e.sheetRows[index] = rowObjList

	return rowObjList
}

// LoadRowsFromSheetIndex 按 sheet 的索引读取 所有行
func (e *Exsoul) LoadRowsFromSheetIndex(index int) []*Row {
	firstSheetName := e.eObj.GetSheetName(index)
	rawRowList, _ := e.eObj.GetRows(firstSheetName)

	if vList, has := e.sheetRows[index]; has {
		return vList
	}

	// 也可以使用 excelize 的 Rows 方法以及对应的。但感觉比较复杂一些?
	rowObjList := make([]*Row, 0)
	for _, rowOne := range rawRowList {
		if isEmptyStrList(rowOne) {
			continue
		}
		rowObjList = append(rowObjList, &Row{rowOne})
	}

	e.sheetRows[index] = rowObjList

	return rowObjList
}

// LoadRowsFromSheetName 按 sheet 的 名字 读取 所有行
func (e *Exsoul) LoadRowsFromSheetName(sheetName string) []*Row {
	idx := e.eObj.GetSheetIndex(sheetName)
	return e.LoadRowsFromSheetIndex(idx)
}

// Clone 克隆一个
func (e *Exsoul) Clone() (*Exsoul, error) {
	f, err := ioutil.TempFile("", "menduo-tmp-*.xlsx")

	if err != nil {
		return nil, err
	}

	defer f.Close()

	err = e.eObj.SaveAs(f.Name())
	if err != nil {
		return nil, err
	}

	esfile, err := NewFromReader(f)
	return esfile, err
}

package main

import "fmt"

var EmptyByte = []byte(`.`)[0]

func isValidSudoku(board [][]byte) bool {
	su := Su{}
	su.init()
	for indRow, rows := range board {
		for indCol, val := range rows {
			if val != EmptyByte && !su.validCell(indRow, indCol, val) {
				return false
			}
		}
	}
	return true
}

type Do map[byte]struct{}

func (m *Do) Push(item byte) bool {
	_, exists := (*m)[item]
	if exists {
		return false
	} else {
		(*m)[item] = struct{}{}
		return true
	}
}

type Su struct {
	EmptyByte byte
	Rows      []Do
	Cols      []Do
	Grid      []Do
}

func (m *Su) init() {
	m.EmptyByte = []byte(`.`)[0]
	m.Rows = make([]Do, 9)
	for i := range m.Rows {
		m.Rows[i] = Do{}
	}
	m.Cols = make([]Do, 9)
	for i := range m.Cols {
		m.Cols[i] = Do{}
	}
	m.Grid = make([]Do, 9)
	for i := range m.Grid {
		m.Grid[i] = Do{}
	}
}

func (m *Su) validCell(indRow, indCol int, val byte) bool {
	if m.pushRow(indRow, val) &&
		m.pushCol(indCol, val) &&
		m.pushGrid(indRow, indCol, val) {
		return true
	} else {
		return false
	}
}

func (m *Su) pushRow(indRow int, val byte) bool {
	return m.Rows[indRow].Push(val)
}

func (m *Su) pushCol(indCol int, val byte) bool {
	return m.Cols[indCol].Push(val)
}

func (m *Su) pushGrid(indRow, indCol int, val byte) bool {
	ind := (indRow%3)*3 + (indCol % 3)
	return m.Grid[ind].Push(val)
}

func main() {
	nums := [][]byte{{}, {}}
	k := isValidSudoku(nums)
	fmt.Printf("k:%+v\n", k)
}

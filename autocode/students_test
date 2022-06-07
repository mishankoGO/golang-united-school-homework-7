package coverage

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
var bDayMock time.Time = time.Now()

func TestLen(t *testing.T) {
	tData := []struct {
		p        People
		Expected int
	}{
		{p: People{
			Person{"a",
				"b",
				bDayMock,
			},
			Person{"c",
				"d",
				bDayMock,
			},
			Person{"e",
				"f",
				bDayMock,
			},
		}, Expected: 3},
		{p: People{
			Person{"g",
				"h",
				bDayMock,
			},
			Person{"i",
				"j",
				bDayMock,
			},
		}, Expected: 2},
		{p: People{},
			Expected: 0},
	}
	for k, v := range tData {
		got := v.p.Len()
		if got != v.Expected {
			t.Errorf("[%d] expected: %d, got %d", k, v.Expected, got)
		}
	}
}

func TestSwap(t *testing.T) {
	tData := struct {
		p        People
		Expected People
	}{
		p: People{
			Person{"a",
				"b",
				bDayMock,
			},
			Person{"c",
				"d",
				bDayMock,
			},
		},
		Expected: People{
			Person{"c",
				"d",
				bDayMock,
			},
			Person{"a",
				"b",
				bDayMock,
			},
		},
	}
	tData.p.Swap(0, 1)
	if tData.p[0] != tData.Expected[0] || tData.p[1] != tData.Expected[1] {
		t.Errorf("[%v] expected: %v, got %v", tData, tData.Expected, tData.p)
	}
}

func TestLess(t *testing.T) {
	tData := []struct {
		p        People
		Expected People
	}{
		{p: People{
			Person{"a",
				"c",
				bDayMock,
			},
			Person{"a",
				"b",
				bDayMock,
			},
		},
			Expected: People{
				Person{"a",
					"b",
					bDayMock,
				},
				Person{"a",
					"c",
					bDayMock,
				},
			},
		},
		{p: People{
			Person{"b",
				"b",
				bDayMock,
			},
			Person{"a",
				"b",
				bDayMock,
			},
		},
			Expected: People{
				Person{"a",
					"b",
					bDayMock,
				},
				Person{"b",
					"b",
					bDayMock,
				},
			},
		},
		{p: People{
			Person{"b",
				"b",
				bDayMock,
			},
			Person{"b",
				"b",
				bDayMock.Add(time.Hour * 10),
			},
		},
			Expected: People{
				Person{"b",
					"b",
					bDayMock.Add(time.Hour * 10),
				},
				Person{"b",
					"b",
					bDayMock,
				},
			},
		},
	}
	for k, v := range tData {
		sort.Slice(v.p, v.p.Less)
		if v.p[0] != v.Expected[0] || v.p[1] != v.Expected[1] {
			t.Errorf("[%d] expected: %v, got %v", k, v.p, v.Expected)
		}
	}
}

///////////////////////////////////////////////////////////////////////////////////////

func sliceMatch(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestNew(t *testing.T) {

	tData := []struct {
		str      string
		Expected *Matrix
		err      error
	}{
		{
			str:      "1 2\n3 4\n5 6",
			Expected: &Matrix{3, 2, []int{1, 2, 3, 4, 5, 6}},
			err:      nil,
		},
		{
			str:      "1 2\n3 4\n5",
			Expected: nil,
			err:      fmt.Errorf("Rows need to be the same length"),
		},
		{
			str:      "1 2\n3 4\n5 a",
			Expected: nil,
			err:      strconv.ErrSyntax,
		},
		{
			str:      "4 2\n3 4",
			Expected: &Matrix{2, 2, []int{4, 2, 3, 4}},
			err:      nil,
		},
	}
	for k, v := range tData {
		got, err := New(v.str)
		if err != nil && !errors.As(err, &v.err) {
			t.Errorf("[%d] error happend while not expected: %s", k, err.Error())
		}
		if got != nil && (got.rows != v.Expected.rows || got.cols != v.Expected.cols || !sliceMatch(got.data, v.Expected.data)) {
			t.Errorf("[%d] expected: %v, got %v", k, v.Expected, got)
		}
	}
}

func TestRows(t *testing.T) {
	tData := []struct {
		matrix   Matrix
		Expected [][]int
	}{
		{
			Matrix{3, 2, []int{1, 2, 3, 4, 5, 6}},
			[][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			Matrix{2, 1, []int{1, 2}},
			[][]int{{1}, {2}},
		},
	}
	for k, v := range tData {
		got := v.matrix.Rows()
		for i, elem := range got {
			if !sliceMatch(elem, v.Expected[i]) {
				t.Errorf("[%d] expected: %v, got %v", k, v.Expected, got)
				break
			}
		}
	}
}

func TestCols(t *testing.T) {
	tData := []struct {
		matrix   Matrix
		Expected [][]int
	}{
		{
			Matrix{3, 2, []int{1, 2, 3, 4, 5, 6}},
			[][]int{{1, 3, 5}, {2, 4, 6}},
		},
		{
			Matrix{2, 1, []int{1, 2}},
			[][]int{{1, 2}},
		},
	}
	for k, v := range tData {
		got := v.matrix.Cols()
		for i, elem := range got {
			if !sliceMatch(elem, v.Expected[i]) {
				t.Errorf("[%d] expected: %v, got %v", k, v.Expected, got)
				break
			}
		}
	}
}

func TestSet(t *testing.T) {
	tData := []struct {
		matrix   *Matrix
		row      int
		col      int
		Expected Matrix
	}{
		{
			&Matrix{3, 2, []int{1, 2, 3, 4, 5, 6}},
			0,
			0,
			Matrix{3, 2, []int{2, 2, 3, 4, 5, 6}},
		},
		{
			&Matrix{1, 2, []int{1, 2, 3, 4, 5, 6}},
			0,
			3,
			Matrix{1, 2, []int{1, 2, 3, 2, 5, 6}},
		},
		{
			&Matrix{1, 2, []int{1, 2, 3, 4, 5, 6}},
			3,
			0,
			Matrix{1, 2, []int{1, 2, 3, 2, 5, 6}},
		},
		{
			&Matrix{1, 1, []int{1}},
			0,
			0,
			Matrix{1, 1, []int{2}},
		},
	}
	for k, v := range tData {
		got := v.matrix.Set(v.row, v.col, 2)
		if !sliceMatch(v.matrix.data, v.Expected.data) && got {
			t.Errorf("[%d] expected: %v, got %v", k, v.Expected, v.matrix)
		}
	}
}

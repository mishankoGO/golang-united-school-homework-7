package coverage

import (
	"os"
	"sort"
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

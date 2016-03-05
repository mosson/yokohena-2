package tetoris

import "testing"

var tests = map[string]string{
	"ff-2f-23-f3-77-7f-3b": "1f-03-00-1c-0d-0f-06",
	"01":                      "00",
	"00":                      "00",
	"7a-4e":                   "0c-02",
	"56-b6":                   "08-14",
	"12-12-12":                "00-00-00",
	"de-ff-7b":                "0a-0f-05",
	"95-be-d0":                "05-1e-20",
	"7c-b0-bb":                "1c-20-2b",
	"7a-b6-31-6a":             "3a-56-11-2a",
	"32-0e-23-82":             "18-06-11-40",
	"ff-7f-bf-df-ef":          "0f-07-0b-0d-0e",
	"75-df-dc-6e-42":          "35-5f-5c-2e-02",
	"62-51-ef-c7-f8":          "22-11-6f-47-78",
	"0c-47-8e-dd-5d-17":       "04-23-46-6d-2d-0b",
	"aa-58-5b-6d-9f-1f":       "52-28-2b-35-4f-0f",
	"ff-55-d5-75-5d-57":       "0f-00-08-04-02-01",
	"fe-fd-fb-f7-ef-df-bf":    "7e-7d-7b-77-6f-5f-3f",
	"fd-fb-f7-ef-df-bf-7f":    "7e-7d-7b-77-6f-5f-3f",
	"d9-15-b5-d7-1b-9f-de":    "69-05-55-67-0b-4f-6e",
	"38-15-fd-50-10-96-ba":    "18-05-7d-20-00-46-5a",
	"fe-fd-fb-f7-ef-df-bf-7f": "fe-fd-fb-f7-ef-df-bf-7f",
}

func TestDescription(t *testing.T) {
	f := New("ff-2f-23-f3-77-7f-3b")

	f.Description()
}

func TestSolve(t *testing.T) {
	for k, v := range tests {
		f := New(k)
		s := f.solve()
		if s != v {
			t.Errorf("%v: expected %v, actual %v", k, v, s)
		}
	}
}

func TestAll(t *testing.T) {
	n := []uint{0, 1, 1, 1}
	if all(n) == true {
		t.Errorf("expected false: %v", n)
	}

	n = []uint{1, 1, 1, 1}
	if all(n) == false {
		t.Errorf("expected true: %v", n)
	}
}

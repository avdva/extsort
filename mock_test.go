package extsort

import (
	"context"
	"testing"
)

func Test50StringMock(t *testing.T) {
	a := makeTestStringArray(50)
	if IsStringsSorted(a) {
		t.Error("sorted before starting")
	}

	err := sortStringForTestMock(a)
	if err != nil {
		t.Fatalf("sort: %v", err)
	}

	if !IsStringsSorted(a) {
		t.Error("not sorted")
	}
}

func TestStringSmokeMock(t *testing.T) {
	a := make([]string, 3)
	a[0] = "banana"
	a[1] = "orange"
	a[2] = "apple"

	err := sortStringForTestMock(a)
	if err != nil {
		t.Fatalf("sort: %v", err)
	}

	if !IsStringsSorted(a) {
		t.Error("not sorted")
	}
}

func Test1KStringsMock(t *testing.T) {
	a := makeTestStringArray(1024)

	err := sortStringForTestMock(a)
	if err != nil {
		t.Fatalf("sort: %v", err)
	}
	if !IsStringsSorted(a) {
		t.Error("not sorted")
	}
}

func TestRandom1MStringMock(t *testing.T) {
	size := 1024 * 1024

	a := makeRandomStringArray(size)

	err := sortStringForTestMock(a)
	if err != nil {
		t.Fatalf("sort: %v", err)
	}
	if !IsStringsSorted(a) {
		t.Error("not sorted")
	}
}

func sortStringForTestMock(inputData []string) error {
	// make array of all data in chan
	inputChan := make(chan string, 2)
	go func() {
		for _, d := range inputData {
			inputChan <- d
		}
		close(inputChan)
	}()
	sort := StringsMockContext(context.Background(), inputChan, nil)
	outChan, errChan := sort.Sort()
	i := 0
	for {
		select {
		case err := <-errChan:
			return err
		case rec, more := <-outChan:
			if !more {
				return nil
			}
			inputData[i] = rec
			i++
		}
	}
}
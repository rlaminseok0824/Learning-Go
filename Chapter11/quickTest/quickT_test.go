package quickTest

import (
	"testing"
	"testing/quick" //블랙박스 테스팅을 위한 패키지 제공
)

var N = 1000000

func TestWithItself(t *testing.T) {
	condition := func(a, b Point2D) bool {
		return Add(a, b) == Add(b, a)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: N}) //Check 후 config Maxcount : N 랜덤 인수 집어넣음
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestThree(t *testing.T) {
	condition := func(a, b, c Point2D) bool {
		return Add(Add(a, b), c) == Add(a, b)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

package go_cond_broadcast

import "testing"

func Test_Solution(t *testing.T) {

	expected := 10
	result := solution(10)

	if expected != result {
		t.Fatalf("expected %d got %d", expected, result)
	}
}

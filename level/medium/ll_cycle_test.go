package medium

import (
	"testing"

	. "github.com/DeepikaAzad/leetCode/linked-list"
	utils "github.com/DeepikaAzad/leetCode/util"
	is "gotest.tools/v3/assert/cmp"
)

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		gotList []int
		want    bool
	}{
		{
			gotList: []int{3, 2, 0, -4},
			want:    true,
		},
	}

	for _, testCase := range testCases {
		actual := hasCycle(FromSlice(testCase.gotList))

		utils.Checkf(t, is.DeepEqual(actual, testCase.want), testCase)
	}
}

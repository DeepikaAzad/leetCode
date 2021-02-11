package easy

import (
	"testing"

	utils "github.com/DeepikaAzad/leetCode/util"
	is "gotest.tools/v3/assert/cmp"
)

func TestContainDuplicate(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    bool
	}{
		{
			gotNums: []int{2, 7, 11, 12},
			want:    false,
		},
		{
			gotNums: []int{3, 3},
			want:    true,
		},
	}

	for _, testCase := range testCases {
		actual := containDuplicate(testCase.gotNums)

		utils.Checkf(t, is.DeepEqual(actual, testCase.want), testCase)
	}
}

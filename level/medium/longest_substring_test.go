package medium

import (
	"testing"

	utils "github.com/DeepikaAzad/leetCode/util"
	is "gotest.tools/v3/assert/cmp"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		gotList string
		want    int
	}{
		{
			gotList: "aabbccef",
			want:    3,
		},
		{
			gotList: "dvdf",
			want:    3,
		},
		{
			gotList: "pwwkew",
			want:    3,
		},
		{

			gotList: "",
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := lengthOfLongestSubstring(testCase.gotList)

		utils.Checkf(t, is.DeepEqual(actual, testCase.want), testCase)
	}
}

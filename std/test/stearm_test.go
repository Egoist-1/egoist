package test

import (
	"testing"
	// "github.com/stretchr/testify/assert"
)
func Test_V1(t *testing.T){
	testCases := []struct{
		input          string
		expectedOutput int
	  }{
		{
		  input:          "aabbcc",
		  expectedOutput: 3,
		},
		{
		  input:          "abcdefg",
		  expectedOutput: 1,
		},
	  }
	  for _, _ = range testCases {
		// output := getRepetitions(tc.input)
		// assert.Equal
		// assert.Equal(t, tc.expectedOutput, output)
	  }
}
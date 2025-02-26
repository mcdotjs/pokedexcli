package main
import (
  "testing"
)
// just must end _test.go

func TestCleanInput(t *testing.T){
cases := []struct{
  input string
  expected []string
}{
  {
    input:" hello world  ",
    expected: []string{"hello", "world"},
  },
  {
    input:" oSrat them  ",
    expected: []string{"osrat", "them"},
  },
	{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
}

for _, c := range cases {
  actual := cleanInput(c.input)
  if len(actual) != len(c.expected){
    t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
    continue
  }
  for i, _ := range actual{
    word := actual[i]
    expectedWord := c.expected[i]
    if word != expectedWord {
      t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
   }
  }
}
}



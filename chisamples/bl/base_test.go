package bl

import "testing"

func TestTokenizeText(t *testing.T) {
	testInput := defaultInput()
	res := TokenizeText(testInput)
	if len(res) != 19 {
		t.Error("expected token size to be of 19 but found ", len(res))
	}

	res = TokenizeText("")
	if res != nil {
		t.Error("expected nil for empty text, but found ", res)
	}
}

func TestWordCount(t *testing.T) {
	testInput := defaultInput()
	res := WordCount(testInput)
	if res == nil {
		t.Error("expected map to non nil")
	}

	if count, found := res["of"]; !found || count != 2 {
		t.Error("expected count to be 2 but found ", count)
	}

	testInput = ""
	res = WordCount(testInput)
	if res != nil {
		t.Error("expected map to be nil for empty text")
	}
}

func defaultInput() string {
	return "Virat Kohli's men have found a bit of time to recover ahead of their match vs Afghanistan on Saturday"
}

package main

import "testing"

var data = "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"

func TestGroup(t *testing.T) {
	groups := Group(data)

	if len(groups) != 5 {
		t.Errorf("expected %d, got %d groups", 5, len(groups))
	}
}

func TestCountVotes(t *testing.T) {
	groups := Group(data)
	votes := CountVotes(groups)

	if votes != 11 {
		t.Errorf("expected %d, got %d votes", 11, votes)
	}
}

func TestCountAllYesVotes(t *testing.T) {
	groups := Group(data)
	votes := CountAllYesVotes(groups)

	if votes != 6 {
		t.Errorf("expected %d, got %d votes that were all yes for a group", 6, votes)
	}
}

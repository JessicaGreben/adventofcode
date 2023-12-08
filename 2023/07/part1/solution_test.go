package main

import "testing"

func TestSortHand(t *testing.T) {
	hands := []*hand{
		{
			cards: []*card{
				{"K", kindCard}, {"T", tenCard}, {"3", threeCard}, {"3", threeCard}, {"2", twoCard},
			},
			cardsStr: "32T3K",
			strength: onePair,
		},
		{
			cards: []*card{
				{"A", aceCard}, {"Q", queenCard}, {"Q", queenCard}, {"Q", queenCard}, {"J", jackCard},
			},
			cardsStr: "QQQJA",
			strength: threeKind,
		},
		{
			cards: []*card{
				{"J", jackCard}, {"T", tenCard}, {"5", fiveCard}, {"5", fiveCard}, {"5", fiveCard},
			},
			cardsStr: "T55J5",
			strength: threeKind,
		},
		{
			cards: []*card{
				{"K", kindCard}, {"K", kindCard}, {"7", sevenCard}, {"7", sevenCard}, {"6", sixCard},
			},
			cardsStr: "KK677",
			strength: twoPair,
		},
		{
			cards: []*card{
				{"K", kindCard}, {"J", jackCard}, {"J", jackCard}, {"T", tenCard}, {"T", tenCard},
			},
			cardsStr: "KTJJT",
			strength: twoPair,
		},
	}
	reverseSortHands(hands)
	want := []string{"32T3K", "KTJJT", "KK677", "T55J5", "QQQJA"}
	for i := range want {
		if want, got := want[i], hands[i].cardsStr; want != got {
			t.Errorf("want=%v, got=%v", want, got)
		}
	}

}
func TestSolutionPart1(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		out, err := solution("../input_test.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(6440), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("input", func(t *testing.T) {
		out, err := solution("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if want, got := int64(251121738), out; want != got {
			t.Errorf("want=%d, got=%d", want, got)
		}
	})

	t.Run("parse input", func(t *testing.T) {
		h, err := ForEachLine("../input_test.txt", processLine)
		if err != nil {
			t.Fatal(err)
		}
		wantHand := []string{"32T3K", "T55J5", "KK677", "KTJJT", "QQQJA"}
		if want, got := len(wantHand), len(h); want != got {
			t.Fatalf("want=%d, got=%d", want, got)
		}
	})
}

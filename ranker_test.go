package ranker_test

import "testing"
import "github.com/dlapets/ranker"

func TestRanker_Top(t *testing.T) {

	input := []string{
		"a", "a", "a", "a", "a",
		"z", "z",
		"b", "b", "b", "b",
		"c", "c", "c",
		"d", "d",
		"e",
		"f",
		"g",
		"h",
		"ze",
		"zf",
		"zg",
		"zh",
		"zze",
		"zzf",
		"zzg",
		"zzh",
	}

	r := ranker.NewRanker(8)

	// Call r.Add n=(100 * input) times:
	for i := 0; i < 100; i++ {
		for _, s := range input {
			r.Add(s)
			if l := len(r.Occurrences); l > 8 {
				t.Errorf("%d is too big", l)
				t.FailNow()
			}
		}
	}

	expected := []string{"a", "b", "c", "d"}

	// Get the top 4
	top := r.Top(4)

	if len(top) != 4 {
		t.Errorf("%s is too short", top)
		t.FailNow()
	}

	for i, s := range expected {
		if top[i] != s {
			t.Errorf("Item at index %d should be %s but is %s", i, s, top[i])
		}
	}
}

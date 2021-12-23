package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	t.Run("with general types", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				Name: "Struct with one string field",
				Input: struct {
					Name string
				}{"Chris"},
				ExpectedCalls: []string{"Chris"},
			}, {
				Name: "Struct with two string fields",
				Input: struct {
					Name string
					City string
				}{"Chris", "London"},
				ExpectedCalls: []string{"Chris", "London"},
			}, {
				Name: "Struct with non string field",
				Input: struct {
					Name string
					Age  int
				}{"Chris", 33},
				ExpectedCalls: []string{"Chris"},
			}, {
				Name:          "Nested fields",
				Input:         Person{"Chris", Profile{33, "London"}},
				ExpectedCalls: []string{"Chris", "London"},
			}, {
				Name:          "Pointers to things",
				Input:         &Person{"Chris", Profile{33, "London"}},
				ExpectedCalls: []string{"Chris", "London"},
			}, {
				Name: "Slices",
				Input: []Profile{
					{33, "London"},
					{34, "Reykjavik"},
				},
				ExpectedCalls: []string{"London", "Reykjavik"},
			}, {
				Name: "Arrays",
				Input: [2]Profile{
					{33, "London"},
					{34, "Reykjavik"},
				},
				ExpectedCalls: []string{"London", "Reykjavik"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(test.ExpectedCalls, got) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}
	})
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(s string) {
			got = append(got, s)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, got []string, want string) {
	t.Helper()
	contains := false
	for _, s := range got {
		if s == want {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", got, want)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

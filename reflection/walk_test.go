package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "Paris"},
			[]string{"Chris", "Paris"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 32},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{"Chris", Profile{32, "Paris"}},
			[]string{"Chris", "Paris"},
		},
		{
			"slices",
			[]Profile{
				{23, "London"},
				{36, "Madrid"},
			},
			[]string{"London", "Madrid"},
		},
		{
			"array",
			[2]Profile{
				{23, "London"},
				{36, "Madrid"},
			},
			[]string{"London", "Madrid"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Bhee",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Bhee")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{34, "Belo Jardim"}
			aChannel <- Profile{28, "New Mexico"}

			close(aChannel)
		}()
  
		var got []string
		want := []string{
			"Belo Jardim",
			"New Mexico",
		}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got: %v, want: %v", got, want)
    }
	})

	t.Run("with functions", func(t *testing.T) {
    aFunction := func() (Profile, Profile) {
			return Profile{34, "Belo Jardim"}, Profile{28, "New Mexico"}
    }

		var got []string
		want := []string{
			"Belo Jardim",
			"New Mexico",
		}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got: %v, want: %v", got, want)
    }
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, value := range haystack {
		if value == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contains %q but it didn't", haystack, needle)
	}
}

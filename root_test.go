package backedEnum

import (
	"fmt"
	"strings"
	"testing"

	"github.com/samber/lo"
)

func TestBackedEnum(testUtils *testing.T) {

	testUtils.Run("works", func(t *testing.T) {

		enum := BackedEnum[int, map[string]int]{
			structure: map[string]int{
				"FOO": 1,
				"BAR": 2,
				"BAZ": 8,
			},
		}

		result := enum.Structure()["FOO"]

		expectation := 1

		if result != expectation {
			t.Fatalf("This %v is not %d ", result, expectation)
		}

	})

	enum := BackedEnum[int, map[string]int]{
		structure: map[string]int{
			"FOO": 1,
			"BAR": 2,
			"BAZ": 8,
		},
	}

	testUtils.Run(
		"all values are returned when Values() is called",
		func(t *testing.T) {

			expectedValues := []int{1, 2, 8}

			values := enum.Values()

			result := lo.EveryBy(values, func(item int) bool {

				return lo.Contains(expectedValues, item)

			})

			createCommaSeparatedStringFromIntArray := func(array []int) string {

				return strings.Join(lo.Map(
					array,
					func(item int, index int) string {
						return fmt.Sprint(item)
					}), ",")

			}

			if !result {

				t.Fatalf(
					"The values from the enum are supposed to have %s they have %s",
					createCommaSeparatedStringFromIntArray(expectedValues),
					createCommaSeparatedStringFromIntArray(values),
				)

			}

		})

	testUtils.Run(
		"An error is returned when the input passed to Parse() is an invalid value",
		func(t *testing.T) {

			input := 9
			result := enum.Parse(input)

			if result == nil {
				t.Fatalf("This input %d is improper input an error was supposed to be returned", input)
			}

		})

	testUtils.Run(
		"The bool false is returned when the input passed to Validate() is an invalid value",
		func(t *testing.T) {

			input := 5
			result := enum.Validate(input)

			if result == true {
				t.Fatalf("This input %d is improper input false is supposed to be returned", input)
			}

		})

}

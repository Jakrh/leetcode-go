package solution

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		d3 := i%3 == 0
		d5 := i%5 == 0

		if d3 && d5 {
			result = append(result, "FizzBuzz")
		} else if d3 {
			result = append(result, "Fizz")
		} else if d5 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

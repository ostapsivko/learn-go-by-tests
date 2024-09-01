package arrays

func Sum(numbers []int) int {
	sum := func(x, y int) int {
		return x + y
	}
	return Reduce(numbers, sum, 0)
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sumTail = func(acc, numbers []int) []int {
		if len(numbers) == 0 {
			return append(acc, 0)
		} else {
			tail := numbers[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[T any](values []T, f func(T, T) T, zero T) T {
	var sum = zero

	for _, value := range values {
		sum = f(sum, value)
	}
	return sum
}

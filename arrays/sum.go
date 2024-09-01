package arrays

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(current float64, t Transaction) float64 {
		if t.To == name {
			current += t.Sum
		}

		if t.From == name {
			current -= t.Sum
		}
		return current
	}

	return Reduce(transactions, adjustBalance, 0.0)
}

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

func Reduce[T, Y any](values []T, f func(Y, T) Y, zero Y) Y {
	var sum = zero

	for _, value := range values {
		sum = f(sum, value)
	}
	return sum
}

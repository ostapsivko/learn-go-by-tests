package arrays

type Transaction struct {
	From, To string
	Sum      float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, amount float64) Transaction {
	return Transaction{from.Name, to.Name, amount}
}

func NewBalanceFor(acc Account, transactions []Transaction) Account {
	calculate := func(a Account, t Transaction) Account {
		if t.From == a.Name {
			a.Balance -= t.Sum
		}

		if t.To == a.Name {
			a.Balance += t.Sum
		}

		return a
	}

	return Reduce(transactions, calculate, acc)
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

package add

func Sum[T ~int](numbers []T) T {
	var sum T = 0
	for _, valueInt := range numbers {
		sum += valueInt
	}
	return sum
}

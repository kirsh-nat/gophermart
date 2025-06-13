package orderservices

func checkBalance(sum, accural float32) error {
	if (sum > accural) || (sum < 0) {
		return NewInsufficientFundsError("Insufficient funds", nil)
	}

	return nil
}

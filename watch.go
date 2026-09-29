package main

// Delta reports how many transactions were sent between two nonce readings.
func Delta(prev, now uint64, seen bool) (uint64, bool) {
	if !seen || now <= prev {
		return 0, false
	}
	return now - prev, true
}

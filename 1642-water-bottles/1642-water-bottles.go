func numWaterBottles(numBottles int, numExchange int) int {
	totDrank := numBottles

	for numBottles >= numExchange {
		totDrank += numBottles / numExchange
		numBottles = (numBottles / numExchange) + (numBottles % numExchange)
	}
	return totDrank
}
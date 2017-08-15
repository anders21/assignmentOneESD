package main

/* A typical style for me */
func sqaure(aNumber int) int {
	result := aNumber * aNumber
	return result
}

/* Go's way of developing */
func goSquare(aNumber *int) {
	*aNumber = *aNumber * *aNumber
}
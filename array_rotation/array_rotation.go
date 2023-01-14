package arrayrotation

func ArrayRotation(arr []int, n int) []int {

	moveToBack := arr[0:n]
	moveToFront := arr[n:]

	rotated := append(moveToFront, moveToBack...)

	return rotated
}

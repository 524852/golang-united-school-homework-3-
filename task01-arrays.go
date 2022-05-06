package homework

/*Task 1: Arrays
Implement function that returns an average value of array (sum / N)
input -> [1,2,3,4,5,6]
output -> 3.5
*/
func average(input [15]float32) (result float32) {
	for i := 0; i < len(input); i++ {
		result += input[i]
	}

	return result / 15
}

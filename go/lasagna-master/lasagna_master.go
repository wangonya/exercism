package lasagna

func PreparationTime(layers []string, prep_time_per_layer int) int {
	if prep_time_per_layer == 0 {
		prep_time_per_layer = 2
	}
	return len(layers) * prep_time_per_layer
}

func Quantities(layers []string) (int, float64) {
	noodle_layers := 0
	sauce_layers := 0
	for _, layer := range layers {
		if layer == "noodles" {
			noodle_layers += 1
		} else if layer == "sauce" {
			sauce_layers += 1
		}
	}
	return noodle_layers * 50, float64(sauce_layers) * 0.2
}
func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var newQuantities []float64

	for _, quantity := range quantities {
		newQuantities = append(newQuantities, (quantity/2)*float64(portions))
	}

	return newQuantities
}

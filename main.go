package main

import "fmt"

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(0),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(2),
		Sodium:              SodiumMilligram(400),
		Fruits:              FruitsPercent(60),
		Fibre:               FibreGram(40),
		Protein:             ProteinGram(2),
	}, Water)

	fmt.Printf("Nutrional Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s \n", ns.GetNutriScore())
}

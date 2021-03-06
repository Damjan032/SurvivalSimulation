package services

import (
	"../models"
	"../utils"
	"fmt"
	"math/rand"
)

func Merge(population []*models.Individual, foodSources []*models.FoodSource) {
	var populationFreeArray = utils.MakeArray(len(population))   //Lista id-ova
	var foodSourcesFreeArray = utils.MakeArray(len(foodSources)) //Lista id-ova
	// prolazim kroz listu dok svi ne nadju izvori za hranu ili dok se ne popune svi izvori hrane
	for len(populationFreeArray) != 0 {
		// posto lista sadrzi sve id-jove uzmemo random iz nje
		if len(foodSourcesFreeArray) == 0 {
			break
		}
		var foodRandomId = rand.Intn(len(foodSourcesFreeArray))
		var foodSourceId = foodSourcesFreeArray[foodRandomId] //random id samo
		var individualRandomId = rand.Intn(len(populationFreeArray))
		var plebId = populationFreeArray[individualRandomId]
		var pleb = population[plebId]
		if foodSources[foodSourceId].Occupied == 0 {
			foodSources[foodSourceId].OccupiedBy = append(foodSources[foodSourceId].OccupiedBy, pleb)
			//pleb.resources = append(pleb.resources, foodSources[foodSourceId].resources[0])
			foodSources[foodSourceId].Occupied++
		} else {
			foodSources[foodSourceId].OccupiedBy = append(foodSources[foodSourceId].OccupiedBy, pleb)
			//pleb.resources = append(pleb.resources, foodSources[foodSourceId].resources[1])
			foodSources[foodSourceId].Occupied++
			foodSourcesFreeArray = append(foodSourcesFreeArray[:foodRandomId], foodSourcesFreeArray[foodRandomId+1:]...)
		}
		populationFreeArray = append(populationFreeArray[:individualRandomId], populationFreeArray[individualRandomId+1:]...)
	}
}

func merge2(population []*models.Individual, foodSources []*models.FoodSource) {
	var populationFreeArray = utils.MakeArray(len(population))   //Lista id-ova
	var foodSourcesFreeArray = utils.MakeArray(len(foodSources)) //Lista id-ova
	fmt.Printf("%v\n", populationFreeArray)
	fmt.Printf("%v\n", foodSourcesFreeArray)
	for _, pleb := range population {
		// posto lista sadrzi sve id-jove uzmemo random iz nje
		if len(foodSourcesFreeArray) == 0 {
			break
		}
		var randomFoodId = rand.Intn(len(foodSourcesFreeArray))
		var foodSourceId = foodSourcesFreeArray[randomFoodId]
		if foodSources[foodSourceId].Occupied == 0 {
			pleb.Resources = append(pleb.Resources, foodSources[foodSourceId].Resources[0])
			foodSources[foodSourceId].Occupied++
		} else {
			pleb.Resources = append(pleb.Resources, foodSources[foodSourceId].Resources[1])
			foodSources[foodSourceId].Occupied++
			foodSourcesFreeArray = append(foodSourcesFreeArray[:randomFoodId], foodSourcesFreeArray[randomFoodId+1:]...)
		}
		//pleb.resources = append(pleb.resources, foodSources[foodSourceId])
		/*for !isFreeFoodSource(foodSources[foodSourceId]){
			fmt.Println("USO MICO")
			foodSourceId = rand.Intn(len(foodSources))
		}*/
		//fmt.Print(rand.Intn(len(foodSources)), ",")
		//fmt.Println("My animal is:", pleb)
	}
}

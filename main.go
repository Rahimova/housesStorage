package main

import (
	"sort"
)

type house struct {
	id                 int64
	title              string
	price              int64
	numberOfRooms      int64
	numberOfFloors     int64
	location           Location
	landArea           int64
	squareArea         int64
	distanceFromCentre int64
	district           string
}

type Location struct {
	city     string
	district string
	street   string
}

func sortByPriceAscAndDesc(houses [] house) (resultForAsc, resultForDesc [] house) {
	resultForAsc = make([]house, len(houses))
	copy(resultForAsc, houses)
	resultForDesc = make([]house, len(houses))
	copy(resultForDesc, houses)
	sort.Slice(resultForAsc, func(i, j int) bool {
		return resultForAsc[i].price < resultForAsc [j].price
	})
	sort.Slice(resultForDesc, func(i, j int) bool {
		return resultForDesc[i].price > resultForDesc [j].price
	})
	return
}

func sortByDistanceFromCentreAscAndDesc(houses [] house) (resultForAsc, resultForDesc []house) {
	resultForAsc = make([]house, len(houses))
	copy(resultForAsc, houses)
	sort.Slice(resultForAsc, func(i, j int) bool {
		return resultForAsc[i].distanceFromCentre > resultForAsc [j].distanceFromCentre
	})

	resultForDesc = make([]house, len(houses))
	copy(resultForDesc, houses)
	sort.Slice(resultForDesc, func(i, j int) bool {

		return resultForDesc[i].distanceFromCentre < resultForDesc [j].distanceFromCentre
	})
	return
}

func searchByMaxPrice(houses [] house, maxPrice int64) []house {

	result := make([]house, 0)
	for _, house := range houses {
		if house.price <= maxPrice {

			result = append(result, house)
		}
	}
	return result
}

func searchByMaxAndMinPrice(houses [] house, maxPrice int64, minPrice int64, ) []house {

	result := make([]house, 0)
	for _, house := range houses {
		if house.price < minPrice {
			continue
		}
		if house.price > maxPrice {
			continue
		}
		result = append(result, house)
	}
	return result
}

func searchByDistrict(houses [] house, district string) []house {

	result := make([]house, 0)
	for _, house := range houses {
		if house.location.district == district {
			result = append(result, house)
		}
	}
	return result
}

func searchByDistricts(houses [] house, districts []string) []house {

	result := make([]house, 0)
	for _, house := range houses {
		for _, district := range districts {
			if house.location.district == district {
				result = append(result, house)
			}
		}
	}
	return result
}

func main() {

}

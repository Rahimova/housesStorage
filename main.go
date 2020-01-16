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

func sortByPriceAsc(houses [] house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		// TODO:
		return result[i].price < result [j].price
	})
	return result
}

func sortByPriceDesc(houses [] house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		// TODO:
		return result[i].price > result [j].price
	})
	return result
}
func sortByDistanceFromCentreAsc(houses [] house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {

		return result[i].distanceFromCentre > result [j].distanceFromCentre
	})
	return result
}

func sortByDistanceFromCentreDesc(houses [] house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {

		return result[i].distanceFromCentre < result [j].distanceFromCentre
	})
	return result
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

func searchByDistricts(houses [] house, districts []string ) []house {

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
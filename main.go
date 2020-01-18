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

func sortBy(houses []house, less func(a, b house) bool) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return less(houses[i], houses[j])
	})
	return result
}

func sortByPriceAsc(houses [] house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.price > b.price
	})
}

func sortByPriceDesc(houses [] house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.price > b.price
	})
}

func sortByDistanceFromCentreAsc(houses [] house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.distanceFromCentre < b.distanceFromCentre
	})
}

func sortByDistanceFromCentreDesc(houses [] house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.distanceFromCentre > b.distanceFromCentre
	})
}

func searchBy(houses []house, predicate func(houses house) bool) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if predicate(house) {
			result = append(result, house)
		}
	}
	return result
}

func searchByMaxPrice(houses []house, limit int64) []house {
	return searchBy(houses, func(a house) bool {
		return a.price <= limit
	})
}

func searchByMaxPriceAndMinPrice(houses []house, maxPrice, minPrice int64) []house {
	return searchBy(houses, func(houses house) bool {
		if houses.price < minPrice {
			return false
		}
		if houses.price > maxPrice {
			return false
		}
		return true
	})
}


func searchByDistrict(houses [] house, district string) []house {
	return searchBy(houses, func(house house) bool {
		return house.location.district == district
	})
}

func searchByDistricts(houses [] house, district string) []house {
	return searchBy(houses, func(house house) bool {
		return house.location.district == district
	})
}



func main() {

}
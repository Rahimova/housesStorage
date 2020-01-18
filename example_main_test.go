package main

import "fmt"

var (
	houses []house = []house{
		{   id:                    1,
			title:              "Дом1",
			price:              500_000,
			numberOfRooms:      5,
			numberOfFloors:     2,
			location:           Location{city: "Dushanbe", district: "Shohmansur", street: "Ayni street"},
			landArea:           24,
			squareArea:         87,
			distanceFromCentre: 150,
		},

		{   id:                  2,
			title:              "Дом2",
			price:              600_000,
			numberOfRooms:      5,
			numberOfFloors:     2,
			location:           Location{city: "Dushanbe", district: "Shohmansur", street: "Ayni street"},
			landArea:           24,
			squareArea:         87,
			distanceFromCentre: 50,
		},

		{   id:                  3,
			title:              "Дом3",
			price:              1600_000,
			numberOfRooms:      15,
			numberOfFloors:     2,
			location:           Location{city: "Dushanbe", district: "Shohmansur", street: "Ayni street"},
			landArea:           22,
			squareArea:         87,
			distanceFromCentre: 5,
		},
	}
)

func ExampleSortByPriceAsc() {
	sortByPriceAsc := sortByPriceAsc(houses)
	fmt.Println(sortByPriceAsc)
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 }]

}

func ExampleSortByPriceDesc() {
	sortByPriceDesc := sortByPriceDesc(houses)
	fmt.Println(sortByPriceDesc)
	//Output: [{3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 }]
}

func ExampleSortByDistanceFromCentreAsc() {
	sortByDistanceFromCentreAsc := sortByDistanceFromCentreAsc(houses)
	fmt.Println(sortByDistanceFromCentreAsc)
	//Output: [{3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 }]
}

func ExampleSortByDistanceFromCentreDesc() {
	sortByDistanceFromCentreDesc := sortByDistanceFromCentreDesc(houses)
	fmt.Println(sortByDistanceFromCentreDesc)
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 }]
}
func ExampleSearchByMaxPrice_MultipleResults() {
	fmt.Println(searchByMaxPrice(houses, 1_000_000))
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 }]
}

func ExampleSearchByMaxPrice_NoResults() {
	fmt.Println(searchByMaxPrice(houses, 1_000))
	//Output: []
}

func ExampleSearchByMaxPrice_OneResults() {
	fmt.Println(searchByMaxPrice(houses, 500_000))
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 }]
}

func ExampleSearchByWantedPriceOneResult() {
	fmt.Println(searchByMaxPriceAndMinPrice(houses, 2_000_000, 1_000_000))
	//Output: [{3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 }]
}

func ExampleSearchByWantedPriceNoResult() {
	fmt.Println(searchByMaxPriceAndMinPrice(houses, 5_000_000, 2_000_000))
	//Output: []
}

func ExampleSearchByWantedPriceMultiResult() {
	fmt.Println(searchByMaxPrice(houses, 600_000))
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 }]
}

func ExampleSearchByDistrictMultiResult() {
	fmt.Println(searchByDistrict(houses, "Shohmansur"))
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 }]
}

func ExampleSearchByDistrictNoResult() {
	fmt.Println(searchByDistrict(houses, "Somoni"))
	//Output: []
}

func ExampleSearchByDistrictsMultiResult() {
	fmt.Println(searchByDistricts(houses, []string{"Shohmansur"}))
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 }]

}
func ExampleSearchByDistrictsNoResult() {
	fmt.Println(searchByDistricts(houses, []string{"Somoni"}))
	//Output: []
}

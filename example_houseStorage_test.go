package main

import "fmt"

var (
	houses []house = []house{
		{id: 1,
			title:              "Дом1",
			price:              500_000,
			numberOfRooms:      5,
			numberOfFloors:     2,
			location:           Location{city: "Dushanbe", district: "Shohmansur", street: "Ayni street"},
			landArea:           24,
			squareArea:         87,
			distanceFromCentre: 150,
		},

		{id: 2,
			title:              "Дом2",
			price:              600_000,
			numberOfRooms:      5,
			numberOfFloors:     2,
			location:           Location{city: "Dushanbe", district: "Shohmansur", street: "Ayni street"},
			landArea:           24,
			squareArea:         87,
			distanceFromCentre: 50,
		},

		{id: 3,
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
	fmt.Println(sortByPriceAsc(houses))
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 }]
}

func ExampleSortByPriceDesc() {
	fmt.Println(sortByPriceDesc(houses))
	//Output:[{3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 }]
}

func ExampleSortByDistanceFromCentreAsc() {
	fmt.Println(sortByDistanceFromCentreAsc(houses))
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 }]
}

func ExampleSortByDistanceFromCentreDesc() {
	fmt.Println(sortByDistanceFromCentreDesc(houses))
	//Output:  [{3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 }]
}

func ExampleSearchByLimitPrice() {
	fmt.Println(searchByMaxPrice(houses, 1_000_000))
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 }]
}

func ExampleSearchByWantedPrice() {
	fmt.Println(searchByMaxAndMinPrice(houses, 2_000_000, 1_000_000))
	//Output: [{3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 }]
}

func ExampleSearchByDistrict() {
	fmt.Println(searchByDistrict(houses, "Shohmansur"))
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 }]
}

func ExampleSearchByDistricts() {
	fmt.Println(searchByDistricts(houses, []string{"Shohmansur"}))
	//Output: [{1 Дом1 500000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 150 } {2 Дом2 600000 5 2 {Dushanbe Shohmansur Ayni street} 24 87 50 } {3 Дом3 1600000 15 2 {Dushanbe Shohmansur Ayni street} 22 87 5 }]

}

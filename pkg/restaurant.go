package pkg

type Restaurant struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func MasterRestaurant() Restaurant{

	restaurant := Restaurant{
		Id: 3,
		Name: "サイゼリヤ",
	}
	return restaurant
}
package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	tacoma := truck{
		vehicle: vehicle{
			doors: 2,
			color: "maroon",
		},
		fourWheel: true,
	}
	camry := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "gold",
		},
		luxury: true,
	}

	fmt.Println(tacoma)
	fmt.Println(camry)

	fmt.Println(camry.doors, camry.color, camry.luxury)
	fmt.Println(tacoma.doors, tacoma.color, tacoma.fourWheel)
}

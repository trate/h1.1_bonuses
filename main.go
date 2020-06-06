package main

func main() {
	// in kopeykas
	purchase := 3333_33
	// regarding the rubles
	percent := 1
	limit := 100

	bonus := purchase / (percent * 100 * 100)
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
}

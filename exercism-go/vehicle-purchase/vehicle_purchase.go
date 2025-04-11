package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(opt1, opt2 string) string {
	suffix := " is clearly the better choice."
	if opt1 > opt2 {
		return opt2 + suffix
	}
	return opt1 + suffix
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3:
		return 0.8 * originalPrice
	case age < 10:
		return 0.7 * originalPrice
	default:
		return 0.5 * originalPrice
	}
}

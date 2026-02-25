package purchase

// LicenseRequiredVehicles represents vehicles that requires a license
var LicenseRequiredVehicles = []string{"car", "truck"}

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	for _, v := range LicenseRequiredVehicles {
        if kind == v {
            return true
        }
    }
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    better := option2
    // lexicographical order = compare alphabet left to right. first word with lower char wins. everything after's irrelevant.
	if option1 < option2 {
        better = option1
    }
    return better + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    calculatePrice := func(percentage float64) float64 {
        return originalPrice * percentage / 100
    }
	if age < 3 {
        return calculatePrice(80)
    } else if age >= 3 && age < 10 {
        return calculatePrice(70)
    }
    return calculatePrice(50)
}

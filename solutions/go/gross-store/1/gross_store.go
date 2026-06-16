package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 	6,
        "dozen": 	12,
        "small_gross": 	120,
        "gross": 	144,
        "great_gross": 	1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    quantity, exists := units[unit]
    if !exists {
        return false
    }

    bill[item] += quantity // if bill[item] doesn't exist, 0 + quantity gets stored
    return true
}
    
    // RemoveItem removes an item from customer bill.
    func RemoveItem(bill, units map[string]int, item, unit string) bool {
    	if _, exists := bill[item]; !exists {
            return false
        }
        quantity, exists := units[unit]
        newQty := bill[item] - quantity
        if !exists || newQty < 0 {
            return false
        }
    	if newQty == 0 {
            delete(bill, item)
            return true
        }
        bill[item] = newQty
        return true
    }

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]
    return qty, exists
}

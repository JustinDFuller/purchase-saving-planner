package planner

import "time"

type (
	// User is the top level struct for, you guessed it, user data.
	// The data is stored in an object database, so purchases
	// Are stored as a property on the user, rather than by association.
	User struct {
		Email         string     `json:"email"`
		Saved         int64      `json:"saved"`
		Contributions int64      `json:"contributions"`
		Frequency     string     `json:"frequency"`
		LastPaycheck  *time.Time `json:"lastPaycheck"`
		Purchases     []Purchase `json:"purchases"`
	}

	// Purchase is something a User wants to buy.
	// TODO: Link it more appropriately with Product.
	Purchase struct {
		Date    *time.Time `json:"date"`
		Product Product    `json:"product"`
	}

	// Product contains information about the thing a User wants to buy.
	Product struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int64  `json:"price"`
		URL         string `json:"url"`
		Image       string `json:"image"`
	}
)

// IsEmpty can be used to determine if a Product is empty.
func (p Product) IsEmpty() bool {
	return p == Product{}
}

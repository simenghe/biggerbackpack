package core

// Bag Type For Elites
type Bag struct {
	Title string  `json:"Title"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

// GetPackData grabs from mysql database
func GetPackData() []Bag {
	return []Bag{
		Bag{Title: "Jansport", Desc: "eLITE women", Price: 32.3},
		Bag{Title: "Swiss Gear", Desc: "eLITE man", Price: 12.3},
	}
}

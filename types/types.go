package types

type Port struct {
	Name        string    `json:"name"`
	Coordinates []float32 `json:"coordinates"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	// No example of alias and regions in the json file. 
	// I assume they are list of strings
	Alias    []string `json:"alias"`
	Regions  []string `json:"regions"`
	Timezone string   `json:"timezone"`
	Unlocs   []string `json:"unlocs"`
	Code     string   `json:"code"`
}

package structs

type Product struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	SubCategory string `json:"subCategory"`
	Price       string `json:"price"`
	DataType    string `json:"datatype"`
}

package flickr

const PhotoFIles = "photo*"

type Photo struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	DateTaken    string `json:"date_taken"`
	DateImported string `json:"date_imported"`
	Photopage    string `json:"photopage"`
	Original     string `json:"original"`
	Albums       []struct {
		Id    string `json:"id"`
		Url   string `json:"url"`
		Title string `json:"title"`
	} `json:"albums"`

	Geo struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
		Accuracy  string `json:"accuracy"`
	} `json:"geo"`
	FileName string `json:"-"`
}

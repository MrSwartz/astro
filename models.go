package astro

type Pictures struct {
	Pics []Picture `db:"pictures"`
}

type Picture struct {
	Copyright      string `json:"copyright" db:"copyright" `
	Explanation    string `json:"explanation" db:"explanation"`
	Hdurl          string `json:"hdurl" db:"hd_url"`
	MediaType      string `json:"media_type" db:"media_type"`
	ServiceVersion string `json:"service_version" db:"service_version"`
	Title          string `json:"title" db:"title"`
	Url            string `json:"url" db:"url"`
	BinaryPic      string `json:"binary_pic" db:"binary_pic"`
	PicOfTheDay    string `json:"date" db:"pic_of_the_day"`
	Stored         string `db:"stored"`
}

type Response struct {
	Status   int `json:"status"`
	Pictures []Picture
	Error
}

type Error struct {
	Message string `json:"message"`
}

type ResponsePair struct {
	Body  []byte
	Error error
}

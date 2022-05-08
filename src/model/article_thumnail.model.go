package model

type ArticleThumnail struct {
	Data ArticleThumnailData `json:"data"`
}

type ArticleThumnailData struct {
	Id         int64                         `json:"id"`
	Attributes ArticleThumnailDataAttributes `json:"attributes"`
}

type ArticleThumnailDataAttributes struct {
	Name              string                              `json:"name"`
	AlternativeText   string                              `json:"alternativeText"`
	Caption           string                              `json:"caption"`
	Width             int64                               `json:"width"`
	Height            int64                               `json:"height"`
	Hash              string                              `json:"hash"`
	Ext               string                              `json:"ext"`
	Mime              string                              `json:"mime"`
	Size              float64                             `json:"size"`
	Url               string                              `json:"url"`
	PreviewUrl        string                              `json:"previewUrl"`
	Provider          string                              `json:"provider"`
	Provider_metadata string                              `json:"provider_metadata"`
	CreatedAt         string                              `json:"createdAt"`
	UpdatedAt         string                              `json:"updatedAt"`
	Formats           ArticleThumnailDataAttributesFormat `json:"formats"`
}

type ArticleThumnailDataAttributesFormat struct {
	Large    ArticleThumnailDataAttributesFormatLarge     `json:"large"`
	Small    ArticleThumnailDataAttributesFormatSmall     `json:"small"`
	Medium   ArticleThumnailDataAttributesFormatMedium    `json:"medium"`
	Xsmall   ArticleThumnailDataAttributesFormatXsmall    `json:"xsmall"`
	Thumnail ArticleThumnailDataAttributesFormatThumbnail `json:"thumbnail"`
}

type ArticleThumnailDataAttributesFormatLarge struct {
	Ext    string  `json:"ext"`
	Url    string  `json:"url"`
	Hash   string  `json:"hash"`
	Mime   string  `json:"mime"`
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Size   float64 `json:"size"`
	Width  int64   `json:"width"`
	Height int64   `json:"height"`
}

type ArticleThumnailDataAttributesFormatSmall struct {
	Ext    string  `json:"ext"`
	Url    string  `json:"url"`
	Hash   string  `json:"hash"`
	Mime   string  `json:"mime"`
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Size   float64 `json:"size"`
	Width  int64   `json:"width"`
	Height int64   `json:"height"`
}

type ArticleThumnailDataAttributesFormatMedium struct {
	Ext    string  `json:"ext"`
	Url    string  `json:"url"`
	Hash   string  `json:"hash"`
	Mime   string  `json:"mime"`
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Size   float64 `json:"size"`
	Width  int64   `json:"width"`
	Height int64   `json:"height"`
}

type ArticleThumnailDataAttributesFormatXsmall struct {
	Ext    string  `json:"ext"`
	Url    string  `json:"url"`
	Hash   string  `json:"hash"`
	Mime   string  `json:"mime"`
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Size   float64 `json:"size"`
	Width  int64   `json:"width"`
	Height int64   `json:"height"`
}

type ArticleThumnailDataAttributesFormatThumbnail struct {
	Ext    string  `json:"ext"`
	Url    string  `json:"url"`
	Hash   string  `json:"hash"`
	Mime   string  `json:"mime"`
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Size   float64 `json:"size"`
	Width  int64   `json:"width"`
	Height int64   `json:"height"`
}

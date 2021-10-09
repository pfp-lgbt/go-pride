package pride

const (
	BaseURL           = "https://api.pfp.lgbt/v6"
	GetFlagsURL       = BaseURL + "/data/flag"
	GetFlagsByNameURL = BaseURL + "/data/flag/%s"
	SearchFlagsURL    = BaseURL + "/data/flag/search?q=%s&limit=%d"
	PutImageURL       = BaseURL + "/image"
	EditImageURL      = BaseURL + "/image/%s"
)

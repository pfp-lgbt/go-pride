package pride

const (
	BaseURL           = "https://api.pfp.lgbt/v6"
	GetFlagsURL       = BaseURL + "/data/flag"
	GetFlagsByNameURL = BaseURL + "/data/flag/%s"
	SearchFlagsURL    = BaseURL + "/data/flag/search?q=%s"
	PutImageURL       = BaseURL + "/image"
	EditImageURL      = BaseURL + "/image/%s"
)

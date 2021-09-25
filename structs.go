package pride

import (
	"encoding/json"
	"time"
)

type response struct {
	Code int             `json:"code"`
	Msg  string          `json:"message"`
	Data json.RawMessage `json:"data"`
}

type Flag struct {
	DefaultAlpha uint   `json:"default_alpha"`
	Name         string `json:"tooltip"`
}

type ImageData struct {
	ID      string    `json:"id"`
	Expires time.Time `json:"expires"`
	Size    int       `json:"size"`
}

type EditImageParams struct {
	Animated bool `json:"animated"`
	Async    struct {
		CallbackURL string `json:"callbackURL"`
		Key         string `json:"key"`
		UseCallback bool   `json:"useCallback"`
	} `json:"async"`
	BorderWidth int  `json:"borderWidth"`
	Cdn         bool `json:"cdn"`
	Cropping    struct {
		FlipX  bool `json:"flipX"`
		FlipY  bool `json:"flipY"`
		Height int  `json:"height"`
		Left   int  `json:"left"`
		Rotate int  `json:"rotate"`
		Top    int  `json:"top"`
		Width  int  `json:"width"`
	} `json:"cropping"`
	Flags        []string `json:"flags"`
	FlagsOpacity int      `json:"flagsOpacity"`
	Framerate    int      `json:"framerate"`
	Layout       string   `json:"layout"`
	Style        string   `json:"style"`
}

type EditImageResponse struct {
	Async   bool   `json:"async"`
	Expires string `json:"expires"`
	URL     string `json:"url"`
}

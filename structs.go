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
		CallbackURL string `json:"callbackURL,omitempty"`
		Key         string `json:"key,omitempty"`
		UseCallback bool   `json:"useCallback,omitempty"`
	} `json:"async,omitempty"`
	BorderWidth int  `json:"borderWidth,omitempty"`
	Cdn         bool `json:"cdn"`
	Cropping    struct {
		FlipX  bool `json:"flipX,omitempty"`
		FlipY  bool `json:"flipY,omitempty"`
		Height int  `json:"height,omitempty"`
		Left   int  `json:"left,omitempty"`
		Rotate int  `json:"rotate,omitempty"`
		Top    int  `json:"top,omitempty"`
		Width  int  `json:"width,omitempty"`
	} `json:"cropping,omitempty"`
	Flags        []string `json:"flags"`
	FlagsOpacity int      `json:"flagsOpacity,omitempty"`
	Framerate    int      `json:"framerate,omitempty"`
	Layout       string   `json:"layout"`
	Style        string   `json:"style"`
}

type EditImageResponse struct {
	Async   bool   `json:"async"`
	Expires string `json:"expires"`
	URL     string `json:"url"`
}

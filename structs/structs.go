package structs

type Stream struct {
	StreamID       string   `json:"id"`
	UserID         string   `json:"user_id"`
	UserName       string   `json:"user_name"`
	GameID         string   `json:"game_id"`
	StreamType     string   `json:"type"`
	StreamTitle    string   `json:"title"`
	ViewerCount    int      `json:"viewer_count"`
	StartedAt      string   `json:"started_at"`
	StreamLanguage string   `json:"language"`
	ThumbnailURL   string   `json:"thumbnail_url"`
	TagIDS         []string `json:"tag_ids"`
}

type Streams struct {
	StreamData []Stream `json:"data"`
}

type GeneralMessage struct {
	Message string `json:"message"`
}

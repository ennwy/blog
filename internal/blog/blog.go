package blog

type Post struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body,omitempty"`
}

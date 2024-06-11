package astcontructor

type Node struct {
    Type     string `json:"type"`
    Attrs    map[string]string `json:"attributes"`
    Text     string `json:"text"`
    Children []*Node `json:"children"`
}

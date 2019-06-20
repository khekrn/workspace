package reqs

// TokenizerRequest for tokenizer handler
type TokenizerRequest struct {
	Text string `json:"text"`
}

// TokenizerResponse for tokenizer handler
type TokenizerResponse struct {
	Tokens []string `json:"tokens"`
}

// WordCountRequest - word count handler
type WordCountRequest struct {
	Text string `json:"text"`
}

// WordCountResponse - word count response
type WordCountResponse struct {
	WordCount map[string]int `json:"wordcount"`
}

// CleanTextRequest - clean text handler
type CleanTextRequest struct {
	Text string `json:"text"`
}

// CleanTextResponse - response for clean text handler
type CleanTextResponse struct {
	CleanText string `json:"cleanText"`
}

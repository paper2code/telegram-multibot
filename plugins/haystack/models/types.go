package models

type Haystack struct {
	Answers  []HaystackAnswer `json:"answers"`
	Question string           `json:"question"`
}

type HaystackAnswer struct {
	Answer      string                 `json:"answer"`
	Context     string                 `json:"context"`
	DocumentID  string                 `json:"document_id"`
	Meta        map[string]interface{} `json:"meta"`
	OffsetEnd   int                    `json:"offset_end"`
	OffsetStart int                    `json:"offset_start"`
	Probability float64                `json:"probability"`
	Score       int                    `json:"score"`
}

type HaystackAnswerMeta struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SegmentID int    `json:"segment_id"`
	URL       string `json:"url"`
}

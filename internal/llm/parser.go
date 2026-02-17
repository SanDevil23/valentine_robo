package llm

import (
	"encoding/json"
	"errors"
)

type File struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

type Response struct {
	Files []File `json:"files"`
}

func ParseResponse(raw string) (*Response, error) {
	var res Response

	err := json.Unmarshal([]byte(raw), &res)
	if err != nil {
		return nil, errors.New("failed to parse response: " + err.Error())
	}

	if len(res.Files) == 0 {
		return nil, errors.New("no files found in response")
	}

	return &res, nil

}
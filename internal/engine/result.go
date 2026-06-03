package engine

import (
	"fmt"
	"time"
)

type ScanResult struct {
	Target      string    `json:"target"`
	Name        string    `json:"name"`
	Severity    string    `json:"severity"`
	Matched     bool      `json:"matched"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}

type Results struct {
	Items []ScanResult `json:"items"`
}

func NewResults() *Results {
	return &Results{
		Items: []ScanResult{},
	}
}

func (r *Results) Add(result ScanResult) {
	result.Timestamp = time.Now()
	r.Items = append(r.Items, result)
}

func (r *Results) Print() {
	for _, item := range r.Items {
		fmt.Printf("[%s] %s → %s\n", item.Severity, item.Name, item.Target)
	}
}

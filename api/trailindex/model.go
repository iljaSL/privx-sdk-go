//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package trailindex

// Params struct for pagination queries.
type Params struct {
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Sortdir string `json:"sortdir,omitempty"`
}

// IndexResponse trail index response definition
type IndexResponse struct {
	ConnID      string `json:"connection_id,omitempty"`
	ChanID      string `json:"channel_id,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	ChannelType string `json:"type,omitempty"`
	TimeStamp   string `json:"timestamp,omitempty"`
	Content     string `json:"content,omitempty"`
	Position    int    `json:"position,omitempty"`
}

// SearchRequestObject search request object definition
type SearchRequestObject struct {
	ConnID    string `json:"connection_id"`
	ChanID    string `json:"channel_id"`
	Protocol  string `json:"protocol"`
	Keywords  string `json:"keywords"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	StartPos  int    `json:"start_position"`
	EndPos    int    `json:"end_position"`
}

// Connection connection definition
type Connection struct {
	ConnID string `json:"connection_id,omitempty"`
	Status string `json:"status,omitempty"`
}

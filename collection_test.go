package huggo

import (
	"encoding/json"
	"testing"
)

func TestGated_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantVal interface{}
		wantErr bool
	}{
		{
			name:    "bool: true",
			json:    `true`,
			wantVal: true,
			wantErr: false,
		},
		{
			name:    "bool: false",
			json:    `false`,
			wantVal: false,
			wantErr: false,
		},
		{
			name:    "string",
			json:    `"gated"`,
			wantVal: "gated",
			wantErr: false,
		},
		{
			name:    "invalid type (number)",
			json:    `69`,
			wantVal: nil,
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			json:    `not JSON`,
			wantVal: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var g Gated
			err := json.Unmarshal([]byte(tt.json), &g)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if gotVal := g.Value(); gotVal != tt.wantVal {
					t.Errorf("Value() = %v, want %v", gotVal, tt.wantVal)
				}
			}
		})
	}
}

func TestGatedInResponse_UnmarshalJSON(t *testing.T) {
	type Response struct {
		Gated Gated `json:"gated"`
	}

	tests := []struct {
		name    string
		json    string
		wantVal interface{}
		wantErr bool
	}{
		{
			name:    "bool in response",
			json:    `{"gated":true}`,
			wantVal: true,
			wantErr: false,
		},
		{
			name:    "string in response",
			json:    `{"gated":"gated"}`,
			wantVal: "gated",
			wantErr: false,
		},
		{
			name:    "invalid type in response",
			json:    `{"gated":42}`,
			wantVal: nil,
			wantErr: true,
		},
		{
			name:    "missing gated field",
			json:    `{}`,
			wantVal: nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response Response
			err := json.Unmarshal([]byte(tt.json), &response)

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				if gotVal := response.Gated.Value(); gotVal != tt.wantVal {
					t.Errorf("Value() = %v, want %v", gotVal, tt.wantVal)
				}
			}
		})
	}
}

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindDigimons(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		err     string
		wantErr bool
	}{
		{
			name:    "Digimons found, respone OK",
			url:     "https://digimon-api.vercel.app/api/digimon",
			err:     "",
			wantErr: false,
		},
		{
			name:    "Digimons found, respone OK",
			url:     "",
			err:     "Request could not be done",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s := NewService(tt.url)
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.GetData()
			if tt.wantErr {
				assert.Nil(t, resp)
			} else {
				assert.NotNil(t, resp)
			}
			if err != nil {
				assert.Equal(t, tt.err, err.Error())
			} else {
				assert.Equal(t, nil, err)
			}
		})
	}
}

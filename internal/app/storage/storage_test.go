package storage_test

import (
	"github.com/SamoylenkoVadim/golang-practicum/internal/app/storage"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStorage_Save(t *testing.T) {

	s := storage.New()

	tests := []struct {
		name       string
		key        string
		link       string
		wantResult bool
	}{
		{
			name:       "simple test #1",
			key:        "id1",
			link:       "http://yandex.ru",
			wantResult: true,
		},
		{
			name:       "simple test #2",
			key:        "id1",
			link:       "http://yandex.ru",
			wantResult: false,
		},
		{
			name:       "simple test #3",
			key:        "id2",
			link:       "http://google.ru",
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Save(tt.key, tt.link)

			if tt.wantResult == true {
				require.NoError(t, err)
			}

			if tt.wantResult == false {
				require.Error(t, err)
			}
		})
	}
}

func TestStorage_GetValue(t *testing.T) {

	s := storage.New()
	s.Save("id1", "http://yandex.ru")
	s.Save("id2", "http://google.ru")
	s.Save("id3", "http://google.ru")

	tests := []struct {
		name       string
		key        string
		link       string
		wantResult bool
	}{
		{
			name:       "simple test #1",
			key:        "id1",
			link:       "http://yandex.ru",
			wantResult: true,
		},
		{
			name:       "simple test #2",
			key:        "id2",
			link:       "http://google.ru",
			wantResult: true,
		},
		{
			name:       "simple test #3",
			key:        "id3",
			link:       "http://google.ru",
			wantResult: true,
		},
		{
			name:       "simple test #4",
			key:        "id4",
			link:       "",
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, ok := s.GetValue(tt.key)

			if tt.wantResult == true {
				require.True(t, ok)
				require.Equal(t, tt.link, value)
			}

			if tt.wantResult == false {
				require.False(t, ok)
				require.Equal(t, "", value)
			}
		})
	}
}

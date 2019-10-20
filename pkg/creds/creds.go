package creds

import (
	"encoding/json"

	"github.com/andrskom/jwa-console/pkg/storage/file"
)

type Component struct {
	db   file.LazyReadWriter
}

func New(db file.LazyReadWriter) *Component {
	return &Component{db: db}
}

func (s *Component) Save(m *Model) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return s.db.WriteData(data)
}

func (s *Component) Get() (*Model, error) {
	data, err := s.db.ReadData()
	if err != nil {
		return nil, err
	}
	var res Model
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

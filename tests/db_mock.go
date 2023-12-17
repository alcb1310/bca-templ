package tests

type DBMock struct{}

func (m *DBMock) Health() map[string]string {
	return map[string]string{
		"message": "It's healthy",
	}
}

package api

import "sync"

const (
	BASE_GET_URL = "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/%s"
	DATE_PATTERN = `^(?P<Year>\d{4})-(?P<Month>\d{1,2})-(?P<Day>\d{1,2})$`
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

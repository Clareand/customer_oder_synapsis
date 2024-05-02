package pagination

import "math"

func FormatPage(total int, limit int, page int) (newPage int, lastPage int) {

	lastPage = int(math.Ceil(float64(total) / float64(limit)))
	newPage = page
	if page < 1 {
		newPage = 1
	}

	return newPage, (lastPage)
}

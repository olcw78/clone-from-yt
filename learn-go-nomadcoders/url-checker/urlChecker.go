package urlchecker

import (
	"errors"
	"fmt"
	"net/http"
)

// errors
var errRequestFail = errors.New("request failed")

func HitURL(url string) error {
	fmt.Println("start hit URL", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFail
	}

	fmt.Println(err, res.StatusCode)
	return nil
}

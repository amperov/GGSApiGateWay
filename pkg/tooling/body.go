package tooling

import (
	"encoding/json"
	"io"
)

func UnmarshallAll(closer io.ReadCloser, out any) error {
	bodyBytes, err := io.ReadAll(closer)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		return err
	}
	return nil
}

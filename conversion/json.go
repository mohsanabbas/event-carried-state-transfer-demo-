package conversion

import (
	"encoding/json"

	"event-carried-state-transfer/state"
)

func MarshalIndent(data state.State) (string, error) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

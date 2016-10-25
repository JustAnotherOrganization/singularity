package singularity

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

//HTTPHelper helps with HTTP Calls.
type HTTPHelper struct {
	Client    *http.Client
	Transport string
}

var (
	errFailedToMarshal   = errors.New("Failed to marshal body")
	errFailedToUnmarshal = errors.New("Failed to unmarshal to body")
)

//Path is SlackAPI + path
func (helper *HTTPHelper) post(path string, response interface{}, i ...string) ([]byte, error) {
	resp, err := helper.Client.PostForm(helper.Transport+SlackAPI+"rtm.start", stringsToValues(i...))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if response != nil {
		err = json.Unmarshal(bytes, response)
		if err != nil {
			return bytes, errFailedToUnmarshal
		}
	}

	return bytes, nil
}

func stringsToValues(values ...string) url.Values {
	rValues := make(url.Values)
	if len(values)%2 != 0 {
		values = values[:len(values)-1] //Must be even.
	}
	for i := 0; i < len(values); i++ {
		rValues.Add(values[i], values[i+1])
		i++
	}
	return rValues
}

func (helper *HTTPHelper) get() error {

	return nil
}

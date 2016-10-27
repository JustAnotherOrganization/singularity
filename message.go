package singularity

import "encoding/json"

//tMessage is a message that is thrown between channels. TODO Replace with something less dumb.
type tMessage struct {
	Body interface{} //ByteForm please.
}

//GetBytes ...
func (message *tMessage) GetBytes() ([]byte, error) {
	switch s := message.Body.(type) {
	case []byte:
		return s, nil
	default:
		return json.Marshal(s)

	}
}

//GetInterface ...
func (message *tMessage) GetInterface() (interface{}, error) {
	switch s := message.Body.(type) {
	case []byte:
		var i interface{}
		err := json.Unmarshal(s, &i)
		return i, err
	default:
		return s, nil
	}
}

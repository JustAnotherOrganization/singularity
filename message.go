package singularity

import "encoding/json"

//ChanMessage is a message that is thrown between channels. TODO Replace with something less dumb.
type ChanMessage struct {
	Body interface{} //ByteForm please.
}

//GetBytes ...
func (message *ChanMessage) GetBytes() ([]byte, error) {
	switch s := message.Body.(type) {
	case []byte:
		return s, nil
	default:
		return json.Marshal(s)

	}
}

//GetInterface ...
func (message *ChanMessage) GetInterface() (interface{}, error) {
	switch s := message.Body.(type) {
	case []byte:
		var i interface{}
		err := json.Unmarshal(s, &i)
		return i, err
	default:
		return s, nil
	}
}

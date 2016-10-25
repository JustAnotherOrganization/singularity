package singularity

import "encoding/json"

//Message Not to be confused with slack.Message
type Message struct {
	Body interface{} //ByteForm please.
}

//GetBytes ...
func (message *Message) GetBytes() ([]byte, error) {
	switch s := message.Body.(type) {
	case []byte:
		return s, nil
	default:
		return json.Marshal(s)

	}
}

//GetInterface ...
func (message *Message) GetInterface() (interface{}, error) {
	switch s := message.Body.(type) {
	case []byte:
		var i interface{}
		err := json.Unmarshal(s, &i)
		return i, err
	default:
		return s, nil
	}
}

//TODO json validation.
//WARNING Not safe.
func (message *Message) preParseString(key string) string {
	var buffer []byte
	body, ok := message.Body.([]byte) //FIXME Make this an actual thing
	if !ok {
		return ""
	}
	qO := false   //quoteOpen
	eonf := false //end on next flush
	escaped := false
	for ii := 0; ii < len(body); ii++ {
		switch body[ii] {
		case '"':
			if !escaped {
				qO = !qO //Toggle.
				//Flush.
				if !qO {
					if len(buffer) > 0 {
						if eonf {
							return string(buffer)
						}
						if string(buffer) == key {
							eonf = true
						}
						buffer = make([]byte, 0) //Clear.
					}
				}
			} else {
				escaped = false
			}
		case '{', '[', ':': // Level handling.
			//Ignore.
		case ',', '}', ']':
			if eonf && !qO {
				return ""
			}
		case '\\':
			escaped = !escaped //TODO make sure the next character is the one we are escaping.
		default:
			if escaped {
				// Ignore?
				escaped = false
			} else if qO {
				buffer = append(buffer, body[ii])
			}
		}
	}
	return "" //If it isn't found, reutrn nil.
}

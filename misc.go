package singularity

import "strings"

func getFunctionParameters(function string) []string {
	if function == "" {
		return []string{} //Empty string because nils are mean.
	}
	start := 0
	end := 1
	for ii, char := range function {
		if char == '(' {
			start = ii
		}
		if char == ')' {
			end = ii
			break //Need to break or else it might process return values
		}
	}
	return strings.Split(function[start+1:end], ", ")
}

//TODO json validation.
//WARNING Not safe.
func preParseString(key string, body []byte) string {
	var buffer []byte
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

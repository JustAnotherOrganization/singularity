package singularity

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

//NewHandler1 ...
func NewHandler1() *EventAPIHandler {
	eapih := &EventAPIHandler{}
	eapih.handlers = make(map[string][]interface{})
	return eapih
}

//EventAPIHandler ...
type EventAPIHandler struct {
	sync.Mutex
	handlerList
}

type handlerList struct {
	sync.Mutex
	handlers map[string][]interface{}
}

/*
One of the things that I'm thinking about, is requiting a naming convention for finding the `type`
*/

//WARNING, does not lock.
func (handler *EventAPIHandler) registerHandler(key string, function interface{}) error {
	if function == nil {
		return errors.New("Function can not be nul!")
	}
	rtype := fmt.Sprintf("%v", reflect.TypeOf(function))
	if strings.HasPrefix(rtype, "func(") {
		params := getFunctionParameters(rtype)
		if len(params) != 2 && params[1] == "*singularity.SlackInstance" {
			return errors.New("Function parameter missmatch!")
		}
		handler.handlerList.handlers[key] = append(handler.handlerList.handlers[key], function)
		return nil
	}

	return nil
}

func (handler *EventAPIHandler) execute(key string, body []byte, instance *SlackInstance) (err error) {
	handler.handlerList.Lock()
	defer handler.handlerList.Unlock()
	for _, function := range handler.handlerList.handlers[key] {
		if func() bool {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("Recovered from Panic: %v\n", err) //TODO Better Error Handling.
				}
			}()
			param0 := reflect.TypeOf(function).In(0) //Should be Param 0.
			value := reflect.New(param0).Interface()
			err = json.Unmarshal(body, &value)
			if err != nil {
				fmt.Printf("Error: %v\n", err) //TODO Better Error Handling.
				return false
			}
			reflect.ValueOf(function).Call([]reflect.Value{reflect.ValueOf(value).Elem(), reflect.ValueOf(instance)})
			return false
		}() {
			break
		}
	}

	return nil
}

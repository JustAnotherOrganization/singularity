package singularity

import "testing"

//TODO: move this file when config.go is moved

//TestConfigCheckBoolNotExistsReturnsFalse test
func TestConfigCheckBool(t *testing.T) {
	config := &defaultConfig{config: make(map[string]interface{})}
	v, ok := config.CheckBool("does_not_exist")

	if v != false {
		t.Error("Expected value of false, got ", v)
	}

	if ok != false {
		t.Error("Expected ok of false, got ", v)
	}
}

//TestConfigCheckBoolExistsReturnsTrue test
func TestConfigCheckBoolExistsReturnsTrue(t *testing.T) {
	config := &defaultConfig{config: make(map[string]interface{})}
	config.config["item"] = true
	v, ok := config.CheckBool("item")

	if v != true {
		t.Error("Expected value of true, got ", v)
	}

	if ok != true {
		t.Error("Expected ok of true, got ", v)
	}
}

//TestConfigGetBoolExistsReturnsTrue test
func TestConfigGetBoolExistsReturnsTrue(t *testing.T) {
	config := &defaultConfig{config: make(map[string]interface{})}
	config.config["item"] = true
	v := config.GetBool("item")

	if v != true {
		t.Error("Expected value of true, got ", v)
	}
}

//TestConfigCheckStringDoesNotExistsReturnsFalse test
func TestConfigCheckStringDoesNotExistsReturnsFalse(t *testing.T) {
	config := &defaultConfig{config: make(map[string]interface{})}
	v, ok := config.CheckString("item")

	if v != "" {
		t.Error("Expected value of 'item-value', got ", v)
	}

	if ok != false {
		t.Error("Expected ok of false, got ", v)
	}
}

//TestConfigCheckStringExistsReturnsTrue test
func TestConfigCheckStringExistsReturnsTrue(t *testing.T) {
	config := &defaultConfig{config: make(map[string]interface{})}
	config.config["item"] = "item-value"
	v, ok := config.CheckString("item")

	if v != "item-value" {
		t.Error("Expected value of 'item-value', got ", v)
	}

	if ok != true {
		t.Error("Expected ok of true, got ", v)
	}
}

//TestConfigGetStringExistsReturnsTrue test
func TestConfigGetStringExistsReturnsTrue(t *testing.T) {
	config := &defaultConfig{config: make(map[string]interface{})}
	config.config["item"] = "item-value"
	v := config.GetString("item")

	if v != "item-value" {
		t.Error("Expected value of 'item-value', got ", v)
	}
}

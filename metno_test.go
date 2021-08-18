package env

import (
	"testing"
	"time"
)

type ValidTimeStruct struct {
	// Home should match Environ because it has a "env" field tag.
	RefTime time.Time `env:"REF_TIME"`
}

func TestTimeMarshal(t *testing.T) {
	timeInRFC3339 := "2021-08-18T00:00:00Z"
	refTime, err := time.Parse(time.RFC3339, timeInRFC3339)
	if err != nil {
		t.Errorf("Expected parsing of timestamp %s to work; Got error: %v", timeInRFC3339, err)
	}
	s := ValidTimeStruct{
		RefTime: refTime,
	}
	envSet, err := Marshal(&s)
	if err != nil {
		t.Errorf("Expected marshal to env with no errors; Got failed to marshal struct to: %v", err)
	}

	if envSet["REF_TIME"] != timeInRFC3339 {
		t.Errorf("Expected marshalling of time to be in RFC3339; Got %s", envSet["REF_TIME"])
	}
}

func TestTimeUnmarshal(t *testing.T) {
	environ := map[string]string{
		"REF_TIME": "2021-08-18T00:00:00Z",
	}
	var validTimeStruct ValidTimeStruct
	err := Unmarshal(environ, &validTimeStruct)
	if err != nil {
		t.Errorf("Expected no error but got '%s'", err)
	}
}

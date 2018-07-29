package test

import (
	"testing"

	"github.com/rabadiw/gotester/src/assert"
	"github.com/rabadiw/w3k/src"
)

func Test_EmptyApp(t *testing.T) {

	app := &w3k.App{}

	(&assert.T{T: t}).
		NotNil(app)
}

func Test_NewApp(t *testing.T) {

	T := &assert.T{T: t}

	app := &w3k.App{
		Name: "TestApp",
		Config: map[string]interface{}{
			"KeyA": "ValueA",
		},
		Pipe: []w3k.Piper{},
	}

	// Add

	T.NotNil(app)
}

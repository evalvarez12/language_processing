package processing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessing(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		[]string{"word1", "word2"},
		Process([]byte("word1 word2")),
	)

	assert.Equal(
		[]string{"word"},
		Process([]byte("(word?!)")),
	)

	assert.Equal(
		[]string{"words"},
		Process([]byte("<brwords>")),
	)

	assert.Equal(
		[]string{"words"},
		Process([]byte(" words ?")),
	)

	assert.Equal(
		[]string{"this", "is", "a", "#"},
		Process([]byte("this is a 23241")),
	)

}

package commons

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_template(t *testing.T) {

	input := map[string]interface{}{
		"name": "this iis the name",
		"age":  1,
		"dead": true,
	}
	api := GenerateApiSpecFile("name", "", input)

	fmt.Println(string(api))
}

const ExpectedProperties = `
      age:
        type: integer
        required: false

      dead:
        type: boolean
        required: false

      name:
        type: string
        required: false
`

func TestGenerateApiSpecFile(t *testing.T) {
	input := map[string]interface{}{
		"name": "this iis the name",
		"age":  1,
		"dead": true,
	}

	spec := generateProperties(input)
	assert.Equal(t, ExpectedProperties, string(spec))

	model := generateModelProperties(input)
	fmt.Println(string(model))
}

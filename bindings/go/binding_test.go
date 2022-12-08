package norg_test

import (
	"context"
	"testing"

	"github.com/icyrogue/tree-sitter-norg/blob/main/bindings/go"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(`* Heading`), norg.GetLanguage())
	assert.NoError(err)
}

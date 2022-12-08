package norg

//#include "../../src/tree_sitter/parser.h"
//TSLanguage *tree_sitter_norg();

import "C"
import (
	"unsafe"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_norg())
	return sitter.NewLanguage(ptr)
}

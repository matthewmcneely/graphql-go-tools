package astvalidation

import (
	"github.com/matthewmcneely/graphql-go-tools/v2/pkg/astvisitor"
)

var reservedFieldPrefix = []byte("__")

// Rule is hook to register callback functions on the Walker
type Rule func(walker *astvisitor.Walker)

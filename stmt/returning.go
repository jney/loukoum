package stmt

import (
	"github.com/ulule/loukoum/v3/token"
	"github.com/ulule/loukoum/v3/types"
)

// Returning is a RETURNING clause.
type Returning struct {
	Columns []SelectExpression
}

// NewReturning returns a new Returning instance.
func NewReturning(exprs []SelectExpression) Returning {
	return Returning{
		Columns: exprs,
	}
}

// Write exposes statement as a SQL query.
func (returning Returning) Write(ctx types.Context) {
	ctx.Write(token.Returning.String())
	ctx.Write(" ")

	for i := range returning.Columns {
		if i > 0 {
			ctx.Write(", ")
		}
		returning.Columns[i].Write(ctx)
	}
}

// IsEmpty returns true if statement is undefined.
func (returning Returning) IsEmpty() bool {
	return len(returning.Columns) == 0
}

// Ensure that Returning is a Statement
var _ Statement = Returning{}

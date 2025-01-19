package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction object is Imutable and has no identifier
type Transaction struct {
	amount int
	from uuid.UUID
	to uuid.UUID
	createdAt time.Time
}
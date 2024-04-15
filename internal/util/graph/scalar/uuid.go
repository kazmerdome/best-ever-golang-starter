package scalar

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

// MarshalUuid implements the graphql.Marshaler interface for the Uuid scalar.
func MarshalUuid(uuid uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, uuid.String()))
	})
}

// UnmarshalUuid implements the graphql.Unmarshaler interface for the Uuid scalar.
func UnmarshalUuid(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		parsedUUID, err := uuid.Parse(v)
		if err != nil {
			return uuid.UUID{}, fmt.Errorf("invalid UUID format")
		}
		return parsedUUID, nil
	default:
		return uuid.UUID{}, fmt.Errorf("UUID must be a string")
	}
}

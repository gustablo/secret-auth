package commons

import "github.com/google/uuid"

func Generate(str *string) {
	*str = uuid.NewString()
}

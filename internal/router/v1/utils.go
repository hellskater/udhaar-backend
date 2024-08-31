package v1

import (
	"fmt"
	"strings"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofrs/uuid"
	"github.com/hellskater/udhaar-backend/internal/router/extension"
	"github.com/labstack/echo"
)

func bindAndValidate(c echo.Context, i interface{}) error {
	return extension.BindAndValidate(c, i)
}

func parseUUIDsFromQueryParam(target *[]uuid.UUID) vd.RuleFunc {
	return func(value interface{}) error {
		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("invalid input type")
		}

		ids := strings.Split(str, ",")

		var uuids []uuid.UUID
		for _, id := range ids {
			uid, err := uuid.FromString(id)
			if err != nil {
				return fmt.Errorf("invalid UUID: %w", err)
			}
			uuids = append(uuids, uid)
		}

		*target = uuids
		return nil
	}
}

package v1

import (
	"net/http"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofrs/uuid"
	"github.com/hellskater/udhaar-backend/internal/router/extension/herror"
	"github.com/labstack/echo"
)

type GetGroupsRequest struct {
	GroupIDs       string      `query:"groupIds" json:"groupIds"`
	ParsedGroupIDs []uuid.UUID `json:"-"`
}

func (r *GetGroupsRequest) Validate() error {
	return vd.ValidateStruct(r,
		vd.Field(&r.GroupIDs, vd.Required, vd.By(parseUUIDsFromQueryParam(&r.ParsedGroupIDs))),
	)
}

func (h *Handlers) GetGroups(c echo.Context) error {
	var req GetGroupsRequest
	if err := bindAndValidate(c, &req); err != nil {
		return err
	}

	cfs, err := h.Repo.GetGroups(req.ParsedGroupIDs)
	if err != nil {
		return herror.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, cfs)
}

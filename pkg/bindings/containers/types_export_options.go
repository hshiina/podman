package containers

import (
	"net/url"

	"github.com/containers/podman/v3/pkg/bindings/internal/util"
)

/*
This file is generated automatically by go generate.  Do not edit.
*/

// Changed
func (o *ExportOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams
func (o *ExportOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

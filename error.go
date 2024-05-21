package registerer

import "github.com/PeerXu/errors-go"

var (
	ErrUnsupported, ErrUnsupportedFn = errors.NewErrorAndErrorFunc[string]("unsupported")
)

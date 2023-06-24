package upload

import (
	"io"

	"github.com/turt2live/matrix-media-repo/common/rcontext"
)

func LimitStream(ctx rcontext.RequestContext, r io.ReadCloser) io.ReadCloser {
	if ctx.Config.Uploads.MaxSizeBytes > 0 {
		return io.NopCloser(io.LimitReader(r, ctx.Config.Uploads.MaxSizeBytes))
	} else {
		return r
	}
}
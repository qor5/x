package statusx

import (
	"cmp"
	"context"
	"log/slog"
	"net/http"
	"strings"

	"connectrpc.com/connect"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"

	vproto "github.com/theplant/validator/proto"

	"github.com/qor5/x/v3/httpx"
	"github.com/qor5/x/v3/i18nx"
	"github.com/qor5/x/v3/jsonx"
)

const HeaderEnsureConnectError = "x-ensure-connect-error"

var AllowHeaders = []string{
	http.CanonicalHeaderKey(HeaderEnsureConnectError),
}

func EnsureConnectError(ctx context.Context) bool {
	return metadata.ExtractIncoming(ctx).Get(HeaderEnsureConnectError) == "true"
}

func NewVProtoHTTPErrorWriter(ib *i18nx.I18N) func(http.Handler) http.Handler {
	conf := &HTTPErrorWriterConfig{I18N: ib}
	conf = conf.WithHTTPWriteErrorHook(VProtoHTTPWriteErrorHook) // Compatible with vproto
	return HTTPErrorWriter(conf)
}

func VProtoHTTPWriteErrorHook(next HTTPWriteErrorFunc) HTTPWriteErrorFunc {
	errWriter := connect.NewErrorWriter()
	return func(ctx context.Context, input *HTTPWriteErrorInput) (*HTTPWriteErrorOutput, error) {
		if EnsureConnectError(ctx) {
			// Why not use statusx.TranslateError? Just to avoid affecting the original prottp releated logic.
			err, _ := TranslateStatusErrorOnly(ctx, input.Conf.I18N, input.Err)
			written := WriteConnectErrorOnly(errWriter, input.W, input.R, err)
			return &HTTPWriteErrorOutput{Written: written}, nil
		}

		err, w, r := input.Err, input.W, input.R

		// Why not use statusx.TranslateError? Just to avoid affecting the original prottp releated logic.
		err, translated := TranslateStatusErrorOnly(ctx, input.Conf.I18N, err)
		if !translated {
			return &HTTPWriteErrorOutput{Written: false}, nil // ignore errors that are not statusx.StatusError
		}

		werr := WriteVProtoHTTPError(err, w, r)
		if werr != nil {
			slog.ErrorContext(ctx, "Failed to write vproto http error", "error", werr)
		}

		return &HTTPWriteErrorOutput{Written: true}, nil
	}
}

func WriteVProtoHTTPError(err error, w http.ResponseWriter, r *http.Request) (xerr error) {
	// Explicitly close the request body following the pattern from connectUnaryHandlerConn.Close().
	// This ensures consistent resource management and proper cleanup, even though the HTTP server
	// would normally handle this automatically.
	defer func() {
		if err := r.Body.Close(); err != nil {
			if xerr == nil {
				xerr = WrapCode(err, codes.Internal, "failed to close request body").Err()
			}
		}
	}()

	st := Convert(err)

	code := st.Code()
	statusCode := HTTPStatusFromCode(code)

	var errInfo *errdetails.ErrorInfo
	var localizedMessage *errdetails.LocalizedMessage
	var badRequest *errdetails.BadRequest
	for _, d := range st.Details() {
		switch v := d.(type) {
		case *errdetails.ErrorInfo:
			errInfo = v
		case *errdetails.LocalizedMessage:
			localizedMessage = v
		case *errdetails.BadRequest:
			badRequest = v
		}
	}

	verr := &vproto.ValidationError{
		Code:           cmp.Or(errInfo.GetReason(), code.String()),
		DefaultViewMsg: cmp.Or(localizedMessage.GetMessage(), st.Message()),
	}

	if code == codes.InvalidArgument {
		statusCode = http.StatusUnprocessableEntity // 422

		for _, fv := range badRequest.GetFieldViolations() {
			verr.FieldViolations = append(verr.FieldViolations, &vproto.ValidationError_FieldViolation{
				Field:          fv.GetField(),
				Code:           fv.GetReason(),
				DefaultViewMsg: cmp.Or(fv.GetLocalizedMessage().GetMessage(), fv.GetDescription()),
			})
		}
	}

	isJSON := isMimeTypeJSON(w.Header().Get(httpx.HeaderContentType))
	if w.Header().Get(httpx.HeaderContentType) == "" {
		isJSON = shouldReturnJSON(r)
		contentType := xprottpContentType
		if isJSON {
			contentType = jsonContentType
		} else if strings.Contains(strings.ToLower(r.Header.Get("Accept")), protoContentType) {
			contentType = protoContentType
		}
		w.Header().Set(httpx.HeaderContentType, contentType)
	}

	w.WriteHeader(statusCode)

	var data []byte
	var werr error
	if isJSON {
		data, werr = jsonx.Marshal(verr)
	} else {
		data, werr = proto.Marshal(verr)
	}
	if werr != nil {
		return WrapCodef(werr, codes.Internal, "failed to marshal error (isJSON:%t)", isJSON).Err()
	}

	_, werr = w.Write(data)
	if werr != nil {
		return WrapCodef(werr, codes.Internal, "failed to write error (isJSON:%t)", isJSON).Err()
	}
	return nil
}

// from connectCodeToHTTP of connect library
func HTTPStatusFromCode(code codes.Code) int {
	// Return literals rather than named constants from the HTTP package to make
	// it easier to compare this function to the Connect specification.
	switch code {
	case codes.Canceled:
		return 499
	case codes.Unknown:
		return 500
	case codes.InvalidArgument:
		return 400
	case codes.DeadlineExceeded:
		return 504
	case codes.NotFound:
		return 404
	case codes.AlreadyExists:
		return 409
	case codes.PermissionDenied:
		return 403
	case codes.ResourceExhausted:
		return 429
	case codes.FailedPrecondition:
		return 400
	case codes.Aborted:
		return 409
	case codes.OutOfRange:
		return 400
	case codes.Unimplemented:
		return 501
	case codes.Internal:
		return 500
	case codes.Unavailable:
		return 503
	case codes.DataLoss:
		return 500
	case codes.Unauthenticated:
		return 401
	default:
		return 500 // same as CodeUnknown
	}
}

// from prottp
const (
	jsonContentType    = "application/json"
	xprottpContentType = "application/x.prottp"
	protoContentType   = "application/proto"
)

func isMimeTypeJSON(contentType string) bool {
	return strings.Contains(strings.ToLower(contentType), jsonContentType)
}

func isContentTypeJSON(r *http.Request) bool {
	return isMimeTypeJSON(r.Header.Get(httpx.HeaderContentType))
}

// shouldReturnJSON returns true if the response to the given request
// should be a JSON object. This is determined by examining the request's
// Accept header. If the Accept header is empty, this function returns
// true if the request's Content-Type header is "application/json".
// Otherwise, this function returns true if "application/json" appears
// before "application/x.prottp" and "application/proto" in the Accept
// header.
func shouldReturnJSON(r *http.Request) bool {
	acceptString := strings.ToLower(r.Header.Get("Accept"))
	if len(acceptString) == 0 {
		return isContentTypeJSON(r)
	}

	jsonIndex := strings.Index(acceptString, jsonContentType)
	xprottpIndex := strings.Index(acceptString, xprottpContentType)
	protoIndex := strings.Index(acceptString, protoContentType)

	if jsonIndex < 0 && xprottpIndex < 0 && protoIndex < 0 {
		return isContentTypeJSON(r)
	}

	if jsonIndex < 0 {
		jsonIndex = 9999
	}
	if xprottpIndex < 0 {
		xprottpIndex = 10000
	}
	if protoIndex < 0 {
		protoIndex = 10001
	}

	return jsonIndex < xprottpIndex && jsonIndex < protoIndex
}

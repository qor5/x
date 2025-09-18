package statusx

import (
	"time"

	"github.com/pkg/errors"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// toProtoMessage converts Go values to proto messages for use with Any.
func toProtoMessage(v any) (proto.Message, error) {
	switch val := v.(type) {
	case proto.Message:
		return val, nil
	case string:
		return wrapperspb.String(val), nil
	case int:
		return wrapperspb.Int64(int64(val)), nil
	case int32:
		return wrapperspb.Int32(val), nil
	case int64:
		return wrapperspb.Int64(val), nil
	case uint:
		return wrapperspb.UInt64(uint64(val)), nil
	case uint32:
		return wrapperspb.UInt32(val), nil
	case uint64:
		return wrapperspb.UInt64(val), nil
	case float32:
		return wrapperspb.Float(val), nil
	case float64:
		return wrapperspb.Double(val), nil
	case bool:
		return wrapperspb.Bool(val), nil
	case []byte:
		return wrapperspb.Bytes(val), nil
	case time.Time:
		return timestamppb.New(val), nil
	case time.Duration:
		return durationpb.New(val), nil
	case map[string]any:
		s, err := structpb.NewStruct(val)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert map to struct")
		}
		return s, nil
	case nil:
		return &emptypb.Empty{}, nil
	default:
		return nil, errors.Errorf("unsupported type for protobuf Any conversion: %T", val)
	}
}

// extractValueFromAny converts protobuf Any back to Go values.
func extractValueFromAny(anyVal *anypb.Any) (any, error) {
	if anyVal == nil {
		return nil, nil
	}

	msg, err := anypb.UnmarshalNew(anyVal, proto.UnmarshalOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal Any")
	}

	switch v := msg.(type) {
	case *wrapperspb.StringValue:
		return v.GetValue(), nil
	case *wrapperspb.Int64Value:
		return v.GetValue(), nil
	case *wrapperspb.Int32Value:
		return v.GetValue(), nil
	case *wrapperspb.UInt64Value:
		return v.GetValue(), nil
	case *wrapperspb.UInt32Value:
		return v.GetValue(), nil
	case *wrapperspb.FloatValue:
		return v.GetValue(), nil
	case *wrapperspb.DoubleValue:
		return v.GetValue(), nil
	case *wrapperspb.BoolValue:
		return v.GetValue(), nil
	case *wrapperspb.BytesValue:
		return v.GetValue(), nil
	case *timestamppb.Timestamp:
		return v.AsTime(), nil
	case *durationpb.Duration:
		return v.AsDuration(), nil
	case *structpb.Struct:
		return v.AsMap(), nil
	case *emptypb.Empty:
		return nil, nil
	default:
		return nil, errors.Errorf("unsupported well-known type in Any: %T", v)
	}
}

// convertArgsToAny converts Go values to protobuf Any types
func convertArgsToAny(args []any) []*anypb.Any {
	anyArgs := make([]*anypb.Any, len(args))
	for i, arg := range args {
		protoMsg, err := toProtoMessage(arg)
		if err != nil {
			panic(errors.Wrap(err, "failed to convert arg to proto message"))
		}
		anyArg, marshalErr := anypb.New(protoMsg)
		if marshalErr != nil {
			panic(errors.Wrap(marshalErr, "failed to marshal proto message to Any"))
		}
		anyArgs[i] = anyArg
	}
	return anyArgs
}

// FieldViolation represents a field-level validation violation with localization capability
type FieldViolation struct {
	Field       string
	Description string
	Reason      string
	Key         string
	Args        []any
}

// Proto converts FieldViolation to protobuf message
func (f *FieldViolation) Proto() *statusv1.FieldViolation {
	violation := &statusv1.FieldViolation{
		Field:       f.Field,
		Description: f.Description,
		Reason:      f.Reason,
	}

	// Add localization if key is provided
	if f.Key != "" {
		violation.Localized = &statusv1.Localized{
			Key:  f.Key,
			Args: convertArgsToAny(f.Args),
		}
	}

	return violation
}

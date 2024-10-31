package responses

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Format formats an ErrorResponse, replacing whatever {$PLACEHOLDER} out there with
// the given values. This may be used right now, and maybe has no place, but it'll exist.
func (errorResponse *ErrorResponse) Format(args ...any) *ErrorResponse {
	text := errorResponse.Error
	for _, arg := range args {
		var replacement string
		switch arg.(type) {
		case float32:
			replacement = humanize.Ftoa(float64(arg.(float32)))
		case float64:
			replacement = humanize.Ftoa(arg.(float64))
		case string:
			replacement = arg.(string)
		case bool:
			replacement = strconv.FormatBool(arg.(bool))
		default:
			t := reflect.ValueOf(arg)
			switch t.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				replacement = humanize.Comma(t.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				replacement = humanize.Comma(int64(t.Uint()))
			default:
				replacement = fmt.Sprintf("%v", arg)
			}
		}
		text = strings.Replace(text, "{$PLACEHOLDER}", replacement, 1)
	}
	return &ErrorResponse{Code: errorResponse.Code, Error: text}
}

// Reply is a short-hand method that allows for an idiomatic way of replying to the
// request with an error response.
func (errorResponse *ErrorResponse) Reply(ctx *fiber.Ctx) {
	err := ctx.Status(errorResponse.Code).JSON(errorResponse)
	if err != nil {
		Logger(ctx).Err(err).Msg("Failed to send error response")
	}
}

// HandleErr is a utility method that handles errors on a request, this sends a VagueError to the client
// and logs the error into the console.
func HandleErr(ctx *fiber.Ctx, err error) {
	VagueError.Reply(ctx)
	Logger(ctx).Err(err).Msg("Encountered an Error")
}

var InvalidPayload = ErrorResponse{Code: http.StatusBadRequest, Identifier: 1, Error: "Invalid payload."}
var VagueError = ErrorResponse{Code: http.StatusBadRequest, Identifier: 2, Error: "An error occurred while trying to execute this task."}
var UnknownClient = ErrorResponse{Code: http.StatusForbidden, Identifier: 3, Error: "Unknown client."}

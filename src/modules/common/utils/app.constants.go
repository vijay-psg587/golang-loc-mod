package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type API_CONTEXT_VARS struct {
	APPCONFIG string
	TABLENAME string
}

type CORS_HEADERS struct {
	ACCESS_CONTROL_ALLOW_ORIGIN_NAME      string
	ACCESS_CONTROL_ALLOW_HEADERS_NAME     string
	ACCESS_CONTROL_ALLOW_METHODS_NAME     string
	ACCESS_CONTROL_ALLOW_CREDENTIALS_NAME string
	ACCESS_CONTROL_MAX_AGE_NAME           string
	ACCEPT_METHODS                        []string
	ACCEPT_HEADERS                        []string
	ORIGIN                                []string
	WHITELIST                             []string
}

const (
	HYPEN_CHAR                string = "-"
	COMMA_CHAR                string = ","
	DOT_CHAR                  string = "."
	COLON_CHAR                string = ":"
	SLASH                     string = "/"
	APP_NAME                  string = "golang-loc-mod"
	APP_ENV                   string = "APP_ENV"
	APP_GO_ENV                string = "APP_GO_ENV"
	LAYOUT_WITHOUT_TIME       string = "01-02-2006"
	LAYOUT_US                 string = "January 2, 2006"
	LAYOUT_DEFAULT            string = "01-02-2006 15:04:05"
	EMPTY_STR                 string = ""
	BASE_MB_BYTES             int64  = 1048576
	THRESHOLD_FILE_SIZE       int8   = 2
	AWS_SSM_PARAM_PATH_PREFIX string = "/go-loc-mod/"
)

var (
	CORS = CORS_HEADERS{
		ACCESS_CONTROL_ALLOW_ORIGIN_NAME:      "Access-Control-Allow-Origin",
		ACCESS_CONTROL_ALLOW_METHODS_NAME:     "Access-Control-Allow-Methods",
		ACCESS_CONTROL_ALLOW_HEADERS_NAME:     "Access-Control-Allow-Headers",
		ACCESS_CONTROL_MAX_AGE_NAME:           "Access-Control-Max-Age",
		ACCESS_CONTROL_ALLOW_CREDENTIALS_NAME: "Access-Control-Allow-Credentials",
		ACCEPT_METHODS:                        []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodHead, http.MethodOptions},
		ACCEPT_HEADERS:                        []string{fiber.HeaderOrigin, fiber.HeaderAccept, fiber.HeaderContentType, fiber.HeaderConnection, fiber.HeaderKeepAlive, "X-Reqeuested-With", "x-api-token"},
		ORIGIN:                                []string{"*"},
	}
	CONTEXT_VARS = API_CONTEXT_VARS{
		APPCONFIG: "appConfig",
		TABLENAME: "tableName",
	}
)

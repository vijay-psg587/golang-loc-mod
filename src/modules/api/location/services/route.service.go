package location

import (
	"errors"
	"fmt"

	"os"

	"github.com/gofiber/fiber/v2"

	zlog "github.com/rs/zerolog/log"
	appUtils "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/app-utils"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/enums"
	cerr "github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/models/errors"
	"github.com/vijayakumar-psg587/golang-loc-mod/src/modules/common/utils"
)

func RegisterRoutes(api fiber.Router) {

	zlog.Info().Msg("inside route registration")
	err := &cerr.CustomErrModel{
		Timestamp: appUtils.GetTimeStamp(enums.DEFAULT_LAYOUT),
		Status:    utils.SLASH,
		Code:      enums.ACCESS_ERROR,
		Message:   models.AppConfigModel{}.AppName + os.Getenv("TT"),
		Err:       errors.New("tt"),
	}
	fmt.Println(err)

	api.Get("/loc/all", GetAllLocations)

}

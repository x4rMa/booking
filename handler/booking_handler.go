package handler

import (
	"net/http"

	"github.com/bborbe/log"
	"github.com/bborbe/server/handler/cachingheader"
	"github.com/bborbe/server/handler/contenttype"
	"github.com/bborbe/server/handler/fallback"
	log_handler "github.com/bborbe/server/handler/log"
	"github.com/bborbe/server/handler/static"
	"github.com/bborbe/server/handler_finder"
	"github.com/bborbe/server/handler_finder/part"
	"github.com/bborbe/server/handler_finder/rest"

	booking_date_handler_create "github.com/bborbe/booking/date/handler/create"
	booking_date_handler_delete "github.com/bborbe/booking/date/handler/delete"
	booking_date_handler_get "github.com/bborbe/booking/date/handler/get"
	booking_date_handler_list "github.com/bborbe/booking/date/handler/list"
	booking_date_handler_update "github.com/bborbe/booking/date/handler/update"
	booking_date_service "github.com/bborbe/booking/date/service"

	booking_model_handler_create "github.com/bborbe/booking/model/handler/create"
	booking_model_handler_delete "github.com/bborbe/booking/model/handler/delete"
	booking_model_handler_get "github.com/bborbe/booking/model/handler/get"
	booking_model_handler_list "github.com/bborbe/booking/model/handler/list"
	booking_model_handler_update "github.com/bborbe/booking/model/handler/update"
	booking_model_service "github.com/bborbe/booking/model/service"

	booking_shooting_handler_create "github.com/bborbe/booking/shooting/handler/create"
	booking_shooting_handler_delete "github.com/bborbe/booking/shooting/handler/delete"
	booking_shooting_handler_get "github.com/bborbe/booking/shooting/handler/get"
	booking_shooting_handler_list "github.com/bborbe/booking/shooting/handler/list"
	booking_shooting_handler_update "github.com/bborbe/booking/shooting/handler/update"
	booking_shooting_handler_book "github.com/bborbe/booking/shooting/handler/book"
	booking_shooting_service "github.com/bborbe/booking/shooting/service"

	booking_user_handler_create "github.com/bborbe/booking/user/handler/create"
	booking_user_handler_delete "github.com/bborbe/booking/user/handler/delete"
	booking_user_handler_get "github.com/bborbe/booking/user/handler/get"
	booking_user_handler_list "github.com/bborbe/booking/user/handler/list"
	booking_user_handler_update "github.com/bborbe/booking/user/handler/update"
	booking_user_handler_verifylogin "github.com/bborbe/booking/user/handler/verifylogin"
	booking_user_service "github.com/bborbe/booking/user/service"
)

var logger = log.DefaultLogger

func NewHandler(documentRoot string, dateService booking_date_service.Service, modelService booking_model_service.Service, shootingService booking_shooting_service.Service, userService booking_user_service.Service) http.Handler {
	logger.Debugf("root: %s", documentRoot)
	fileServer := cachingheader.NewCachingHeaderHandler(contenttype.NewContentTypeHandler(http.FileServer(http.Dir(documentRoot))))
	handlerFinder := part.New("")
	handlerFinder.RegisterHandler("/", fileServer)
	handlerFinder.RegisterHandler("/css", fileServer)
	handlerFinder.RegisterHandler("/js", fileServer)
	handlerFinder.RegisterHandler("/images", fileServer)
	handlerFinder.RegisterHandlerFinder("/date", createDateHandlerFinder("/date", dateService))
	handlerFinder.RegisterHandlerFinder("/model", createModelHandlerFinder("/model", modelService))
	handlerFinder.RegisterHandlerFinder("/shooting", createShootingHandlerFinder("/shooting", shootingService))
	handlerFinder.RegisterHandlerFinder("/user", createUserHandlerFinder("/user", userService))
	return log_handler.NewLogHandler(fallback.NewFallback(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", 404)))
}

func createDateHandlerFinder(prefix string, dateService booking_date_service.Service) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(booking_date_handler_list.New(dateService))
	hf.RegisterCreateHandler(booking_date_handler_create.New(dateService))
	hf.RegisterDeleteHandler(booking_date_handler_delete.New(dateService))
	hf.RegisterGetHandler(booking_date_handler_get.New(dateService))
	hf.RegisterUpdateHandler(booking_date_handler_update.New(dateService))
	hf.RegisterPatchHandler(booking_date_handler_update.New(dateService))
	return hf
}

func createModelHandlerFinder(prefix string, modelService booking_model_service.Service) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(booking_model_handler_list.New(modelService))
	hf.RegisterCreateHandler(booking_model_handler_create.New(modelService))
	hf.RegisterDeleteHandler(booking_model_handler_delete.New(modelService))
	hf.RegisterGetHandler(booking_model_handler_get.New(modelService))
	hf.RegisterUpdateHandler(booking_model_handler_update.New(modelService))
	hf.RegisterPatchHandler(booking_model_handler_update.New(modelService))
	return hf
}

func createShootingHandlerFinder(prefix string, shootingService booking_shooting_service.Service) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(booking_shooting_handler_list.New(shootingService))
	hf.RegisterCreateHandler(booking_shooting_handler_create.New(shootingService))
	hf.RegisterDeleteHandler(booking_shooting_handler_delete.New(shootingService))
	hf.RegisterGetHandler(booking_shooting_handler_get.New(shootingService))
	hf.RegisterUpdateHandler(booking_shooting_handler_update.New(shootingService))
	hf.RegisterPatchHandler(booking_shooting_handler_update.New(shootingService))
	hf.RegisterHandler("POST", "/book", booking_shooting_handler_book.New(shootingService))
	return hf
}

func createUserHandlerFinder(prefix string, userService booking_user_service.Service) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(booking_user_handler_list.New(userService))
	hf.RegisterCreateHandler(booking_user_handler_create.New(userService))
	hf.RegisterDeleteHandler(booking_user_handler_delete.New(userService))
	hf.RegisterGetHandler(booking_user_handler_get.New(userService))
	hf.RegisterUpdateHandler(booking_user_handler_update.New(userService))
	hf.RegisterPatchHandler(booking_user_handler_update.New(userService))
	hf.RegisterHandler("POST", "/verifyLogin", booking_user_handler_verifylogin.New(userService))
	return hf
}

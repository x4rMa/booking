package handler_configuration

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

	booking_shooting_handler_book "github.com/bborbe/booking/shooting/handler/book"
	booking_shooting_handler_create "github.com/bborbe/booking/shooting/handler/create"
	booking_shooting_handler_delete "github.com/bborbe/booking/shooting/handler/delete"
	booking_shooting_handler_get "github.com/bborbe/booking/shooting/handler/get"
	booking_shooting_handler_list "github.com/bborbe/booking/shooting/handler/list"
	booking_shooting_handler_update "github.com/bborbe/booking/shooting/handler/update"
	booking_shooting_service "github.com/bborbe/booking/shooting/service"

	booking_user_handler_create "github.com/bborbe/booking/user/handler/create"
	booking_user_handler_delete "github.com/bborbe/booking/user/handler/delete"
	booking_user_handler_get "github.com/bborbe/booking/user/handler/get"
	booking_user_handler_list "github.com/bborbe/booking/user/handler/list"
	booking_user_handler_update "github.com/bborbe/booking/user/handler/update"
	booking_user_service "github.com/bborbe/booking/user/service"

	booking_authentication_handler_verifylogin "github.com/bborbe/booking/authentication/handler/verifylogin"
	booking_authentication_service "github.com/bborbe/booking/authentication/service"

	booking_error_handler "github.com/bborbe/booking/error_handler"
)

var logger = log.DefaultLogger

func NewHandler(documentRoot string, dateService booking_date_service.Service, modelService booking_model_service.Service, shootingService booking_shooting_service.Service, userService booking_user_service.Service, authenticationService booking_authentication_service.Service) http.Handler {
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
	handlerFinder.RegisterHandlerFinder("/authentication", createAuthenticationHandlerFinder("/authentication", authenticationService))
	return log_handler.NewLogHandler(fallback.NewFallback(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", 404)))
}

func createAuthenticationHandlerFinder(prefix string, authenticationService booking_authentication_service.Service) handler_finder.HandlerFinder {
	hf := part.New(prefix)
	hf.RegisterHandler("/verifyLogin", booking_error_handler.New(booking_authentication_handler_verifylogin.New(authenticationService.VerifyLogin)))
	return hf
}

func createDateHandlerFinder(prefix string, dateService booking_date_service.Service) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(booking_error_handler.New(booking_date_handler_list.New(dateService.List)))
	hf.RegisterCreateHandler(booking_error_handler.New(booking_date_handler_create.New(dateService.Create)))
	hf.RegisterDeleteHandler(booking_error_handler.New(booking_date_handler_delete.New(dateService.Delete)))
	hf.RegisterGetHandler(booking_error_handler.New(booking_date_handler_get.New(dateService.Get)))
	hf.RegisterUpdateHandler(booking_error_handler.New(booking_date_handler_update.New(dateService.Update)))
	hf.RegisterPatchHandler(booking_error_handler.New(booking_date_handler_update.New(dateService.Update)))
	hf.RegisterHandler("GET", "/free", booking_error_handler.New(booking_date_handler_list.New(dateService.ListFree)))
	return hf
}

func createModelHandlerFinder(prefix string, modelService booking_model_service.Service) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(booking_error_handler.New(booking_model_handler_list.New(modelService.List, modelService.FindByToken)))
	hf.RegisterCreateHandler(booking_error_handler.New(booking_model_handler_create.New(modelService.Create)))
	hf.RegisterDeleteHandler(booking_error_handler.New(booking_model_handler_delete.New(modelService.Delete)))
	hf.RegisterGetHandler(booking_error_handler.New(booking_model_handler_get.New(modelService.Get)))
	hf.RegisterUpdateHandler(booking_error_handler.New(booking_model_handler_update.New(modelService.Update)))
	hf.RegisterPatchHandler(booking_error_handler.New(booking_model_handler_update.New(modelService.Update)))
	return hf
}

func createShootingHandlerFinder(prefix string, shootingService booking_shooting_service.Service) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(booking_error_handler.New(booking_shooting_handler_list.New(shootingService.List)))
	hf.RegisterCreateHandler(booking_error_handler.New(booking_shooting_handler_create.New(shootingService.Create)))
	hf.RegisterDeleteHandler(booking_error_handler.New(booking_shooting_handler_delete.New(shootingService.Delete)))
	hf.RegisterGetHandler(booking_error_handler.New(booking_shooting_handler_get.New(shootingService.Get)))
	hf.RegisterUpdateHandler(booking_error_handler.New(booking_shooting_handler_update.New(shootingService.Update)))
	hf.RegisterPatchHandler(booking_error_handler.New(booking_shooting_handler_update.New(shootingService.Update)))
	hf.RegisterHandler("POST", "/book", booking_error_handler.New(booking_shooting_handler_book.New(shootingService.Book)))
	return hf
}

func createUserHandlerFinder(prefix string, userService booking_user_service.Service) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(booking_error_handler.New(booking_user_handler_list.New(userService.List)))
	hf.RegisterCreateHandler(booking_error_handler.New(booking_user_handler_create.New(userService.Create)))
	hf.RegisterDeleteHandler(booking_error_handler.New(booking_user_handler_delete.New(userService.Delete)))
	hf.RegisterGetHandler(booking_error_handler.New(booking_user_handler_get.New(userService.Get)))
	hf.RegisterUpdateHandler(booking_error_handler.New(booking_user_handler_update.New(userService.Update)))
	hf.RegisterPatchHandler(booking_error_handler.New(booking_user_handler_update.New(userService.Update)))
	return hf
}

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

	booking_model_handler_complete "github.com/bborbe/booking/model/handler/complete"
	booking_model_handler_create "github.com/bborbe/booking/model/handler/create"
	booking_model_handler_current "github.com/bborbe/booking/model/handler/current"
	booking_model_handler_delete "github.com/bborbe/booking/model/handler/delete"
	booking_model_handler_get "github.com/bborbe/booking/model/handler/get"
	booking_model_handler_list "github.com/bborbe/booking/model/handler/list"
	booking_model_handler_update "github.com/bborbe/booking/model/handler/update"
	booking_model_service "github.com/bborbe/booking/model/service"

	booking_shooting_handler_book "github.com/bborbe/booking/shooting/handler/book"
	booking_shooting_handler_create "github.com/bborbe/booking/shooting/handler/create"
	booking_shooting_handler_current "github.com/bborbe/booking/shooting/handler/current"
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

	booking_authentication_converter "github.com/bborbe/booking/authentication/converter"
	booking_authentication_handler_verifylogin "github.com/bborbe/booking/authentication/handler/verifylogin"
	booking_authentication_service "github.com/bborbe/booking/authentication/service"

	booking_authorization "github.com/bborbe/booking/authorization"
	booking_authorization_service "github.com/bborbe/booking/authorization/service"
	booking_error_handler "github.com/bborbe/booking/error_handler"
	booking_handler "github.com/bborbe/booking/handler"
	booking_permission_check_handler "github.com/bborbe/booking/permission_check_handler"
)

var logger = log.DefaultLogger

type handlerConfiguration struct {
	documentRoot            string
	dateService             booking_date_service.Service
	modelService            booking_model_service.Service
	shootingService         booking_shooting_service.Service
	userService             booking_user_service.Service
	authenticationService   booking_authentication_service.Service
	authorizationService    booking_authorization_service.Service
	authenticationConverter booking_authentication_converter.Converter
}

func New(documentRoot string, dateService booking_date_service.Service, modelService booking_model_service.Service, shootingService booking_shooting_service.Service, userService booking_user_service.Service, authenticationService booking_authentication_service.Service, authorizationService booking_authorization_service.Service, authenticationConverter booking_authentication_converter.Converter) *handlerConfiguration {
	h := new(handlerConfiguration)
	h.documentRoot = documentRoot
	h.dateService = dateService
	h.modelService = modelService
	h.shootingService = shootingService
	h.userService = userService
	h.authenticationService = authenticationService
	h.authorizationService = authorizationService
	h.authenticationConverter = authenticationConverter
	return h
}

func (h *handlerConfiguration) GetHandler() http.Handler {
	logger.Debugf("root: %s", h.documentRoot)
	fileServer := cachingheader.NewCachingHeaderHandler(contenttype.NewContentTypeHandler(http.FileServer(http.Dir(h.documentRoot))))
	handlerFinder := part.New("")
	handlerFinder.RegisterHandler("/", fileServer)
	handlerFinder.RegisterHandler("/css", fileServer)
	handlerFinder.RegisterHandler("/js", fileServer)
	handlerFinder.RegisterHandler("/images", fileServer)
	handlerFinder.RegisterHandlerFinder("/date", h.createDateHandlerFinder("/date"))
	handlerFinder.RegisterHandlerFinder("/model", h.createModelHandlerFinder("/model"))
	handlerFinder.RegisterHandlerFinder("/shooting", h.createShootingHandlerFinder("/shooting"))
	handlerFinder.RegisterHandlerFinder("/user", h.createUserHandlerFinder("/user"))
	handlerFinder.RegisterHandlerFinder("/authentication", h.createAuthenticationHandlerFinder("/authentication"))
	return log_handler.NewLogHandler(fallback.NewFallback(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", 404)))
}

func (h *handlerConfiguration) createAuthenticationHandlerFinder(prefix string) handler_finder.HandlerFinder {
	hf := part.New(prefix)
	hf.RegisterHandler("/verifyLogin", h.handle_errors(h.check_permission(booking_authentication_handler_verifylogin.New(h.authenticationService.VerifyLogin), booking_authorization.None)))
	return hf
}

func (h *handlerConfiguration) createDateHandlerFinder(prefix string) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(h.handle_errors(h.check_permission(booking_date_handler_list.New(h.dateService.List), booking_authorization.Organizer)))
	hf.RegisterCreateHandler(h.handle_errors(h.check_permission(booking_date_handler_create.New(h.dateService.Create), booking_authorization.Organizer)))
	hf.RegisterDeleteHandler(h.handle_errors(h.check_permission(booking_date_handler_delete.New(h.dateService.Delete), booking_authorization.Organizer)))
	hf.RegisterGetHandler(h.handle_errors(h.check_permission(booking_date_handler_get.New(h.dateService.Get), booking_authorization.Participant)))
	hf.RegisterUpdateHandler(h.handle_errors(h.check_permission(booking_date_handler_update.New(h.dateService.Update), booking_authorization.Organizer)))
	hf.RegisterPatchHandler(h.handle_errors(h.check_permission(booking_date_handler_update.New(h.dateService.Update), booking_authorization.Organizer)))
	hf.RegisterHandler("GET", "/free", h.handle_errors(h.check_permission(booking_date_handler_list.New(h.dateService.ListFree), booking_authorization.Participant)))
	return hf
}

func (h *handlerConfiguration) createModelHandlerFinder(prefix string) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(h.handle_errors(h.check_permission(booking_model_handler_list.New(h.modelService.List), booking_authorization.Organizer)))
	hf.RegisterCreateHandler(h.handle_errors(h.check_permission(booking_model_handler_create.New(h.modelService.Create), booking_authorization.Organizer)))
	hf.RegisterDeleteHandler(h.handle_errors(h.check_permission(booking_model_handler_delete.New(h.modelService.Delete), booking_authorization.Organizer)))
	hf.RegisterGetHandler(h.handle_errors(h.check_permission(booking_model_handler_get.New(h.modelService.Get), booking_authorization.Organizer)))
	hf.RegisterUpdateHandler(h.handle_errors(h.check_permission(booking_model_handler_update.New(h.modelService.Update), booking_authorization.Organizer)))
	hf.RegisterPatchHandler(h.handle_errors(h.check_permission(booking_model_handler_update.New(h.modelService.Update), booking_authorization.Organizer)))
	hf.RegisterHandler("GET", "/current", h.handle_errors(h.check_permission(booking_model_handler_current.New(h.authenticationConverter.HttpRequestToAuthentication, h.modelService.GetByToken), booking_authorization.Participant)))
	hf.RegisterHandler("PUT", "/complete", h.handle_errors(h.check_permission(booking_model_handler_complete.New(h.modelService.Update), booking_authorization.Participant)))
	return hf
}

func (h *handlerConfiguration) createShootingHandlerFinder(prefix string) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(h.handle_errors(h.check_permission(booking_shooting_handler_list.New(h.shootingService.List), booking_authorization.Organizer)))
	hf.RegisterCreateHandler(h.handle_errors(h.check_permission(booking_shooting_handler_create.New(h.shootingService.Create), booking_authorization.Organizer)))
	hf.RegisterDeleteHandler(h.handle_errors(h.check_permission(booking_shooting_handler_delete.New(h.shootingService.Delete), booking_authorization.Organizer)))
	hf.RegisterGetHandler(h.handle_errors(h.check_permission(booking_shooting_handler_get.New(h.shootingService.Get), booking_authorization.Participant)))
	hf.RegisterUpdateHandler(h.handle_errors(h.check_permission(booking_shooting_handler_update.New(h.shootingService.Update), booking_authorization.Organizer)))
	hf.RegisterPatchHandler(h.handle_errors(h.check_permission(booking_shooting_handler_update.New(h.shootingService.Update), booking_authorization.Organizer)))
	hf.RegisterHandler("POST", "/book", h.handle_errors(h.check_permission(booking_shooting_handler_book.New(h.authenticationConverter.HttpRequestToAuthentication, h.modelService.GetByToken, h.shootingService.Book, h.shootingService.Get), booking_authorization.Participant)))
	hf.RegisterHandler("GET", "/current", h.handle_errors(h.check_permission(booking_shooting_handler_current.New(h.authenticationConverter.HttpRequestToAuthentication, h.modelService.GetByToken, h.shootingService.FindByModelId), booking_authorization.Participant)))
	return hf
}

func (h *handlerConfiguration) createUserHandlerFinder(prefix string) handler_finder.HandlerFinder {
	hf := rest.New(prefix)
	hf.RegisterListHandler(h.handle_errors(h.check_permission(booking_user_handler_list.New(h.userService.List), booking_authorization.Administrator)))
	hf.RegisterCreateHandler(h.handle_errors(h.check_permission(booking_user_handler_create.New(h.userService.Create), booking_authorization.Administrator)))
	hf.RegisterDeleteHandler(h.handle_errors(h.check_permission(booking_user_handler_delete.New(h.userService.Delete), booking_authorization.Administrator)))
	hf.RegisterGetHandler(h.handle_errors(h.check_permission(booking_user_handler_get.New(h.userService.Get), booking_authorization.Administrator)))
	hf.RegisterUpdateHandler(h.handle_errors(h.check_permission(booking_user_handler_update.New(h.userService.Update), booking_authorization.Administrator)))
	hf.RegisterPatchHandler(h.handle_errors(h.check_permission(booking_user_handler_update.New(h.userService.Update), booking_authorization.Administrator)))
	return hf
}

func (h *handlerConfiguration) check_permission(handler booking_handler.Handler, role booking_authorization.Role) booking_handler.Handler {
	return booking_permission_check_handler.New(h.authorizationService.HasRole, h.authenticationConverter.HttpRequestToAuthentication, role, handler)
}

func (h *handlerConfiguration) handle_errors(handler booking_handler.Handler) http.Handler {
	return booking_error_handler.New(handler)
}

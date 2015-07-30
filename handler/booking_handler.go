package handler

import (
	"net/http"

	user_service "github.com/bborbe/booking/user/service"
	"github.com/bborbe/log"
	"github.com/bborbe/server/handler/cachingheader"
	"github.com/bborbe/server/handler/contenttype"
	"github.com/bborbe/server/handler/fallback"
	json_handler "github.com/bborbe/server/handler/json"
	log_handler "github.com/bborbe/server/handler/log"
	"github.com/bborbe/server/handler/static"
	"github.com/bborbe/server/handler_finder/part"
)

var logger = log.DefaultLogger

func NewHandler(documentRoot string) http.Handler {
	logger.Debugf("root: %s", documentRoot)
	fileServer := cachingheader.NewCachingHeaderHandler(contenttype.NewContentTypeHandler(http.FileServer(http.Dir(documentRoot))))
	handlerFinder := part.New("")
	handlerFinder.RegisterHandler("/", fileServer)
	handlerFinder.RegisterHandler("/css", fileServer)
	handlerFinder.RegisterHandler("/js", fileServer)
	handlerFinder.RegisterHandler("/images", fileServer)
	handlerFinder.RegisterHandler("/users.json", createRestHandler())
	return log_handler.NewLogHandler(fallback.NewFallback(handlerFinder, static.NewHandlerStaticContentReturnCode("not found", 404)))
}

func createRestHandler() http.Handler {
	userService := user_service.NewUserService()
	users := userService.List()
	return json_handler.NewJsonHandler(users)
}

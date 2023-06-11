package routers

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

func initValidator() {
	if _, ok := binding.Validator.Engine().(*validator.Validate); ok {

	} else {
		log.Warn("@routers.initValidator: Can not init validator")
	}
}

package model

import (
	"completerr/utils"
	"github.com/reugn/go-quartz/quartz"
)

var logger = utils.GetLogger()

type CompleterrJob interface {
	quartz.Job
	Name() string
}

package console

import (
	"github.com/bbcloudGroup/gothic/cli"
	"github.com/bbcloudGroup/gothic/di"
	"github.com/robfig/cron/v3"
)

func Schedule(cron *cron.Cron) {

	//di.Invoke(func (printer Printer) { _, _ = cron.AddJob("* * * * * *", printer) })
	//di.Invoke(func (printer Printer) {cli.Once(printer, true)})
	//di.Invoke(func (printer Printer) {cli.Forever(printer)})

	di.Invoke(func (migrate AdminMigrate) {cli.Once(migrate, false)})
}



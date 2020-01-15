package getting_started

import (
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var OneMinuteQuickStart = Components(
	md.Markdown(`
This article try to let you use the shortest time to get a taste of how powerful GoPlaid is.

One of the GoPlaid module called presents that can quickly create admin interface like [these](/samples/presets-detail-page-cards/customers):

Install the command line tool with:

~~~
$ go get -v github.com/goplaid/x/goplaid
~~~

And run:

~~~
$ goplaid
~~~

It will promote you to input a Go package, and create the admin app in current directory.

Change to the created package directory, and use ~docker-compose up~ to start the database, and then
Use a new terminal to run ~source dev_env && go run main.go~ to start the admin app

`),
)

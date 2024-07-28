package main

import "EffectiveMobile/internal/app"

// @title           UwU
// @version         1.0
// @description     This is my server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache helicopter
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	app.Run()
}

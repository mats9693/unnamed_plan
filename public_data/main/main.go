package main

import (
	_ "github.com/mats9693/unnamed_plan/public_data/http"
	_ "github.com/mats9693/utils/toy_server/config"
	_ "github.com/mats9693/utils/toy_server/db"

	"github.com/mats9693/utils/toy_server/http"
)

func main() {
	shttp.StartServer()
}

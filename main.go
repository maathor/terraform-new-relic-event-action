package main

import (
	"fmt"
	"github.com/newrelic/go-agent"
	"log"
	"os"
	"time"
)

var newRelicApp newrelic.Application

func main() {
	newRelicEventType := os.Getenv("INPUT_EVENT_TYPE_NAME")
	newRelicLicenseKey := os.Getenv("INPUT_NEW_RELIC_LICENCE_KEY")
	stage := os.Getenv("INPUT_ENV")
	terraformInitStatus := os.Getenv("INPUT_TERRAFORM_INIT_STATUS")
	terraformApplyStatus := os.Getenv("INPUT_TERRAFORM_APPLY_STATUS")
	terraformTagKey := os.Getenv("INPUT_TERRAFORM_TAG_KEY")
	terraformTagValue := os.Getenv("INPUT_TERRAFORM_TAG_VALUE")
	github_repository := os.Getenv("INPUT_GITHUB_REPOSITORY")
	github_run_id := os.Getenv("INPUT_GITHUB_RUN_ID")
	gha_url := fmt.Sprintf("https://github.com/%s/actions/runs/%s", github_repository, github_run_id)

	app := initNewRelicClient(newRelicLicenseKey)

	if err := app.RecordCustomEvent(newRelicEventType,map[string]interface{}{
		"env": stage,
		"terraformApply": terraformApplyStatus,
		terraformTagKey: terraformTagValue,
		"terraformInit": terraformInitStatus,
		"ghaUrl": gha_url,
	}); err != nil {
		log.Println("error in creating New Relic custom event: ", err)
		return
	}
	app.Shutdown(5 * time.Second)

}

func initNewRelicClient(newRelicLicenseKey string) newrelic.Application {
	config := newrelic.NewConfig("GithubActionsMonitoring", newRelicLicenseKey)
	newRelicApp, err := newrelic.NewApplication(config)
	if err != nil {
		log.Println("error in creating new relic application: ", err)
		panic(err)
	}
	newRelicApp.WaitForConnection(5 * time.Second)
	return newRelicApp
}
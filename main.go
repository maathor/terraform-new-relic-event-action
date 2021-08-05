package main

import (
	"bufio"
	"fmt"
	"github.com/newrelic/go-agent"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var newRelicApp newrelic.Application

type TerraformOperations struct {
	Create int
	NoOp int
	Delete int
	Update int
	Read int
}

func main() {
	//inputs
	newRelicEventType := os.Getenv("INPUT_EVENT_TYPE_NAME")
	newRelicLicenseKey := os.Getenv("INPUT_NEW_RELIC_LICENCE_KEY")
	stage := os.Getenv("INPUT_ENV")
	terraformInitStatus := os.Getenv("INPUT_TERRAFORM_INIT_STATUS")
	terraformApplyStatus := os.Getenv("INPUT_TERRAFORM_APPLY_STATUS")
	terraformTagKey := os.Getenv("INPUT_TERRAFORM_TAG_KEY")
	terraformTagValue := os.Getenv("INPUT_TERRAFORM_TAG_VALUE")
	github_repository := os.Getenv("INPUT_GITHUB_REPOSITORY")
	github_run_id := os.Getenv("INPUT_GITHUB_RUN_ID")
	terraformOperationPath := os.Getenv("INPUT_TERRAFORM_OPERATION_LIST_PATH")
	gha_url := fmt.Sprintf("https://github.com/%s/actions/runs/%s", github_repository, github_run_id)

	app := initNewRelicClient(newRelicLicenseKey)
	byteFile, err := ioutil.ReadFile(terraformOperationPath)
	if err != nil {
		log.Println("error while reading terraform operation file: ", err)
	}
	terraformOperations := computeTerraformOperationsNumber(string(byteFile))

	if err := app.RecordCustomEvent(newRelicEventType,map[string]interface{}{
		"env": stage,
		"terraformApply": terraformApplyStatus,
		terraformTagKey: terraformTagValue,
		"terraformInit": terraformInitStatus,
		"terraformCreate" : terraformOperations.Create,
		"terraformDelete" : terraformOperations.Delete,
		"terraformNoOp" : terraformOperations.NoOp,
		"terraformUpdate" : terraformOperations.Update,
		"ghaUrl": gha_url,
	}); err != nil {
		log.Println("error in creating New Relic custom event: ", err)
		return
	}
	app.Shutdown(5 * time.Second)
	// output
	fmt.Println(fmt.Sprintf(`::set-output name=terraform_update::%s`, strconv.Itoa(terraformOperations.Update)))
	fmt.Println(fmt.Sprintf(`::set-output name=terraform_create::%s`, strconv.Itoa(terraformOperations.Create)))
	fmt.Println(fmt.Sprintf(`::set-output name=terraform_delete::%s`, strconv.Itoa(terraformOperations.Delete)))
	fmt.Println(fmt.Sprintf(`::set-output name=terraform_noop::%s`, strconv.Itoa(terraformOperations.NoOp)))
}

func computeTerraformOperationsNumber(terraformOperation string) TerraformOperations {
	var tOp TerraformOperations
	scanner := bufio.NewScanner(strings.NewReader(terraformOperation))
	for scanner.Scan() {
		operation := scanner.Text()
		switch {
		case strings.Contains(operation, "no-op"):
			tOp.NoOp = tOp.NoOp + 1
		case strings.Contains(operation, "create"):
			tOp.Create = tOp.Create + 1
		case strings.Contains(operation, "delete"):
			tOp.Create = tOp.Create + 1
		case strings.Contains(operation, "read"):
			tOp.Read = tOp.Read + 1
		case strings.Contains(operation, "update"):
			tOp.Update = tOp.Update + 1
		default:
			log.Println("Operation Unknown: ", operation)
		}
	}
	return tOp
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
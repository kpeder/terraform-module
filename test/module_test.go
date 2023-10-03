package test

import (
	"flag"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Flag to destroy the target environment after tests
var destroy = flag.Bool("destroy", false, "destroy environment after tests")

func TestExample1(t *testing.T) {
	// Set execution directory
	terraformOptions := &terraform.Options{
		TerraformDir: "../fixtures/example1",
	}

	// Check for inputs file
	if !assert.FileExists(t, terraformOptions.TerraformDir+"/inputs.yaml") {
		t.Fail()
	}

	// Read the inputs.yaml
	yfile, err := os.ReadFile(terraformOptions.TerraformDir + "/inputs.yaml")
	if err != nil {
		t.Fail()
	}

	inputs := make(map[string]interface{})
	err = yaml.Unmarshal(yfile, &inputs)
	if err != nil {
		t.Fail()
	}

	// Defer Terraform destroy only if flag is set
	if *destroy {
		defer terraform.Destroy(t, terraformOptions)
	}

	// Initialize the deployment and create resources
	terraform.InitAndApply(t, terraformOptions)

	outputValue := terraform.Output(t, terraformOptions, "random_pet")

	// Test for valid output
	if assert.NotNil(t, outputValue) {
		t.Logf("Output test PASSED. Expected output to be string, got %s", outputValue)
	} else {
		t.Error("Output test FAILED. Expected output to be string, got nil")
	}
	if inputs["prefix"] != nil {
		// Test for prefix
		if assert.Equal(t, strings.Split(outputValue, "-")[0], inputs["prefix"].(string)) {
			t.Logf("Prefix test PASSED. Expected output to start with %s, got %s", inputs["prefix"].(string), strings.Split(outputValue, "-")[0])
		} else {
			t.Errorf("Prefix test FAILED. Expected output to start with %s, got %s", inputs["prefix"].(string), strings.Split(outputValue, "-")[0])
		}
		// Test for word count
		if assert.Equal(t, len(strings.Split(outputValue, "-")), inputs["length"].(int)+1) {
			t.Logf("Word count test PASSED. Expected output to contain prefix plus %d more words, got %d", inputs["length"].(int), len(strings.Split(outputValue, "-")))
		} else {
			t.Errorf("Word count test FAILED. Expected output to contain prefix plus %d more words, got %d", inputs["length"].(int), len(strings.Split(outputValue, "-")))
		}
	} else {
		// Test for word count
		if assert.Equal(t, len(strings.Split(outputValue, "-")), inputs["length"].(int)) {
			t.Logf("Word count test PASSED. Expected output to contain %d words, got %d", inputs["length"].(int), len(strings.Split(outputValue, "-")))
		} else {
			t.Errorf("Word count test FAILED. Expected output to contain %d words, got %d", inputs["length"].(int), len(strings.Split(outputValue, "-")))
		}
	}
}

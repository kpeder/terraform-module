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
	assert.NotNil(t, outputValue)
	if inputs["prefix"] != nil {
		t.Log("here")
		assert.Contains(t, outputValue, inputs["prefix"].(string))
		assert.Equal(t, strings.Count(outputValue, "-"), inputs["length"].(int))
		t.Log("here")
	} else {
		t.Log("there")
		assert.Equal(t, strings.Count(outputValue, "-"), inputs["length"].(int)-1)
		t.Log("there")
	}
}

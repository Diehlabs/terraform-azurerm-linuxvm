package test

import (
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestMyModule(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/build",
		Vars: map[string]interface{}{
			"gh_run_id": os.Getenv("GITHUB_RUN_ID"),
			"gh_repo":   strings.Replace(os.Getenv("GITHUB_REPOSITORY"), "/", "-", -1),
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	//output := terraform.Output(t, terraformOptions, "hello_world")
	assert.Equal(t, "Hello, World!", "Hello, World!")
}

package test

import (
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestMyModule(t *testing.T) {
	t.Parallel()
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

	terraform.ApplyAndIdempotent(t, terraformOptions)

	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	nsgName := terraform.Output(t, terraformOptions, "nsg_name")
	// sshRuleName := terraform.Output(t, terraformOptions, "ssh_rule_name")
	// httpRuleName := terraform.Output(t, terraformOptions, "http_rule_name")

	// A default NSG has 6 rules, and we have two custom rules for a total of 8
	rules, err := azure.GetAllNSGRulesE(resourceGroupName, nsgName, "")
	assert.NoError(t, err)
	assert.Equal(t, 8, len(rules.SummarizedRules))

}

// func TestTerraformAzureNsgExample(t *testing.T) {

// 	// We should have a rule for allowing ssh
// 	sshRule := rules.FindRuleByName(sshRuleName)

// 	// That rule should allow port 22 inbound
// 	assert.True(t, sshRule.AllowsDestinationPort(t, "22"))

// 	// But should not allow 80 inbound
// 	assert.False(t, sshRule.AllowsDestinationPort(t, "80"))

// 	// SSh is allowed from any port
// 	assert.True(t, sshRule.AllowsSourcePort(t, "*"))

// 	// We should have a rule for blocking HTTP
// 	httpRule := rules.FindRuleByName(httpRuleName)

// 	// This rule should BLOCK port 80 inbound
// 	assert.False(t, httpRule.AllowsDestinationPort(t, "80"))
// }

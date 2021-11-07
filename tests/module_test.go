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
	os.Setenv("AZURE_TENANT_ID", os.Getenv("ARM_TENANT_ID"))
	os.Setenv("AZURE_CLIENT_ID", os.Getenv("ARM_CLIENT_ID"))
	os.Setenv("AZURE_CLIENT_SECRET", os.Getenv("ARM_CLIENT_SECRET"))
	os.Setenv("AZURE_SUBSCRIPTION_ID", os.Getenv("ARM_SUBSCRIPTION_ID"))

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

	// terraform.ApplyAndIdempotent(t, terraformOptions)

	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	nsgName := terraform.Output(t, terraformOptions, "nsg_name")
	// sshRuleName := terraform.Output(t, terraformOptions, "ssh_rule_name")
	// httpsRuleName := terraform.Output(t, terraformOptions, "https_rule_name")

	// A default NSG has 6 rules, and we have two custom rules for a total of 8
	rules, err := azure.GetAllNSGRulesE(resourceGroupName, nsgName, "")
	assert.NoError(t, err)
	assert.Equal(t, 8, len(rules.SummarizedRules))

	// ssh rule is contained in the module
	sshRule := rules.FindRuleByName("SSH")
	assert.True(t, sshRule.AllowsDestinationPort(t, "22"))
	assert.False(t, sshRule.AllowsDestinationPort(t, "80"))
	assert.True(t, sshRule.AllowsSourcePort(t, "*"))

	// https rule is passed to the module
	httpsRule := rules.FindRuleByName("HTTPS")
	assert.True(t, httpsRule.AllowsDestinationPort(t, "443"))
	assert.False(t, httpsRule.AllowsDestinationPort(t, "80"))
	assert.True(t, httpsRule.AllowsSourcePort(t, "*"))
}

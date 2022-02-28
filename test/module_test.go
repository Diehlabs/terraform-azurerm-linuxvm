package test

import (
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

type RunSettings struct {
	t                         *testing.T
	workingDir                string `default:"../examples/build"`
	tfCliPath                 string `default:"/usr/local/bin/terraform"`
	approleID                 string
	// secretID                  *auth.SecretID
	vaultSecretPath           string
	uniqueID                  string
}

func (r *RunSettings) setDefaults() {
	if r.t == nil {
		panic("No Terratest module provided")
	}

	r.workingDir = "../examples/build"
	if ttwd := os.Getenv("TERRATEST_WORKING_DIR"); ttwd != "" {
		r.workingDir = ttwd
	}

	r.tfCliPath = "/usr/local/bin/terraform"
	if tfcp := os.Getenv("TF_CLI_PATH"); tfcp != "" {
		r.tfCliPath = tfcp
	}

	// VAULT items are used for interaction with HashiCorp Vault
	if vsecp := os.Getenv("VAULT_SECRET_PATH"); vsecp != "" {
		r.vaultSecretPath = vsecp
	}

	if role_id := os.Getenv("VAULT_APPROLE_ID"); role_id != "" {
		r.approleID = role_id
	}

	// if wrapped_token := os.Getenv("VAULT_WRAPPED_TOKEN"); wrapped_token != "" {
	// 	r.secretID = &auth.SecretID{FromEnv: "VAULT_WRAPPED_TOKEN"}
	// }

	if localId := os.Getenv("GITHUB_RUN_ID"); localId != "" {
		r.uniqueID = localId
	} else {
		r.uniqueID = random.UniqueId()
	}

}

func (r *RunSettings) setTerraformOptions() {
	repoName := "local-repo"
	if ghRepo := os.Getenv("GITHUB_REPOSITORY"); ghRepo != "" {
		repoName = strings.Replace(ghRepo, "/", "-", -1)
	}
	

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(r.t, &terraform.Options{
		TerraformDir:    r.workingDir,
		TerraformBinary: r.tfCliPath,
		Vars: map[string]interface{}{
			"gh_run_id": r.uniqueID,
			"gh_repo":   repoName,
		},
	})

	test_structure.SaveTerraformOptions(r.t, r.workingDir, terraformOptions)
}

func TestMyModule(t *testing.T) {
	r := RunSettings{t: t}
	r.setDefaults()

	t.Parallel()

	// to set terraform options
	// to skip execution "export SKIP_set_terraformOptions=true" in terminal
	test_structure.RunTestStage(r.t, "setTerraformOptions", r.setTerraformOptions)

	// Destroy the infra after testing is finished
	// export SKIP_terraformDestroy=true to skip this stage
	defer test_structure.RunTestStage(r.t, "terraformDestroy", r.terraformDestroy)

	// Deploy using Terraform
	// export SKIP_deployTerraform=true to skip this stage
	test_structure.RunTestStage(r.t, "deployTerraform", r.deployUsingTerraform)

		// Redeploy using Terraform and ensure idempotency
	// export SKIP_redeployTerraform=true to skip this stage
	test_structure.RunTestStage(r.t, "redeployTerraform", r.redeployUsingTerraform)

	// Perform tests
	// export SKIP_runTests=true to skip this stage
	test_structure.RunTestStage(r.t, "runTests", r.runTests)

}

func (r *RunSettings) deployUsingTerraform() {
	terraformOptions := test_structure.LoadTerraformOptions(r.t, r.workingDir)
	terraform.InitAndApply(r.t, terraformOptions)
}

func (r *RunSettings) redeployUsingTerraform() {
	terraformOptions := test_structure.LoadTerraformOptions(r.t, r.workingDir)
	terraform.ApplyAndIdempotent(r.t, terraformOptions)
}

func (r *RunSettings) runTests() {
	// start - temp work
	// set xtra vars - will be replaced with go module to pull all secrets from vault
	os.Setenv("AZURE_TENANT_ID", os.Getenv("ARM_TENANT_ID"))
	os.Setenv("AZURE_CLIENT_ID", os.Getenv("ARM_CLIENT_ID"))
	os.Setenv("AZURE_CLIENT_SECRET", os.Getenv("ARM_CLIENT_SECRET"))
	os.Setenv("AZURE_SUBSCRIPTION_ID", os.Getenv("ARM_SUBSCRIPTION_ID"))
	// end - temp work

	terraformOptions := test_structure.LoadTerraformOptions(r.t, r.workingDir)

	resourceGroupName := terraform.Output(r.t, terraformOptions, "resource_group_name")
	nsgName := terraform.Output(r.t, terraformOptions, "nsg_name")
	// sshRuleName := terraform.Output(r.t, terraformOptions, "ssh_rule_name")
	// httpsRuleName := terraform.Output(r.t, terraformOptions, "https_rule_name")

	// A default NSG has 6 rules, and we have two custom rules for a total of 8
	rules, err := azure.GetAllNSGRulesE(resourceGroupName, nsgName, "")
	assert.NoError(r.t, err)
	assert.Equal(r.t, 8, len(rules.SummarizedRules))

	// ssh rule is contained in the module
	sshRule := rules.FindRuleByName("SSH")
	assert.True(r.t, sshRule.AllowsDestinationPort(r.t, "22"))
	assert.False(r.t, sshRule.AllowsDestinationPort(r.t, "80"))
	assert.True(r.t, sshRule.AllowsSourcePort(r.t, "*"))

	// https rule is passed to the module
	httpsRule := rules.FindRuleByName("HTTPS")
	assert.True(r.t, httpsRule.AllowsDestinationPort(r.t, "443"))
	assert.False(r.t, httpsRule.AllowsDestinationPort(r.t, "80"))
	assert.True(r.t, httpsRule.AllowsSourcePort(r.t, "*"))
}

func (r *RunSettings) terraformDestroy() {
	terraformOptions := test_structure.LoadTerraformOptions(r.t, r.workingDir)
	terraform.Destroy(r.t, terraformOptions)
}
package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

var uniqueId = random.UniqueId()

var terraformBinary = "/usr/local/bin/terraform"

var workingDir = "../examples/build"


func TestTerraformModule(t *testing.T) {
	t.Parallel()

	//os.Setenv("SKIP_reterraform_deploy", "true")
	//os.Setenv("SKIP_terraform_redeploy", "true")
	//os.Setenv("SKIP_terraform_destroy", "true")

	if tfbin := os.Getenv("TF_CLI_PATH"); tfbin != "" {
		terraformBinary = tfbin
	}

	if tfdir := os.Getenv("TERRATEST_WORKING_DIR"); tfdir != "" {
		workingDir = tfdir
	}

	terraformVars := map[string]interface{}{
		"unique_id": uniqueId,
		"test_for": fmt.Sprintf("terratest-local-%s", uniqueId),
	}

	setupTesting(t, workingDir, terraformBinary, terraformVars)

	// Destroy the infra after testing is finished
	defer test_structure.RunTestStage(t, "terraform_destroy", func(){
		terraform_destroy(t, workingDir)
	})

	// Deploy using Terraform
	test_structure.RunTestStage(t, "terraform_deploy", func() {
		deployUsingTerraform(t, workingDir)
	})

	// Redeploy using Terraform and ensure idempotency
	test_structure.RunTestStage(t, "terraform_redeploy", func() {
		redeployUsingTerraform(t, workingDir)
	})

	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	t.Run("Ensure NSG rules match inputs", func(t *testing.T){
		testNsgRules(t, terraformOptions, workingDir, resourceGroupName)
	})

	t.Run("Ensure VM size matches input", func(t *testing.T){
		testVmSize(t, terraformOptions, resourceGroupName)
	})

}

func setEnvVars() {
	os.Setenv("AZURE_TENANT_ID", os.Getenv("ARM_TENANT_ID"))
	os.Setenv("AZURE_CLIENT_ID", os.Getenv("ARM_CLIENT_ID"))
	os.Setenv("AZURE_CLIENT_SECRET", os.Getenv("ARM_CLIENT_SECRET"))
	os.Setenv("AZURE_SUBSCRIPTION_ID", os.Getenv("ARM_SUBSCRIPTION_ID"))
}

func testNsgRules(t *testing.T, terraformOptions *terraform.Options, workingDir string, resourceGroupName string) {
	setEnvVars()

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

func testVmSize(t *testing.T, terraformOptions *terraform.Options, resourceGroupName string) {
	setEnvVars()

	virtualMachineName := terraform.Output(t, terraformOptions, "vm_name")

	expectedVMSize := compute.VirtualMachineSizeTypes(terraform.Output(t, terraformOptions, "vm_size"))

	vmByRef := azure.GetVirtualMachine(t, virtualMachineName, resourceGroupName, os.Getenv("AZURE_SUBSCRIPTION_ID"))
	vmInstance := azure.Instance{VirtualMachine: vmByRef}
	actualVMSize := vmInstance.GetVirtualMachineInstanceSize()
	assert.Equal(t, expectedVMSize, actualVMSize)
}

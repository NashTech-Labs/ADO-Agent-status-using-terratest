package main

import (
	"encoding/json"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
    "github.com/gruntwork-io/terratest/modules/terraform"

)
type Agent struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type AgentResponse struct {
	Agents []Agent `json:"value"`
}


var (
	organizationURL     = "< >" 
	organizationName    = "< >"
	expectedStatus      = "online"
	poolID = < >
)



func TestAzureVMWithLogic(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../module",
		VarFiles:     []string{"../test/terraform.tfvars"},
		Vars: map[string]interface{}{
			"vm_size": "Standard_DS1_v2",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)


	personalAccessToken := os.Getenv("TF_VAR_adotoken")
	// fmt.Println("Personal Access Token for the ADO:-",personalAccessToken)
	if personalAccessToken == "" {
		t.Fatal("TF_VAR_token environment variable is not set")
	}

	agentNames := terraform.OutputList(t, terraformOptions, "agent_name")
	for _, name := range agentNames {
		// fmt.Println(name)

		agent, err := getAgent(organizationURL, personalAccessToken, name)
		if err != nil {
			assert.Fail(t, "Failed to retrieve agent:", err.Error())
			continue
		}
		
		assert.NotNil(t, agent, "Agent not found")
		t.Run(fmt.Sprintf("Agent status has been matched for agent: %s", name), func(t *testing.T) {
			// Check the agent status matches the expected status
			assert.Equal(t, expectedStatus, agent.Status, "Agent status mismatch")
   		 })
	}


}




func getAgent(organizationURL, personalAccessToken, agentName string) (*Agent, error) {
	client := &http.Client{}

	url := fmt.Sprintf("%s/%s/_apis/distributedtask/pools/%d/agents?api-version=6.1-preview.1", organizationURL, organizationName, poolID)
	// fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	token := base64.StdEncoding.EncodeToString([]byte(":" + personalAccessToken))
	req.Header.Set("Authorization", "Basic "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve agents. Status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var agentResponse AgentResponse
	err = json.Unmarshal(body, &agentResponse)
	if err != nil {
		return nil, err
	}

	for _, agent := range agentResponse.Agents {
		if agent.Name == agentName {
			return &agent, nil
		}
	}

	return nil, fmt.Errorf("agent '%s' not found", agentName)
}


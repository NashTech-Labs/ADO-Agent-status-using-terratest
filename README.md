# Test the ADO-Agent using terratest.

 

## This repo contains the terratest test cases for the ADO-Agent.

 


-------------
### To run this terratest, You must have the terraform ADO-Agent module and should be in the same root and you need to export the given variale in your environment and then run go test.

 


1. Export the ADO token as

 

        export TF_VAR_token="<>"

 

2. You have to export the credential of azure.

 

        export CLIENT_ID=""

 

        export CLIENT_SECRET=""

 

        export TENANT_ID=""

 

        export SUBSCRIPTION_ID=""%           

3. You need to add the following values this code:

        i.   Add your project namd

        ii.  Add your organisation name
        
        iii. Add your pool id

4. In the last, you need to run the below command to run the test case:-

             go mod init <>

             go mod tidy

             go test -v
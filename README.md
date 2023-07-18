# Test the ADO-Agent using terratest.

 

## This repo contains the terratest test cases for the ADO-Agent.

 


-------------
### To run this terratest, You must have the terraform vm module and should be in the same root and you need to export the given variale in your environment and then run go test.

 


1. Export the ADO token as

 

        export TF_VAR_token="<>"

 

2. You have to export the credential of azure.

 

        export ARM_CLIENT_ID=""

 

        export ARM_CLIENT_SECRET=""

 

        export ARM_TENANT_ID=""

 

        export ARM_SUBSCRIPTION_ID=""%           

3. You need to add the following values this code:

        i.   Add your project namd

        ii.  Add your organisation name
        
        iii. Add your pool id
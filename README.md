**USECASE**

API for finding stores across a given country

**FUNCTIONALITY**

This is a lambda function with the API functionality that caters to all CRUD operations with dynamodb

**Prerequisites**

- Add following env variables
   
  - ```APP_DEUBG_ENABLED``` -  to true if you wanted debug logs
  - APP_GO_ENV -  to **DEV** if you wanted to taken in environment variables from local dev environment. Else by default it will pickup env variables from aws SSM
- Use ```serverless local``` to run the code in local
- To run in a lambda environment make sure to deploy the application to you aws using ```serverless deploy```
  
  

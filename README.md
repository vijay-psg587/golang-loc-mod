**USECASE**

CRUD API to describe the dynamodb usage with golang

**FUNCTIONALITY**

- ```gotdotenv``` is used for env config assignment in local
- Single pattern is implemented with mutex locking
- ```ZeroLog``` is used as a logging framework
- ```common``` module is implemented which can be generalized and used in other microservices module
- It uses ```Go Fiber``` for making the API calls
- CORS along with logging middleware is implemented using ```go routines```. GO routine can be extended if needed be
- Custom Http Error handling is implemented
- Code is modularized as much as possible (Possibility of over engineering)
- Uses AWS DynamoDB as the database server
- Proper recover stratergy is suggested incase of panics in Handler of the API
- Context is used to set the locals and values are being passed arnd if needed be. Middleware pattern is being used

 
**Prerequisites**

- Add following env variables
   
  - ```APP_DEUBG_ENABLED``` -  to true if you wanted debug logs
  - APP_GO_ENV -  to **dev** if you wanted to taken in environment variables from local dev environment. Else **dev_local** to fetch details from ```aws ssm``` connecting from local
  Use **any env name** to fetch details from aws ssm if the workload is going to run inside aws
- Use ```serverless local``` to run the code in local
- To run in a lambda environment make sure to deploy the application to you aws using ```serverless deploy```1


**Points to Note**

- Dynamodb data modelling is different. This is just a sample usuage of the CRUD operations
- Access patterns should be determined prior to working on this with a partionkey chosen with hiher cardinality
- To avoid ```hotspotting``` in dynamodb, it is always better to append hashkeys with unique numbers and create the data so the data is distributed across partitions
- For now Provisioned capacity is used considering the fact that using on-demand mode being costly (use it only during cases of unpredictable workloads)



**TODO's**

- Make a dockerized version
- Dapr needs to be explored
- Kubernets deployment - to be implemented
- ```serverless``` - to be implemented for the dockerized container


**Access Patterns addressed**

 - This is a location table created for a bank , combination of locId&country serve as Primary Key. Basically we are not worried about the storage space , so we can store as many replica of table. Primary function is to make the retrieval faster
 - Fetch all locations for a country
 - Fetch address of the location given locationId
 - Fetch all locations based on a county - PK - combination of locId&county can be used to store the address again
 - Fetch all locations given a state - PK - combination of locId&state can be used to store addresses again
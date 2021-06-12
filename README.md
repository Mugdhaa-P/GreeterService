# GreeterService
GreeterService using Twirp and DynamoDB

Functions:  
`SetGreetingForUser(name)->None`:   Stores a greeting message in Table `Users` by randomly selecting greeting phrases  
`GetGreetingForUser(name)->message`:   Fetches the message from the database by 'name'

To execute file:
- Add DynamoDB credentials   
  `cd GreeterService/internal/greeterserver/server.go`    
  Add credentials to ConnectDynamo() function:   
  `credentials.NewStaticCredentials("AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "TOKEN")`
  
- Run the greeterserver on a terminal:  
  `cd GreeterService/cmd/greeterserver/main.go`      
  `go run main.go  `
  
- Run the greeterclient on another terminal:  
  `cd GreeterService/cmd/greeterclient/main.go`    
  `go run main.go `
  

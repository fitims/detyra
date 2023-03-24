## Decisions made ##
            
The Wallet App will hold its users on a map data structure. The key for the map will be user's email address. 

In real-life scenarios a user would have more than one wallet, but to simplify the example I choose to have only one wallet per user.

A very good data store for the users and wallets would be a key/value data store (like Redis, BadgerDB, etc.) or even a graph database (Dgraph, Neo, etc.) 
but to simplify the example I decided to mock the key/value data store (**in_memory** package that is inside **repo** package).

To keep the app performant, I would use pipeline concurrency pattern, so for every transaction I would publish a transaction in a channel, and would have 
another goroutine that listens to the transaction channel and save the transaction in database concurrently.


## Setting up and executing the code ##

#### Running the app ####

To run a test code, execute the following command:
```
go run main.go
```

#### Running unit tests ####

To run unit tests for the wallet, go to the wallet folder by using the command:
```
cd ./wallet
```
thne run unit test by using the command:
```
go test -v
```

To run unit tests for the store, go to the store folder by using the command:
```
cd ./store
```
thne run unit test by using the command:
```
go test -v
```          

## Areas to be improved ##
                              
* More unit tests and better test coverage
* Make the application concurrent
* If there is a large number of users, split the users into shards and run a separate instance for each shard
* Use a distributed key/value store to store users and wallets
* Use a blob storage to store transactions 
* Setup a monitoring and alerting 


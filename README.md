# AroundHome | Customer<>Provider Match Service

# Project Layout 

## Project Milestones (metric: x=hour)
- Understanding the needs and design architecture as a template (1x)
- Development of the architecture and apply the edge cases (2x)
- Development of services (2x)
- Testing and Documentation (1x)
- Deployment and get ready for sharing (1x)
## Project ToDo's
- [ ] Understanding the needs and requirements.
- [ ] Create a basic project layout and git repo.
- [ ] System design and skeleton development.
- [ ] Docker integration database connection
- [ ] Database and schema design
- [ ] Development of internal services
- [ ] Development of endpoints
- [ ] Development of unit tests & relavent tests
- [ ] Add godoc for autogenerated documents
- [ ] API coding documentation 
- [ ] Deployment 
- [ ] Sharing with AroundHome team.

## Real Life Architecture Desing



## Task Architecture Design (Due to time constraints)




You can access drawings also in here on [DrawIO](https://app.diagrams.net/#G1ku_myFdK3BX2mKC8As7pY5j2icQLDEi-#%7B%22pageId%22%3A%22tLzeWMEuZmfhtKyOjQjI%22%7D)

## Decisions 



## Missing parts due to time constraint



# How to use ?

## Accessing the project 

You can access the project via [github](https://github.com/AkyurekDogan/around-home-task). 

## How to run ?

### Step-1

You can run the code in different ways however, in every case you need the following setup 

- `install docker` if you do not have

### Step-2

You can choose one of the following options in `Step-2`

#### Running the code (Recommended)

- `install the golang version ~1.23.2`
- You can use any IDE to see the code however, I used `VSCode` so if you are `VSCode` user, you can use the `.vscode/lunch` configuration to run the code in `debug mode`
- Run the following make file command (you can see the makefile at the root directory of the code directory) `make run-local-docker-db`. This command will initiate local docker `postgres` instance with some initial data which is defined in the `init.sql` file in the `/scripts/database` folder.
- The database connection is ready at `localhost:5432` with `username:postgress` and `password:mypassword123!`. You can also see these details in `./config.yml` file
- Both you can run the `make run` command to run the service or use the IDE debugger functionality.
- The service will be running on `localhost:1989`

#### Using as a docker compose

- You can see the `docker-compose.yml` file in the root directory of the repository. 
- I recommend to change the database connection in the `./config.yml` file `database.host: go-around-home-task-postgress` so the service can access in the compose base port and host.
- run the `make run-docker-compose` command 
- The service will be running on the `localhost:3000` (ports are used different to understand the traffic root better)

### Step-3

Depending on the previous steps the service must be running on either `localhost:1989` or `localhost:3000` so you can use the following commands to use the service 

- `localhost:1989 GET`
- `localhost:1989 POST`

### Step-4

You can see the definitions of the services as follows;

#### HTTP:GET /cost 

Returns the cost calculation regarding to price and usage data.



#### HTTP:POST /meter 

Gets the metering data and saves to database.


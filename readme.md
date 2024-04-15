
# This is the best ever golang starter 🚀

The goal of the repository is to condense my experience as a Golang developer over the past 8 years into an easily usable, understandable package that incorporates best practices. I hope this contributes to serving as support for users of the go language. My aim is to assist the Golang community in an area where I feel guidance is needed: code structure.

This is not your typical starter template but rather a coding structure guide that provides assistance from folder naming conventions to the integration of various technologies like mongodb, postgres, dataloader, pub-sub, etc.

This kit can be helpful to you if you're looking for the following in your project:

- ✍🏻 Interface-first design
- #️⃣ Hexagonal architectural design
- 🧱 Domain-driven module system
- 👉🏻 Simplicity with minimal abstractions on top of the language-layer
- ♻️ Reusable modules
- 💉 Dependency injection (ohh yes the fancy word of arguments...)
- 🎭 Mock-first testing approach
- 🧐 Monorepository pattern


### Embracing the Diverse Approaches in Go Development

I've worked in many teams as a Go engineer, and I've seen as many approaches to the same problems as there are teams or even more. In my experience, there's no such thing as a standardized Go project. On one hand, this adds a certain romance to the language because many people enjoy this "freedom." However, this freedom also greatly complicates the use and onboarding of the language.

Competing languages on the market, such as JavaScript or Java, have been offering well-structured frameworks and solutions for many years. Think of React, for example, which has set a new direction for frontend development. So, it's a valid question: why isn't there something like this in Go?

A more relevant question would be whether it's good that there isn't such a thing in Go. My answer to this question is not black and white either. I see the language as a kind of superset language (like the Down band), so well put together that it doesn't require robust frameworks like PHP's Laravel or Ruby on Rails. However, it does require good maintainable structure, naming conventions, and adherence to best practices.

I think this has been interpreted in the developer community as not needing frameworks. Since there's no prominent player providing guidance in the absence of frameworks, the question of structure has been left to individual development teams, thus leaving the field more chaotic.

The starter kit aims to take a stance in favor of structuring, in a more constrained, conservative form, without the goal of creating a framework.

# About the project

In the example project, we're building a lightweight blog. The blog will be accessible via a GraphQL gateway in the form of a web server. The blog has two business entities: categories and posts. Each post can be assigned to one or more categories. I've implemented CRUD operations for both domains as well as a data loader to assist with data population on the GraphQL side.

## Run the application

Step 1 - open terminal from project folder and run `docker compose up`
Step 2 - `cp .env.example .env`
Step 3 - run `make run-simple-blog-graphql` OR use Visual Studio Code, Run and Debug Section to run simple-blog-graphql (Recommended way)
Step 4. open `http://localhost:9099`


# About the structure

The starter kit follows a lightweight, hexagonal architecture-based concept. This means that there are domain modules that are in some kind of relationship with various actors, either on the driven or driver side. Actors can be driven, such as a database or cache layer, but they can also be driver-side, such as a web server or a CLI server. The separation of actors and modules related to the domain enables wide reuse of domain modules, for example, within a project, a domain module can be accessible as part of a binary-based CLI tool or as a running web server endpoint.

The project follows monorepository pattern, which allows for not just one entry point for the application, but any number of them. It's important to note that this demo project doesn't adhere to the Go workspace-based multi-service structure. The reason for this is that I didn't want to overcomplicate things from the start, thus aiding in understanding the basics.


## Base Concepts

### Actors

Driving (or primary) actors are the ones that initiate the interaction. For example, a driving adapter could be a controller which is the one that takes the (user) input and passes it to the Application via a Port.

Driven (or secondary) actors are the ones that are “kicked into behavior” by the Application. For example, a database Adapter is called by the Application so that it fetches a certain data set from persistence.

The actors are placed in either the `/internal/actor` or `/pkg/actor` directory. Personally, I don't like to distinguish between driver-side and driven-side actors at the folder level, but if it helps in a particular project, actors can be marked using subdirectories (`/internal/actor/driver`, `/internal/actor/driven`). In my example project, I've structured it as follows:

```
internal/
  actor/
    db/
      ...
      mongodb/
        mongodb.go
        mongodb.mongodriver.go
      ...
```

In this structure, the actors are divided based on technologies. If we focus solely on the database actors, we can see that each database technology has its own folder. Within a given technology, we distinguish between definition files and implementation files. The definition files always bear the domain name. In my example, this is `mongodb.go`. These definition files contain a blueprint that we define for ourselves, specifying what we want to do with a given technology in our project. This definition may provide less functionality than the package with which we implement it, but that's perfectly fine. We can implement the definition files with anything, as long as we indicate the type of implementation somehow in the file and its name. In my example, `mongodb.mongodriver.go` is the implementation file because I implemented my definitions with the mongodriver package. If I were to use another package, the file name would be different, for example `mongodb.mgo.go` in the case of using the mgo package. It's also possible and even recommended to have multiple implementation files simultaneously, especially if we want a smooth transition from an old, outdated, or deprecated package to a new one.

It's not worth creating a generic interface for a specific actor because we might lose the strength of the particular tool if we try to match it generically. For example, in the case of databases, we could create a generic interface with methods like GetOne, GetMany, etc., but this would diminish the power of certain databases.

### Module / Domain Module:

- A module, or domain module, is essentially a logically separable unit that closely covers a business or domain-specific requirement.
- A module can depend on other domain modules and can be a dependency of other modules.
- Each module is always placed in its own directory under the domain name.
- The modules are housed in a collector `/module` folder, although the naming can vary, for example, `/domain`. While it's not mandatory to organize domain modules into separate folders, I've found it highly beneficial in my projects to maintain a separate folder for readability, searchability, and to accommodate the increasing complexity over time.
- Dependency injection - A module's dependency is always an interface definition; otherwise, we get a hard dependency, which among other things, makes unit testing impossible.


#### What does a module consist of?

1. Definition file

   This file always carries the domain name. It contains important definitions, the building blocks of the given domain:
    - Provider Definitions
    - Entities
    - Enums
    - DTOs (Data Transfer Objects)

2. module.go file
   
   This file is responsible for collecting dependencies of individual providers, initializing the providers, exporting the providers, and mapping to external interfaces.

3. Providers and their corresponding implementation files.
  
4. Other files or directories
   
   Closely related to the domain but not fitting into the above categories, such as GraphQL schema files, sqlc querier files, etc.

5. Unit Test files

    Unit test files for the implementation files. It's important that only the unit tests of that particular module can be placed in the domain module. Integration, system, or end-to-end tests should not be placed in the domain module. I recommend taking the definition of unit tests seriously and truly writing tests on a unit basis. This means that a test file for a repository implementation should not test the service or other implementations. Focus solely on testing the specific unit.

```
internal/
  module/
    category/
      category.go
      dataloader.go
      module.go
      repository.mongodb.go
      schema.graphql
      service.go
    post/
      post-querier/
        ...sqlc generated files
      post.go
      dataloader.go
      module.go
      repository.postgres.go
      schema.graphql
      service.go
```


In the example project, we are creating a simple blog. This blog has two domain modules, `category` and `post`. The structure shows that each domain has its own folder, and the various files under them are part of the same package by default. This kind of separation allows for complete isolation of the individual domains, so there is no `service` package containing category and post services, but rather there are post and category packages containing service implementations and definitions.

In the sample code, it can be seen that different technologies were used for implementing the repository definitions in the modules. The `category` repository provider implementation was done with MongoDB, while the `post` repository provider implementation was done with PostgreSQL (using sqlc). The naming conventions for the repositories follow the same pattern as with the actors, where the definition/implementation scheme is followed. The only difference is that if it is known that there will be no other implementation of a particular provider definition in the foreseeable future, the indication of the implementation technology can be omitted from the name. An example of this is the `service.go` file in my case.


### Providers (Provider Definition, Provider Implementation)?

- A provider covers a specific part of the tasks required for executing the business logic.
- There are providers that are intended to introduce a specific technology (Actor) into the domain module, but a provider can also be a logical unit performing complex computational tasks.
- Providers can depend on other providers within or outside the domain module.

It's important to note that this boilerplate is not a framework, so there isn't a predefined set of providers that must be used. In my example code, you'll encounter provider definitions and their implementations of types such as service, repository, and dataloader, but you're free to create your own definitions. Personally, I always ensure that a particular provider implementation focuses only on a specific technology or Actor.

Useful providers could include:

- Service
  
  A provider responsible for serving the defined business needs in a technology-agnostic manner.

- Repository
  
  A provider responsible for stateful operations. A repository implementation could be database-based or even file-based.

- Pub/Sub
  
  Used in publish-subscribe or event-driven systems to handle subscriptions and operations.

- Dataloader
  
  A specialized provider used to address the n+1 problem commonly encountered in GraphQL.

# Folder Structure

```
cmd/
  gateway1/
    ...
    main.go
  gateway2/
    ...
    main.go
db/
  migrations/
    ...
internal/ or pkg/
  actor/
    db/
      actor1/
        actor1.go
        actor1.technology1.go
        actor1.technology2.go
  module/
    domain1/
      domain1.go
      module.go
      provider-implementation1.go
      provider-implementation2.go
      provider-implementation3.technology1.go
      ...
    domain1/
      domain1.go
      module.go
      provider-implementation1.technology1.go
      provider-implementation2.go
      provider-implementation3.technology1.go
      ...
  util/
    config/
      config.go
      config.technology1.go
    ...
mocks/
  GeneratedMock1.go
  GeneratedMock2.go
docker-compose.yml
Makefile
tools.go
...
```

Let's break these down and examine each part individually.
Note: We won't repeat the discussion about the actor and module folders as we've already covered them under the base concepts.

`cmd/`
   
  The cmd folder contains the entry point(s) of the project. This could be a single gateway or even multiple ones. In the case of a more complex application, we might need a worker and a public gateway, but for a simpler application, a single entry point might be sufficient. Think of these entry points as a playground for children. It's a space where we can assemble our own world using many building blocks. So, within an entry point, we can compose applications using the provided domain and actor sets. The choice of technology for a given application always depends on the available actor set. An application could be based on technologies like CLI, gRPC, web, or others.

   - If the project has a single entry point, subfoldering can be omitted.
   - A gateway should contain only files related to the gateway itself. These files are always located within the gateway directory. In the example code, I used a subfolder named `/simple-blog-graphql` and within that, I separated a folder recommended by the gqgen package `/graph` in this folder we can find gqlgen specific files. Of course, if necessary, gateways/entrypoints can be viewed as individual applications within the project, so the naming convention can be maintained here as well in terms of folders (/internal, /pkg, etc...). However adding too much complexity in gateways is not recommended.

`/db`

  The `/db` folder should not be confused with the database folders of the actor set. This folder contains database migration files and other files related to the global database configuration. In my example code, the `/db/migration/` subfolder contains the schema for the PostgreSQL database. This is necessary because unlike NoSQL databases, PostgreSQL and other SQL-based databases need to be migrated before use. This is an operation that is not strictly part of the code but is essential for it to function. In my example code, the database migration is handled by a third-party tool, `go migrate`. This package allows us to maintain a version-controlled, historical migration flow that supports jumping between different database schema versions. If the project does not require the `/db` folder (for example, if it only uses MongoDB), then this folder can be omitted.

`internal/` or `pkg/`

  The `/internal` or `/pkg` folders are used to house domain modules, actors, utilities, and other packages that are integral parts of the codebase. Use the `/internal` if you don't want the packages and modules in it to be accessible to other projects / packages, and use the `/pkg` folder if you do. In general, everything should be placed in the `/internal` folder until there is a need for the opposite.

`/mocks`

  The mocks folder hosts the mock files of providers and other interfaces. Dependency injection and the interface-first approach allow us to mock dependencies, making this folder useful for unit testing. In the example project, we use `github.com/vektra/mockery` for generating mocks, but of course, you can use other generators or write your own mock functions in the `/mocks` folder as well.

`docker-compose.yml` and `.env.example`/`.env`

  These files facilitate local development. The docker-compose file sets up an environment with the dependencies necessary for initializing the actors, while the .env file contains the required environment variables. 

`Makefile`

  The Makefile is a file containing various scripts that help us access/execute tasks related to the project, providing us with an easy-to-use and comprehensible tool. Its usage is simple and recommended. In the example project, the Makefile includes commands responsible for:

  - handling migrations,
  - generating mocks,
  - running gateways,
  - generating gqlgen and sqlc files, and
  - running unit tests.
  
  Of course, it's not necessary to use a Makefile, and other tools can be used instead. However, in my experience, the Makefile is the fastest and simplest script tool, widely known and used. There are other tools, such as Docker container-based ones, but they offer solutions more slowly.


### Using Environment Variables
  It's important not to use environment variables directly in domain modules. Strive to ensure that all dependencies of the domain module can be injected. Therefore, loading environment variables should always be the responsibility of the gateways located in the `/cmd` folder.

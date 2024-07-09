# Graphql 
# GraphQL Jokes API

This project is a GraphQL-based API for storing and retrieving jokes. It demonstrates the use of GraphQL for creating, reading, and managing jokes with a PostgreSQL database as the backend. The API supports basic CRUD operations through GraphQL queries and mutations.

## Prerequisites

- Go 1.22.4 - You should install Go on your system. My version of Go is 1.22.4.
- PostgreSQL
- gqlgen (GraphQL server library for Go)

To get the basic know how of what is Graphql and why we use it follow any latest medium article.
In Graphql POST, PUT, DELETE, GET and PATCH request replaced by QUERIES and MUTATIONS.
I use postgres database for storing the data.
Firstly create a table in some exsisting database with the Migration mutation.
Secondly user enter some random jokes enters in database using CreateJoke mutation.
With the help of queries like joke and jokes user can retrieve a single joke or get all jokes stored in the database with the choice in hand of retrieving id with created joke/jokes or not.
For my own here are some examples of queries and mutation that Graphql made:

mutation Migration {
    migration {
        message
    }
}

mutation CreateJoke {
    createJoke(input: {text:"Any value in text"}) {
        id
        text
    }
}

query Joke {
    joke(id: "1") {
        id
        text
    }
}

query Jokes {
    jokes {
        id
        text
    }
}

NOTE: In queries we can get only id or text or both deponds the type of data which we want to retrieve or not.

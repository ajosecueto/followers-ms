# Go + GqlGen + Neo4j + Kafka

Simple MS using Graphql + Kafka

## Foldering

- Events:
    - Aggregation: it contains all producers events and orchestration
    - Reducer: it contains an event consumer separated by topics and business logic
- Graph
    - Domain, contain all the business logic and api.
    - Include Mutation as commands and Queries
- Infrastructure
    - kafka: Producer and Consumer initialization
    - neo4j: Driver and session management
- Persistence
    - All related to data source  

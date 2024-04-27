# Money Manager

## Tools
- Go Migrate
- Justfile
- Watchexec
- Curl
- Docker + Docker Compose
- PostgreSQL

## Dependencies
- Gin Gonic
- Go sqlx + goqu
- gotestdox
- Testify
- Test Container

## RestAPI Endpoints
- [ ] Create Transaction [POST]
- [ ] Get List Transaction [GET]
- [ ] Get Detail Transaction [GET]
- [ ] Update Transaction [PUT]
- [ ] Delete Transaction [DELETE]

```ts
type Transaction = {
    id: uuid
    title: string
    note: *string
    amount: uint64
    type: 'in' | 'out'
    created_at: timestamp
    updated_at: timestamp
    deleted_at: timestamp
}
```
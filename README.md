# go-project-structure

TODO sum up once blog post is written

Insert tasks:
```bash
# insert expired task
curl -X POST localhost:8080/task -H 'Content-Type: application/json' -d '{"value":"some value...","date":"1000-01-01T00:00:00Z"}'

# insert future task
curl -X POST localhost:8080/task -H 'Content-Type: application/json' -d '{"value":"some other value...","date":"3000-01-01T00:00:00Z"}'
```

Get tasks:
```bash
curl localhost:8080/future_tasks
```
# Mono Repo Cacthus Challenge

## Structure

```
/
├── Backend
│   .├── Golang // Golang files
│   .├── Phoenix // Elixir files
├── Frontend // react files
├── docker-compose.yml

```

### How to run

In a terminal run below command

```
docker-compose up -d
```

the UI will be displayed in [http://localhost:3000](http://localhost:3000)

Info

- The golang service insert fake data every 10 seconds
- Phoenix service send data through broadcast channel every 5 seconds
- Fake data its random.
- In UI show the last 50 records

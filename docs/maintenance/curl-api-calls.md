## How to test API endpoints with curl

During development, it's important to test that all the API calls work properly before deploying. Right now, there are no tests implemented (this will be added in the future). Below is a list of all the available endpoints, and a curl command that calls each of them.

_Note, some endpoints won't work on their own. Most of them require a user to be created in the database. This can be done by calling any of the subsequent `/initialize-game` endpoints. Error messages are provided in the endpoint responses that require a user to be created.`_

### Common Handlers

---

Health handler:

```
curl -i -X GET http://localhost:8080/api/health
```

Remaining time handler:

```
curl -i -X GET http://localhost:8080/api/remaining-time
```

User game result handler:

```
curl -i -X GET "http://localhost:8080/api/finished-games?user_id=123"
```

## Daily Slash Tests

Daily-slash initialize game handler:

```
curl -i -X GET "http://localhost:8080/api/daily-slash/initialize-game?user_id=123"
```

Daily-slash get search weapons:

```
curl -i -X GET "http://localhost:8080/api/daily-slash/search-items"
```

Daily-slash get hint:

```
curl -i -X GET "http://localhost:8080/api/daily-slash/hint?hint=1"
```

Daily-slash check guess:

```
curl -X POST http://localhost:8080/api/daily-slash/check-guess \
 -H "Content-Type: application/json" \
 -d '{"user_id": "123", "guess": 10}'
```

Daily-slash winning data:

```
curl -i -X GET "http://localhost:8080/api/daily-slash/winning-data?user_id=123"
```

Daily-slash player guess count endpoint:

```
curl -i -X GET "http://localhost:8080/api/guess-count?mode=daily-slash"
```

## Connections Tests

Connections initialize game handler:

```
curl -i -X GET "http://localhost:8080/api/connections/initialize-game?user_id=123"
```

Connections check guess:

```
curl -X POST http://localhost:8080/api/connections/check-guess \
 -H "Content-Type: application/json" \
 -d '{"user_id": "123", "guess": ["Dunerider Boots","Shark","Bast Statue","Webbed"]}'
```

Connections reveal answers:

```
curl -X POST http://localhost:8080/api/connections/reveal-answers \
 -H "Content-Type: application/json" \
 -d '{ "user_id": "123" }'
```

Connections winning data:

```
curl -i -X GET "http://localhost:8080/api/connections/winning-data?user_id=123"
```

## Guess the NPC Tests

Guess the NPC initialize game handler:

```
curl -i -X GET "http://localhost:8080/api/guess-the-npc/initialize-game?user_id=123"
```

Guess the NPC get search npcs:

```
curl -i -X GET "http://localhost:8080/api/guess-the-npc/search-items"
```

Guess the NPC guess check:

```
curl -X POST http://localhost:8080/api/guess-the-npc/check-guess \
 -H "Content-Type: application/json" \
 -d '{"user_id": "123", "guess": 10}'
```

Guess the NPC winning data:

```
curl -i -X GET "http://localhost:8080/api/guess-the-npc/winning-data?user_id=123"
```

Guess the NPC name guess check:

```
curl -X POST http://localhost:8080/api/guess-the-npc/check-name-guess \
 -H "Content-Type: application/json" \
 -d '{"user_id": "123", "guess": "Agnew"}'
```

## Hangman Tests

Hangman initialize game handler:

```
curl -i -X GET "http://localhost:8080/api/hangman/initialize-game?user_id=123"
```

Hangman check guess:

```
curl -X POST http://localhost:8080/api/hangman/check-guess \
 -H "Content-Type: application/json" \
 -d '{"user_id": "123", "guess": "A"}'
```

Hangman winning data:

```
curl -i -X GET "http://localhost:8080/api/hangman/winning-data?user_id=123"
```

## TerraTrivia Tests

TerraTrivia initialize game handler:

```
curl -i -X GET "http://localhost:8080/api/terratrivia/initialize-game?user_id=123"
```

TerraTrivia check guess:

```
curl -X POST http://localhost:8080/api/terratrivia/check-guess \
 -H "Content-Type: application/json" \
 -d '{"user_id": "123", "guess": "SHRUBSTAR"}'
```

TerraTrivia winning data:

```
curl -i -X GET "http://localhost:8080/api/terratrivia/winning-data?user_id=123"
```

Here are all the API test calls to make:

## Common Handlers

Health handler:
curl -i -X GET http://localhost:8080/api/health

Remaining time handler:
curl -i -X GET http://localhost:8080/api/remaining-time

## Daily Slash Tests

Daily-slash init game handler:
curl -i -X GET "http://localhost:8080/api/daily-slash/initialize-game?user_id=123"

Daily-slash get search weapons:
curl -i -X GET "http://localhost:8080/api/daily-slash/search-items"

Daily-slash get hint:
curl -i -X GET "http://localhost:8080/api/daily-slash/hint?hint=1"

Daily-shash check guess:
curl -X POST http://localhost:8080/api/daily-slash/check-guess \
  -H "Content-Type: application/json" \
  -d '{"user_id": "123", "guess": 10}'

Daily-slash winning data:
curl -i -X GET "http://localhost:8080/api/daily-slash/winning-data?user_id=123"

## Connections Tests

Connections init game handler:
curl -i -X GET "http://localhost:8080/api/connections/initialize-game?user_id=123"

Connections check guess:
curl -X POST http://localhost:8080/api/connections/check-guess \
  -H "Content-Type: application/json" \
  -d '{"user_id": "123", "guess": ["Dunerider Boots","Shark","Bast Statue","Webbed"]}'

Connections reveal answers:
curl -X POST http://localhost:8080/api/connections/reveal-answers \
  -H "Content-Type: application/json" \
  -d '{ "user_id": "123" }'

Connections winning data:
curl -i -X GET "http://localhost:8080/api/connections/winning-data?user_id=123"

## Guess the NPC Tests

Guess the NPC init game handler:
curl -i -X GET "http://localhost:8080/api/guess-the-npc/initialize-game?user_id=123"

Guess the NPC get search npcs:
curl -i -X GET "http://localhost:8080/api/guess-the-npc/search-items"

Guess the NPC guess check:
curl -X POST http://localhost:8080/api/guess-the-npc/check-guess \
  -H "Content-Type: application/json" \
  -d '{"user_id": "123", "guess": 10}'

Guess the NPC winning data:
curl -i -X GET "http://localhost:8080/api/guess-the-npc/winning-data?user_id=123"

Guess the NPC name guess check:
curl -X POST http://localhost:8080/api/guess-the-npc/check-name-guess \
  -H "Content-Type: application/json" \
  -d '{"user_id": "123", "guess": "Agnew"}'

## Hangman Tests

Hangman init game handler:
curl -i -X GET "http://localhost:8080/api/hangman/initialize-game?user_id=123"

Hangman check guess:
curl -X POST http://localhost:8080/api/hangman/check-guess \
  -H "Content-Type: application/json" \
  -d '{"user_id": "123", "guess": "A"}'

Hangman winning data:
curl -i -X GET "http://localhost:8080/api/hangman/winning-data?user_id=123"
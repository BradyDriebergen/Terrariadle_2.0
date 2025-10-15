**All of these values are within /api/{mode}**

GET /puzzle-data
- This route gets the puzzle data for each of the games. This data will
include all the information necessary for users to start guessing. Data
can include quotes, categories, hints, etc.

GET /user-guesses/
- This route gets all the user guesses for a specified mode. Guesses can
range from a list of indexes to guessed categories.

GET /search
- This is a search function that returns a list of objects when a user
searches. Values change depending on the mode, but this will allow users
to search without having to store all the options on the front end.
Ensure use of a debouncer when using this endpoint to avoid users 
over-guessing

POST /check-guess
- This route checks the guess that a user makes. Make sure to include the
user's ID in the body of the request. The user's guess will be checked, and
It will return the object they guessed alongside a boolean on whether they won or not.

GET /get-position/{user-id}
- This route responds with the position of when the user successfully
guessed compared to all the other users. 

GET /players-guessed
- This route returns the total number of players guessed in a given day

GET /remaining-time
- This route returns the remaining time in the day for users to guess
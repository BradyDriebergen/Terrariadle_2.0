## Running Terrariadle v1.0

Here is how I ran the frontend (was Node.js, not angular):

```bash
#!/usr/bin/env bash
set -euo pipefail

cd ~/Terrariadle/client

while true; do
  echo "[$(date)] starting Angular frontend..."
  serve -s build > /dev/null
  code=$?
  echo "[$(date)] ng serve exited with code $code; restarting in 2s..."
  sleep 2
done
```

And here is how I ran the backend:

```bash
node ./src/DailyGuessAPI.js
```

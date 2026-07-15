## How to update dependencies:

With vulnerabilities being found every day, it's important to keep dependencies up to date. If there are ever vulnerabilities reported by NPM or GitHub's Dependabot, here is how you do it from the command line:

**Go dependencies:**

```
# See what's outdated
go list -u -m all

# Update everything to latest minor/patch versions
go get -u ./...
go mod tidy
```

**NPM dependencies:**

```
# See what's outdated
npm outdated

# Update npm using audit
npm audit fix
```

This won't fix every issue, but it will fix a majority of them. If there is ever a dependency that won't update after these commands, manual updating is required. I don't have an example now, but it usually requires going into `frontend/package-lock.json` and updating the version manually. After doing this, you can run the following commands:

```
# Install the new dependency
npm install

# Check to see if the issue is still there
npm audit
```

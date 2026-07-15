## If versioning ever breaks:

Versioning is a bit delicate, but it shouldn't break because of the branch protection in place. However, if it ever does, you can reset the version by running the following commands in the `main` branch in git:

```
# deletes it from remote
# a version number looks something like v2.0.0
git tag -d <version number>
git push origin :refs/tags/<version number>

git tag <version number>
git push origin <version number>
```

This deletes the old version number, which is stored in a git tag, and replaces it on the current commit. This resets the current HEAD of main version tag.

Also, if it ever mentions `-dirty` on the version, it means that there are uncommitted changes in the default branch. Make sure to switch branches and commit the changes to get rid of this postfix.

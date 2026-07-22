`diff files-before.txt files-after.txt`

```
7,13d6
< backend/data/catagories.json
< backend/data/categories.json
< backend/data/enemies.json
< backend/data/npcs.json
< backend/data/terms.json
< backend/data/terraria_terms_2000.json
< backend/data/weapons.json
```

```
❯ git rev-list --all --count
285
```

```
❯ git rev-list --all --count
280
```

```
❯ git push --force --mirror origin
Enumerating objects: 3766, done.
Writing objects: 100% (3766/3766), 12.17 MiB | 9.79 MiB/s, done.
Total 3766 (delta 0), reused 0 (delta 0), pack-reused 3766 (from 1)
remote: Resolving deltas: 100% (1712/1712), done.
To github.com:BradyDriebergen/Terrariadle_2.0.git
 + d72b51a...542d829 Documentation -> Documentation (forced update)
 + 56ca2af...9de81a8 main -> main (forced update)
 + 56ca2af...9de81a8 v2.0.0 -> v2.0.0 (forced update)
 ! [remote rejected] refs/pull/1/head -> refs/pull/1/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/10/head -> refs/pull/10/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/11/head -> refs/pull/11/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/12/head -> refs/pull/12/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/13/head -> refs/pull/13/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/14/head -> refs/pull/14/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/15/head -> refs/pull/15/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/16/head -> refs/pull/16/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/17/head -> refs/pull/17/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/18/head -> refs/pull/18/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/19/head -> refs/pull/19/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/2/head -> refs/pull/2/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/20/head -> refs/pull/20/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/21/head -> refs/pull/21/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/22/head -> refs/pull/22/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/3/head -> refs/pull/3/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/4/head -> refs/pull/4/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/5/head -> refs/pull/5/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/6/head -> refs/pull/6/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/7/head -> refs/pull/7/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/8/head -> refs/pull/8/head (deny updating a hidden ref)
 ! [remote rejected] refs/pull/9/head -> refs/pull/9/head (deny updating a hidden ref)
error: failed to push some refs to 'github.com:BradyDriebergen/Terrariadle_2.0.git'
```

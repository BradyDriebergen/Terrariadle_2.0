# Removing Data Files from Git History

I've been developing Terrariadle 2.0 for close to a year now. During that time, I've gone through a lot of different design decisions. One of the initial design choices I made was storing my game data in JSON files rather than my database. I stored data like this in my previous iteration, and it was the easiest way of getting my initial POC for this version working. I ended up taking these data files out later in development, but quickly learned that they remained in GitHub. Not wanting to reveal my game data, I started researching a way to remove these specific files from my Git history.

One of the main things I wanted to preserve is the full commit history. It would be super easy to copy the files to a new GitHub repo and pushed them up, but I wanted to show how long I've been working on this project. This introduced a problem where I wanted these data files gone from every commit without losing any of the commit history that shows how long I've actually been building this.

In my research, I quickly found a tool that can strip specific files out of the entire git history while preserving every commit, its author, and its original date. It looked scary, risking my history with no easy way of backing it up, but it was something I had to do.

## How I Did It

### Initial check

I used `git filter-repo`, which is the tool GitHub recommends for rewriting history (it replaced the older, slower `git filter-branch`). Unlike simply deleting a file and committing that change, `filter-repo` goes back through every commit in the repo's history and removes the target file from each one, so it's as if the file was never committed in the first place.

Before changing anything, I first ran commands to confirm which files existed in history:

```bash
git log --all --pretty=format: --name-only --diff-filter=A | sort -u
git log --all --pretty=format: --name-only | sort -u > files-before.txt
```

The files I wanted to remove were the following:

- `backend/data/catagories.json`
- `backend/data/categories.json`
- `backend/data/enemies.json`
- `backend/data/npcs.json`
- `backend/data/terms.json`
- `backend/data/terraria_terms_2000.json`
- `backend/data/weapons.json`

I also ran the following to check how many files were tracked:

```bash
$ git rev-list --all --count
285
```

I also ran `git filter-repo --analyze` on a test clone, which produces a report of every path in history and its size. This told the same thing as the commands above, but provided me with a bit more confidence.

### Making backups before touching anything!

I never ran `filter-repo` on my real working copy. I didn't want to risk messing anything up on the main repo. Instead, I created a full mirror backup clone as an untouched restore point, and a separate mirror clone to actually perform the rewrite on:

```bash
git clone --mirror https://github.com/BradyDriebergen/Terrariadle_2.0.git terrariadle-backup.git
git clone --mirror https://github.com/BradyDriebergen/Terrariadle_2.0.git terrariadle-rewrite.git
```

This meant that no matter what happened during the rewrite, I had an exact copy of the original repo I could restore from.

### 4. Running the filter

Inside the rewrite mirror, I ran `filter-repo` with `--invert-paths` to remove the target files while keeping everything else:

```bash
git filter-repo --path-glob 'backend/data/*' --invert-paths
```

`--path-glob` allowed me to remove everything within the `backend/data/` directory without me having to specify every individual file. If I wanted to specify specific files, I would've used `--path`.

### 5. Verifying the rewrite worked correctly

Before pushing anything to GitHub, I checked that the target files no longer appeared anywhere in history and no other files had been accidentally removed. I checked the ladder by running the following:

```bash
git log --all --pretty=format: --name-only | sort -u > files-after.txt
diff files-before.txt files-after.txt
```

The `diff` command printed the following:

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

I also checked the tracked file count:

```bash
$ git rev-list --all --count
280
```

This verified that my `filer-repo` worked correctly.

### 6. Re-added the remote and pushed

`filter-repo` automatically strips the `origin` remote, so I had to re-add it before pushing:

```bash
git remote add origin https://github.com/BradyDriebergen/Terrariadle_2.0.git
git push --force --mirror origin
```

Running this returned:

```
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

The push force-updated all my real branches and tags with the rewritten history. GitHub also rejected pushes to `refs/pull/*/head`, the hidden refs it auto-generates for pull requests, which wasn't an issue since those aren't real branches, just made for PRs.

### 7. Replaced my local working copy

Since every commit SHA changed during the rewrite, my old local working folder didn't share any history with the updated remote. Rather than trying to reconcile it, I renamed the old folder and cloned a fresh copy from GitHub to continue working from.

## Takeaway

This process was pretty scary dealing with. I've had plenty of issues with GitHub in the past with merging and rebasing. I didn't want to permanently screw up my Git history which I was trying so hard to preserve. In the end however, I successfully removed these files from my Git History and keep the months of history behind this rewrite.

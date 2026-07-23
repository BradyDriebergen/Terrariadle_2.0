# How to update:

This document should be a step by step list on how to update the site with new changes.


## Starting out

Make sure that you close out both tmux instances. Not sure if this helps with restart time
but it doesn't hurt to do it.

Also, before you start this process, make sure all of the files are pushed to github.

After that, restart the instance. This use to take like 5 minutes but recently its taken close
to 20 minutes. It's annoying but be patient.


## Front-end Production Build Steps

Rather than uploading all the code and running it through a development build, use the built-in 
production build command. The first step is to be in the react project directory on your 
development machine:

```bash
$ cd /Terrariadle/client
```

After this, run the following command to make the production build:

```bash
$ npm run build
```

After that, you should have a new/updated build file in your react project directory.


## Uploading Updated Files

Now that you have your files ready to transfer, start by moving to the project directory: 

```bash
cd ~/Projects/Terrariadle
```

After that, run the following command to move all back-end files:

```bash
rsync -avz --exclude-from=./api/.rsyncignore -e "ssh -i ~/Projects/ssh-key-2025-02-17.key" ./api/ ubuntu@146.235.238.252:/home/ubuntu/Terrariadle/api
```

The last command moves the build folder to the instance.

```bash
rsync -avz --delete -e "ssh -i ~/Projects/ssh-key-2025-02-17.key" ./client/build/ ubuntu@146.235.238.252:/home/ubuntu/Terrariadle/client/build
```

After that, all the files should all be updated with the current ones.


## Hosting The Website

After uploading all the updated files, here is how you start up the site.

### Front-end

Make a new tmux session with the following command:

```bash
$ tmux new -s frontend
```

Once in the session, navigate to the ```/client``` directory and run the following command:

```bash
$ serve -s build
```

### Back-end

As with the previous section, make a new tmux session with the following command:

```bash
$ tmux new -s backend
```

Once in the session, navigate to the ```/api``` directory and run the following command:

```bash
$ node src/DailyGuessAPI.js
```

## Closing Notes

After all of these steps, the site should be updated with the correct information. To assist
with this process, I reccomend to make a script that does this all together.

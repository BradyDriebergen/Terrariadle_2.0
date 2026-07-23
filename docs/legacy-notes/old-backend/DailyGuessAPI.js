const { getRandomWeaponData, getRandomConnectionsData, getRandomNpcData, getRandomEnemyData } = require('./DailyReset');
const mongoose = require('mongoose');
const express = require('express');
const cors = require('cors');
const app = express();
const https = require('https');
// const WebSocket = require('ws');
require('dotenv').config();

const fs = require('fs');
const path = require('path');

const isProduction = process.env.NODE_ENV === 'production';

let server;
if (isProduction) {
  server = https.createServer({
    key: fs.readFileSync(path.join(__dirname, '../certs/privkey.pem')),
    cert: fs.readFileSync(path.join(__dirname, '../certs/fullchain.pem')),
  }, app);
} else {
  server = require('http').createServer(app);
}

// const wss = new WebSocket.Server({ server, path: "/ws" });

// server.on('upgrade', (req, socket) => {
//   console.log('Upgrade request received for:', req.url);
// });

const PORT = process.env.PORT || 3001;
const INTERVAL_DURATION = 24 * 60 * 60 * 1000;
let intervalStartTime = Date.now();

const uri = process.env.MONGODB_URL;

mongoose.connect(uri, { 
    useNewUrlParser: true, 
    useUnifiedTopology: true,
    maxPoolSize: 20
});

const guessSchema = new mongoose.Schema({
    userId: String,
    gtw_guesses: [String],
    gtw_hasWon: Boolean,
    gtw_position: Number,
    gtn_guesses: [String],
    gtn_name_guess: String,
    gtn_hasWon: Boolean,
    gtn_position: Number,
    tc_guessed_categories: [String],
    tc_attempts: Number,
    tc_hasWon: Boolean,
    tc_position: Number,
    hm_guesses: [String],
    hm_attempts: Number,
    hm_position: Number
});

const mongooseModel = process.env.MODEL;
const Guess = mongoose.model(mongooseModel, guessSchema);

app.use(cors());
app.use(express.json());

let randomWeapon = null;
let prevRandomWeapon = null;

let randomNPC = null;
let randomNPCQuote = null;
let randomNPCName = null;
let randomNPCNames = null;

let tc_categories = null;
let tc_groups = null;

let randomEnemy = null;

let gtw_dailyGuesses = 0;
let gtn_dailyGuesses = 0;
let tc_dailyGuesses = 0;
let hm_dailyGuesses = 0;

// const clients = new Map();

const dailyReset = async () => {
    try {
        const deleteResult = await Guess.deleteMany({});
        console.log(`Deleted ${deleteResult.deletedCount} documents from the collection.`);
        
        let weaponData = getRandomWeaponData(randomWeapon);
        randomWeapon = weaponData[0];
        prevRandomWeapon = weaponData[1];
        gtw_dailyGuesses = 0;

        let npcData = getRandomNpcData(randomNPC);
        randomNPC = npcData[0];
        randomNPCQuote = npcData[1];
        randomNPCName = npcData[2];
        randomNPCNames = npcData[3];
        gtn_dailyGuesses = 0;

        let connectionData = getRandomConnectionsData(tc_categories);
        tc_categories = connectionData[0];
        tc_groups = connectionData[1];
        tc_dailyGuesses = 0;

        let enemyData = getRandomEnemyData(randomEnemy);
        randomEnemy = enemyData;
        hm_dailyGuesses = 0;

    } catch (error) {
        console.error('Error deleting documents:', error);
    }
};

app.post('/api/init-user/:mode', async (req, res) => {
    const mode = req.params.mode;
    const { userId } = req.body;
    let user = await Guess.findOne({ userId: userId });

    if (user) {
        if (mode === 'gtw') res.json({ success: user.gtw_hasWon, guesses: user.gtw_guesses });
        if (mode === 'gtn') res.json({ 
            success: user.gtn_hasWon, 
            guesses: user.gtn_guesses, 
            nameGuessed: user.gtn_name_guess,
            name: (user.gtn_name_guess !== "") ? randomNPCName : ""
        });
        if (mode === 'tc') res.json({success: user.tc_hasWon, attempts: user.tc_attempts, guesses: user.tc_guessed_categories});
        if (mode === 'hm') res.json({ guesses: user.hm_guesses, attempts: user.hm_attempts });
    } else {
        if (mode === 'gtw') res.json({ success: false, guesses: [] });
        if (mode === 'gtn') res.json({ success: false, guesses: [], nameGuessed: "" });
        if (mode === 'tc') res.json({ success: false, attempts: 4, guesses: [] });
        if (mode === 'hm') res.json({ guesses: [], attempts: 0 });
    }
})

app.get('/api/guess-data/:mode', (req, res) => {
    const mode = req.params.mode;

    if (mode === 'gtw') {
        return res.json([randomWeapon, prevRandomWeapon]);
    } else if (mode === 'gtn') {
        return res.json([randomNPCQuote, randomNPCNames]);
    } else if (mode === 'tc') {
        return res.json([tc_categories, tc_groups]);
    } else if (mode === 'hm') {
        return res.json([randomEnemy]);
    } else {
        return res.status(400).json({ error: 'Invalid mode' });
    }
});

app.get('/api/players-guessed/:mode', (req, res) => {
    const mode = req.params.mode;

    if (mode === 'gtw') res.json([gtw_dailyGuesses]);
    if (mode === 'gtn') res.json([gtn_dailyGuesses]);
    if (mode === 'tc') res.json([tc_dailyGuesses]);
    if (mode === 'hm') res.json([hm_dailyGuesses]);
});

// wss.on('connection', (ws, req) => {
//     console.log('WebSocket connected from:', req.socket.remoteAddress);

//     ws.on('message', (message) => {
//         try {
//             const { mode } = JSON.parse(message);
//             if (['gtw', 'gtn', 'tc', 'hm'].includes(mode)) {
//                 clients.set(ws, mode);
//                 ws.send(JSON.stringify({ value: getGuessCount(mode) }));
//             }
//         } catch (err) {
//             console.error('Invalid message from client:', err);
//         }
//     });

//     ws.on('close', () => {
//         clients.delete(ws);
//     });
// });

// function getGuessCount(mode) {
//     if (mode === 'gtw') return gtw_dailyGuesses;
//     if (mode === 'gtn') return gtn_dailyGuesses;
//     if (mode === 'tc') return tc_dailyGuesses;
//     if (mode === 'hm') return hm_dailyGuesses;
// }

// // Call this whenever a guess count changes
// function broadcastUpdate(mode) {
//     const value = getGuessCount(mode);

//     for (const [client, clientMode] of clients.entries()) {
//         if (clientMode === mode && client.readyState === WebSocket.OPEN) {
//             client.send(JSON.stringify({ mode, value }));
//         }
//     }
// }

app.get('/api/remaining-time', (req, res) => {
    const currentTime = Date.now();
    const elapsedTime = currentTime - intervalStartTime;
    res.json(Math.floor((INTERVAL_DURATION - (elapsedTime % INTERVAL_DURATION)) / 1000));
});

app.get('/api/get-position/:mode/:userId', async (req, res) => {
    const mode = req.params.mode;
    const userId = req.params.userId;
    
    const user = await Guess.findOne({ userId: userId });
    
    if (!user) {
        return res.json({ pos: -1 });
    }

    if (mode === 'gtw') res.json({ pos: user.gtw_position });
    if (mode === 'gtn') res.json({ pos: user.gtn_position });
    if (mode === 'tc') res.json({ pos: user.tc_position });
    if (mode === 'hm') res.json({ pos: user.hm_position });
})

app.post('/api/guess/:mode', async (req, res) => {
    const mode = req.params.mode;
    const { userId, guess } = req.body;

    const user = await saveGuess(userId, guess, mode);

    if (mode === 'gtw') res.json({ success: user.gtw_hasWon, guesses: user.gtw_guesses });
    if (mode === 'gtn') res.json({ success: user.gtn_hasWon, guesses: user.gtn_guesses });
    if (mode === 'gtn-name') res.json({ name: randomNPCName });
    if (mode === 'tc' || mode === 'hm') res.status(200).json({ message: 'Success' })
});

const saveGuess = async (userID, guess, mode) => {
    let user = await Guess.findOne({ userId: userID });

    if (!user) {
        user = await Guess.create({
            userId: userID,
            gtw_guesses: [],
            gtw_hasWon: false,
            gtw_position: -1,
            gtn_guesses: [],
            gtn_name_guess: "",
            gtn_hasWon: false,
            gtn_position: -1,
            tc_guessed_categories: [],
            tc_position: -1,
            tc_attempts: 4,
            hm_guesses: [],
            hm_attempts: 0,
            hm_position: -1
        });
    }

    if (mode === 'gtw') {
        const newGuesses = [guess, ...user.gtw_guesses];
        const hasWon = guess === randomWeapon.name;
        let updateData = { 
            gtw_guesses: newGuesses, 
            gtw_hasWon: hasWon
        };
        
        if (hasWon) {
            gtw_dailyGuesses++;
            // broadcastUpdate('gtw');
            updateData.gtw_position = gtw_dailyGuesses;
        }
        await Guess.updateOne(
            { userId: userID },
            { $set: updateData }
        );

        user.gtw_guesses = newGuesses;
        user.gtw_hasWon = hasWon;
        user.gtw_position = gtw_dailyGuesses;
    } else if (mode === 'gtn') {
        const newGuesses = [guess, ...user.gtn_guesses];
        const hasWon = guess === randomNPC.npc;
        let updateData = {
            gtn_guesses: newGuesses,
            gtn_hasWon: hasWon
        }

        if (hasWon) {
            gtn_dailyGuesses++;
            // broadcastUpdate('gtn');
            updateData.gtn_position = gtn_dailyGuesses;
        }
        await Guess.updateOne(
            { userId: userID },
            { $set: updateData }
        );

        user.gtn_guesses = newGuesses;
        user.gtn_hasWon = hasWon;
        user.gtn_position = gtn_dailyGuesses;
    } else if (mode === 'gtn-name') {
        await Guess.updateOne(
            { userId: userID },
            { $set: { gtn_name_guess: guess } }
        );

        user.gtn_name_guess = guess;
    } else if (mode === 'tc') {
        let updateData = {};
        if (guess === 'attempt') {
            updateData.tc_attempts = user.tc_attempts - 1;
        } else {
            const newGuesses = [...user.tc_guessed_categories, guess];
            if (newGuesses.length === 4) {
                updateData.tc_hasWon = true;
                tc_dailyGuesses++;
                // broadcastUpdate('tc');
                updateData.tc_position = tc_dailyGuesses;
            }
            updateData.tc_guessed_categories = newGuesses;
        }

        await Guess.updateOne(
            { userId: userID },
            { $set: updateData }
        );
    } else if (mode === 'hm') {
        const newGuesses = [guess, ...user.hm_guesses];
        let updateData = { 
            hm_guesses: newGuesses
        };

        if (!randomEnemy.name.toUpperCase().includes(guess)) {
            updateData.hm_attempts = user.hm_attempts + 1;
        }

        if (containsAllLetters(newGuesses, randomEnemy.name)) {
            hm_dailyGuesses++;
            // broadcastUpdate('hm');
            updateData.hm_position = hm_dailyGuesses;
        }

        await Guess.updateOne(
            { userId: userID },
            { $set: updateData }
        );
    }

    return user;
};

function containsAllLetters(charArray, inputString) {
  // Convert input string to uppercase and extract only letters
  const lettersInString = new Set(
    inputString.toUpperCase().match(/[A-Z]/g)
  );
  
  // Convert the array to a Set (it’s already uppercase)
  const lettersInArray = new Set(
    charArray.filter(c => /[A-Z]/.test(c))
  );

  // Check if every letter in the string is in the array
  for (let letter of lettersInString) {
    if (!lettersInArray.has(letter)) {
      return false;
    }
  }
  return true;
}

dailyReset();
setInterval(() => {
    dailyReset();
    intervalStartTime = Date.now();
}, INTERVAL_DURATION);

server.listen(PORT, '0.0.0.0', () => {
    console.log(`Server running on http://0.0.0.0:${PORT}`);
});

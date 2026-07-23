const weapons = require('../item.json');
const npcs = require('../npcs.json');
const catagories = require('../Catagories.json');
const enemies = require('../enemies.json');

function getRandomWeaponData(weapon) {
    if (!weapon) weapon = weapons[0];

    let newWeapon = null;
    let oldWeapon = weapon;

    do {
        newWeapon = weapons[Math.floor(Math.random() * weapons.length)];
    } while(newWeapon === weapon);

    return [newWeapon, oldWeapon];
}

function getRandomConnectionsData(oldCatagories) {
    let tc_groups = null;
    let tc_catagories = null;
    let shuffledCategories = null;

    // Shuffles all the catagories and picks out the first four
    if (!oldCatagories) {
        shuffledCategories = shuffle(catagories).slice(0, 4);
        tc_catagories = shuffledCategories.map(cat => cat.category);
    } else {
        let duplicate = false;
        do {
            shuffledCategories = shuffle(catagories).slice(0, 4);
            const newCatagories = shuffledCategories.map(cat => cat.category);
            duplicate = oldCatagories.some(item => newCatagories.includes(item))
        } while(duplicate);
        tc_catagories = shuffledCategories.map(cat => cat.category);
    }

    tc_groups = shuffledCategories.map(cat => shuffle(cat.options).slice(0, 4));

    const allItems = tc_groups.flat();
    const uniqueItems = new Set(allItems);

    if (tc_groups.length !== 4 || tc_catagories.length !== 4 || allItems.length !== uniqueItems.size) {
        let data = getRandomConnectionsData(oldCatagories);
        tc_catagories = data[0];
        tc_groups = data[1];
    }

    return [tc_catagories, tc_groups]
}

function getRandomNpcData(oldNpc) {
    let npc = null;

    do {
        npc = npcs[Math.floor(Math.random() * npcs.length)];
    } while (npc === oldNpc);

    let npcQuote = npc.quotes[Math.floor(Math.random() * npc.quotes.length)];
    let npcName = npc.names[Math.floor(Math.random() * npc.names.length)];
    let npcNames = pickRandomNames(npcName);

    return [npc, npcQuote, npcName, npcNames];
}

function pickRandomNames(npcName) {
    let npcList = shuffle(npcs);
    npcList = npcList.filter(n => !n.names.includes(npcName));
    let npcNames = new Set();

    npcNames.add(npcName);

    for (let i = 0; i < 3; i++) {
        let name = npcList[i].names[Math.floor(Math.random() * npcList[i].names.length)];
        npcNames.add(name)
    }

    npcNames = shuffle([...npcNames]);

    return npcNames;
}

function getRandomEnemyData(enemy) {
    let randomEnemy = enemies[Math.floor(Math.random() * enemies.length)];

    do {
        randomEnemy = enemies[Math.floor(Math.random() * enemies.length)];
    } while (enemy === randomEnemy);

    return randomEnemy;
}

function shuffle(array) {
    let arr = [...array];
    for (let i = arr.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [arr[i], arr[j]] = [arr[j], arr[i]];
    }
    return arr;
}


module.exports = {getRandomWeaponData, getRandomNpcData, getRandomConnectionsData, getRandomEnemyData}
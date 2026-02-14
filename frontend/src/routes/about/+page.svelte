<script lang="ts">
    function handleScroll(event: MouseEvent, id: string) {
        event.preventDefault();
        const element = document.getElementById(id);
        if (element) element.scrollIntoView({ behavior: "smooth" });
    }
</script>

<div class="how-to-play">
    <div class="directory">
        <h3>How to Play</h3>
        <nav>
            <a class="link" href="#dw" on:click={(e) => handleScroll(e, "dw")}>Daily Weapon</a>
            <a class="link" href="#tc" on:click={(e) => handleScroll(e, "tc")}>Terraria Connections</a>
            <a class="link" href="#gtn" on:click={(e) => handleScroll(e, "gtn")}>Guess The NPC</a>
            <a class="link" href="#hm" on:click={(e) => handleScroll(e, "hm")}>Hangman</a>
        </nav>
    </div>

    <h1>About</h1>
    <hr />

    <p>
        Welcome to Terrariadle! This is a passion project based on the popular sandbox
        game,
        <a class="link" href="https://terraria.org/" target="_blank" rel="noopener noreferrer">Terraria</a>.
        I appreciate you stopping by and taking the time to read this.
    </p>

    <p>
        I'm a solo developer that loved playing daily games such as
        <a class="link" href="https://www.nytimes.com/games/wordle/index.html" target="_blank" rel="noopener noreferrer">Wordle</a>,
        <a class="link" href="https://loldle.net/" target="_blank" rel="noopener noreferrer">Loldle</a>,
        <a class="link" href="https://www.nytimes.com/games/connections" target="_blank" rel="noopener noreferrer">Connections</a>,
        <a class="link" href="https://puzzles.usatoday.com/" target="_blank" rel="noopener noreferrer">Crossword</a>,
        etc. I wanted to make a version of these daily guessing games based on my all-time
        favorite game, Terraria.
    </p>

    <p>
        This project aims to remake another developer's project:
        <a class="link" href="https://www.terradle.com/" target="_blank" rel="noopener noreferrer">Terradle</a>.
        They did an amazing job on his site but I wanted to expand and improve on the idea.
        Please check out their
        <a class="link" href="https://github.com/cxhuy" target="_blank" rel="noopener noreferrer">github</a>
        to see some of their other projects.
    </p>

    <p>
        Found a bug <img src="/emojis/Bug.png" alt="" class="emoji-icons" />, have a suggestion
        <img src="/emojis/Suggestion.png" alt="" class="emoji-icons" />, or you want to say hi
        <img src="/emojis/Heart.png" alt="" class="emoji-icons" /> ?
    </p>

    <p>
        If you'd like to support the game, you can
        <a class="link" href="https://ko-fi.com/bytebeard" target="_blank" rel="noopener noreferrer">
        buy me a cup of coffee!
        </a>
        <img src="/emojis/Coffee.png" alt="" id="coffee" />
    </p>

    <h2 id="dw">How to play: Daily Weapon</h2>
    <hr />
    <p>
        The daily weapon is a random weapon found in Terraria picked every 24 hours. The daily
        weapon is picked from a pool of ~450 weapons. Users can then try to guess the daily weapon
        using hints given by previous guesses. The game is won by guessing the correct weapon.
    </p>
    <p>
        Guesses are displayed below the input prompt. Each guess contains traits specific to the guessed
        weapon. Each trait can have one of three backgrounds:
    </p>
    <ul>
        <li><span style="color: var(--color-red);">Red</span> - Trait is not the same as the daily weapon.</li>
        <li><span style="color: var(--color-yellow);">Yellow</span> - Trait is partialy similar to the daily weapon.</li>
        <li><span style="color: var(--color-green);">Green</span> - Trait is the same as the daily weapon.</li>
    </ul>

    <h5>
        <em>
        *Note: Some weapons traits match exactly with others. If you get all the traits right, it doesn't
        mean you won. It means that there is another weapon with the exact same traits.
        </em>
    </h5>

    <p>
        You know you've won when a prompt appears above the guesses displaying "You Got It!" The prompt will also
        display the remaining time until the daily weapon refreshes.
    </p>

    <h4>Extra notes about the game:</h4>
    <ul style="font-size: 16px;">
        <li>- Melee weapons that are thrown count at throwables (hint).</li>
        <li>- Lances count as spears</li>
        <li>- Repeaters and crossbows count as bows</li>
        <li>- Specialty guns count as guns</li>
        <li>- Staffs and wands are grouped together to avoid confusion</li>
        <li>- Wooden swords and phaseblades just count as phaseblade and wooden sword/bow</li>
    </ul>

    <h2 id="tc">How to play: Terraria Connections</h2>
    <hr />
    <p>
        This game is based heavily on the NYT game:
        <a class="link" href="https://www.nytimes.com/games/connections" target="_blank" rel="noopener noreferrer">Connections</a>.
        The goal of this game is to find groups of four objects that share something in common. Once you've selected four
        objects, click submit to check if your guess is correct. You have to find all four groups without making four
        mistakes.
    </p>
    <p>
        Because this is based on Terraria, each catagory will be based on different parts of the game. This can range from
        accessories, potions, fish, weapons, biomes, etc. However, catagories are not as straightforward as they seem. Use
        your best judgement to find the similarities between the objects.
    </p>

    <h3>Examples of catagories:</h3>
    <ul style="margin-top: -10px;">
        <li>- Slime Bonus Drops</li>
        <li>- Types of Sharks</li>
        <li>- Snow Critters</li>
        <li>- Underground Plants</li>
        <li>- Rarity Colors</li>
    </ul>

    <p>
        Unlike Connections, this game doesn't have difficulty tiers. Each catagory ranges in difficulty, so you can have days
        where the solution is obvious, and others where its more challenging.
    </p>

    <h2 id="gtn">How to play: Guess The NPC</h2>
    <hr />
    <p>
        Terraria's NPCs say all kinds of things. Everyday, a random quote said by a NPC is picked and you are prompted
        to guess who said it. The quote can be from any NPC found in the game excluding Santa Claus. A full list of
        NPCs can be found
        <a class="link" href="https://terraria.fandom.com/wiki/NPCs" target="_blank" rel="noopener noreferrer">here</a>.
        You can guess through entering the trade of the NPC (Guide, Dryad, Arms Dealer, etc.). You will know when your right
        when the winning screen appears.
    </p>

    <p>
        When you guess correctly, you will be prompted to play a mini-game. Like NPC quotes, many NPCs have unique names.
        The mini-game will task you with guessing the name of the Daily NPC. You will be prompted with four NPC names.
        One of the names is a name the Daily NPC can have. After you have made your guess, the correct guess will appear.
        If you guessed correctly, you will be shown the message "Nailed it!". If not, it will will show you the correct guess
        and say "Better Luck Next Time!". Be careful, you only have one guess.
    </p>

    <h2 id="hm">How to play: Hangman</h2>
    <hr />
    <p>
        This gamemode is a twist on the classic game; Hangman. The goal of this game is to guess the name of a random enemy
        before you hang the guide. You can guess by selecting letters on the on-screen keyboard. If the letter is within the
        enemy's name, it will show up on the position of the letter(s). If the letter is not in the name, it will add a limb to
        the guide. If you add all 6 body parts to the guide, you lose and the correct name is revealed. The enemies can consist
        of any enemy found in the game and bosses.
    </p>

    <p>
        Some helpful tips to get started is to guess vowels first. Pay attention to punctuation. More often than not, words that
        end with '_' usually end in the letter s. If your stuck after that, guess common letters such as h, t, s, r, etc. Also
        remember that a lot of creatures in the game are zombies, slimes, and skeletons. The more you play, the more you'll develop
        patterns that can help you. Good luck guessing Terrarian!
    </p>
</div>

<style>
    .directory {
        padding: 10px 20px;
        margin: auto;
        margin-top: -45px;
        width: fit-content;

        background-color: var(--color-lightblue);
        border: 2px solid black;
        border-radius: 25px;
    }
    .directory h3 {
        margin: 5px;
    }
    .directory nav {
        display: flex;
        gap: 20px;
    }
    .directory a {
        font-size: 18px;
    }
    @media (max-width: 600px) {
        .directory a {
            font-size: 14px;
        }
    }

    .link {
        color: rgb(111, 255, 255);
    }
    .link:hover {
        color: yellow;
    }

    .how-to-play {
        background-color: var(--color-button);
        text-align: left;
        margin: 20px auto;
        margin-top: 60px;
        max-width: 600px;
        padding: 1px 15px 1px 15px;
        font-size: 20px;

        border: 2px solid #000;
        border-radius: 15px;
        background-color: var(--color-button);
    }
    .how-to-play h1 {
        margin-bottom: -5px;
    }
    .how-to-play h2 {
        padding-top: 20px;
        margin-bottom: -5px;
    }
    .how-to-play ul {
        list-style-type: none;
        margin: 0;
        padding: 0;
    }

    .emoji-icons {
        max-width: 15px;
        max-height: 15px;
        width: 15px;
        height: 15px;
        object-fit: contain;
    }
</style>

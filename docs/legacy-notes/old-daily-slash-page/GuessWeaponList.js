import styles from "./GuessWeaponList.module.css"

export default function GuessWeaponList(guess, weapon) {

    const guessCorrect = { backgroundColor: 'rgb(61, 212, 41)' };
    const guessPartial = { backgroundColor: 'yellow' };
    const guessWrong = { backgroundColor: 'red' };

    const guessObtained = (g, w) => {
        if (JSON.stringify(g) === JSON.stringify(w)) {
            return guessCorrect;
        } else if (Array.isArray(w) && w.includes(g)) {
            return guessPartial
        } else if (Array.isArray(g) && g.includes(w)) {
            return guessPartial
        } else {
            return guessWrong
        }
    }

    const guessTrueFalse = (g, w) => {
        if (g === w) {
            return guessCorrect;
        } else {
            return guessWrong;
        }
    }
    const guessHighLowDmg = (g, w) => {
        if (g < w) {
            return ' ↑';
        } else if (g > w) {
            return ' ↓';
        } else {
            return '';
        }
    }

    const guessHighLowRar = (g, w) => {
        const rarities = ['White', 'Blue', 'Green', 'Orange', 'Light_Red', 'Pink', 
            'Light_Purple', 'Lime', 'Yellow', 'Cyan', 'Red'];

        if (rarities.indexOf(g) < rarities.indexOf(w)) {
            return '  ↑';
        } else if (rarities.indexOf(g) > rarities.indexOf(w)) {
            return '  ↓';
        } else {
            return '';
        }
    }

    const guessHighLowUse = (g, w) => {
        const useTimes = ['Snail', 'Extremely Slow', 'Very Slow', 'Slow', 'Average', 
                        'Fast', 'Very Fast', 'Insanely Fast'];

        if (useTimes.indexOf(g) < useTimes.indexOf(w)) {
            return '  ↑';
        } else if (useTimes.indexOf(g) > useTimes.indexOf(w)) {
            return '  ↓';
        } else {
            return '';
        }
    }

    const displayObtained = (msg) => {
        if (Array.isArray(msg)) {
            return msg.join(' ');
        } else {
            return msg;
        }
    }

    return (
        <div className={styles["guess-row"]}>
            <div>
                <img src={guess.info['image-path'].replace(".", "assets/")} alt='' />
            </div>
            <div style={guessTrueFalse(guess.info['damage-type'], weapon.info['damage-type'])}>
                {guess.info['damage-type']}
            </div>
            <div style={guessTrueFalse(guess.info['damage'], weapon.info['damage'])}>
                {guess.info['damage'] + guessHighLowDmg(guess.info['damage'], weapon.info['damage'])}
            </div>
            <div style={guessTrueFalse(guess.info['use-time'], weapon.info['use-time'])}>
                {guess.info['use-time'] + guessHighLowUse(guess.info['use-time'], weapon.info['use-time'])}
            </div>
            <div style={guessTrueFalse(
                guess.info['rarity'].replace('./rarities/', '').replace('.png', ''),
                weapon.info['rarity'].replace('./rarities/', '').replace('.png', '')
                )}>
                <img src={guess.info['rarity'].replace(".", "assets/")} alt='' />
                <p style={{display:"inline-block",paddingLeft:"3px"}}> {guessHighLowRar(
                    guess.info['rarity'].replace('./rarities/', '').replace('.png', ''),
                    weapon.info['rarity'].replace('./rarities/', '').replace('.png', '')
                )}</p>
            </div>
            <div style={guessTrueFalse(guess.info['operation'], weapon.info['operation'])}>
                {guess.info['operation']}
            </div>
            <div style={guessTrueFalse(guess.info['material'], weapon.info['material'])}>
                {guess.info['material']}
            </div>
            {/* Make it so it's separated by space/slash */}
            <div style={guessObtained(guess.info['obtained'], weapon.info['obtained'])}>
                {displayObtained(guess.info['obtained'])}
            </div>
        </div>
    )
}
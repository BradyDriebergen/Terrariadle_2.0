import styles from "./IntroScreen.module.css"

export default function IntroScreen(guessCount, prevWeapon) {   
    
    const fontColor = (p) => {
        const colors = {
            White: '#ffffff', 
            Blue: '#1732ff', 
            Green: '#35d400', 
            Orange: '#ffa600', 
            Light_Red: '#ff8080', 
            Pink: '#ffbdf0', 
            Light_Purple: '#de28b3', 
            Lime: '#69f75c',
            Yellow: '#fbff00', 
            Cyan: '#00fff7', 
            Red: '#ff0000'
        };
        const color = colors[p.info['rarity'].replace('./rarities/', '').replace('.png', '')];

        return {borderColor: color, color: color};
    }

    return (
        <div className={styles.container}>
            <p>{guessCount} people have guessed today's weapon!</p>
            <h3>Yesterdays weapon was: </h3>
            <div style={fontColor(prevWeapon)}>
                <img src={prevWeapon.info['image-path'].replace(".", "assets/")} alt=''/>
            </div>
            <h3 style={fontColor(prevWeapon)}>{prevWeapon.name}</h3>
            <p>Guess any weapon to begin.</p>
        </div>
    );
}
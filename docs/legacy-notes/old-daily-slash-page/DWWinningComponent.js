import { useEffect, useState } from "react";
import CustomConfetti from "../ConfettiExplosion";
import GetRemainingTimeComponent from "../GetRemainingTimeComponent"; // Ensure the correct path
import { motion } from "framer-motion";
import axios from "axios";
import styles from "./DWWinningComponent.module.css";

const ip = process.env.REACT_APP_LOCAL_IP || '';

export default function WinningComponent({ guesses, randomWeapon, remainingTime, user }) { 
    const [pos, setPos] = useState(0);

    useEffect(() => {
        fetchPosition();
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    const fetchPosition = async () => {
        try {
            const response = await axios.get(ip + '/api/get-position/gtw/' + user);
            setPos(response.data.pos);
        } catch (error) {
            console.error('Error fetching random weapon:', error);
        }
    }

    const backgroundAndBorder = (p) => {
        let background = null;
        let border = null;

        switch (p.info['rarity'].replace('./rarities/', '').replace('.png', '')) {
            case 'White':
                background = 'url(\'/backgrounds/WoodWall.png\')';
                border = 'url(\'/backgrounds/Wood.png\')';
                break;
            case 'Blue':
                background = 'url(\'/backgrounds/StoneWall.png\')';
                border = 'url(\'/backgrounds/StoneBlock.png\')';
                break;
            case 'Green':
                background = 'url(\'/backgrounds/DungeonWall.png\')';
                border = 'url(\'/backgrounds/DungeonBrick.png\')';
                break;
            case 'Orange':
                background = 'url(\'/backgrounds/HellstoneWall.png\')';
                border = 'url(\'/backgrounds/HellstoneBrick.png\')';
                break;
            case 'Light_Red':
                background = 'url(\'/backgrounds/PearlstoneWall.png\')';
                border = 'url(\'/backgrounds/PearlstoneBlock.png\')';
                break;
            case 'Pink':
                background = 'url(\'/backgrounds/CrystalBlockWall.png\')';
                border = 'url(\'/backgrounds/CrystalBlock.png\')';
                break;
            case 'Light_Purple':
                background = 'url(\'/backgrounds/ChlorophyteBrickWall.png\')';
                border = 'url(\'/backgrounds/ChlorophyteBrick.png\')';
                break;
            case 'Lime':
                background = 'url(\'/backgrounds/LihzahrdBrickWall.png\')';
                border = 'url(\'/backgrounds/LihzahrdBrick.png\')';
                break;
            case 'Yellow':
                background = 'url(\'/backgrounds/MartianConduitWall.png\')';
                border = 'url(\'/backgrounds/MartianConduitPlating.png\')';
                break;
            case 'Cyan':
                background = 'url(\'/backgrounds/SmoothMarbleWall.png\')';
                border = 'url(\'/backgrounds/SmoothMarbleBlock.png\')';
                break;
            case 'Red':
                background = 'url(\'/backgrounds/LuminiteBrickWall.png\')';
                border = 'url(\'/backgrounds/LuminiteBrick.png\')';
                break;
            default:
                background = 'url(\'/backgrounds/WoodWall.png\')';
                border = 'url(\'/backgrounds/Wood.png\')'
                break;
        }

        return { 
            border: '20px solid transparent',
            borderImage: border, 
            borderImageSlice: '17',
            borderImageRepeat: 'round',
            borderRadius: '5px',

            backgroundImage: background,
            backgroundRepeat: 'repeat',
            backgroundSize: '20%'
        };
    };
    
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

        return { borderColor: color, color: color };
    }

    const getGuessCount = (n) => {
        switch (n) {
            case 1:
                return '1st';
            case 2:
                return '2nd';
            case 3:
                return '3rd';
            default:
                return n + 'th'
        }
    };

    return (
        <div>
            <motion.div 
                className={styles.container} 
                style={backgroundAndBorder(randomWeapon)}
                initial={{ opacity: 0, scale: 0.5 }}
                animate={{ 
                    opacity: 0.9, 
                    scale: 1,
                    x: [0, -5, 10, -10, 5, 0]
                }}
                transition={{ 
                    duration: 0.8,
                    type: "spring",
                    stiffness: 260,
                    damping: 20
                }}
            >
                <h1>You Got It!</h1>
                <p>You were the {getGuessCount(pos)} person to guess today's weapon!</p>
                <div className={styles.weapon} style={fontColor(randomWeapon)}>
                    <img className={styles.weaponImg} src={randomWeapon.info['image-path'].replace(".", "assets/")} alt=''/>
                </div>
                <h3 style={fontColor(randomWeapon)}>{randomWeapon.name}</h3>
                <p>{guesses} people have guessed today's weapon!</p>
                <GetRemainingTimeComponent remainingTime={remainingTime} numGuessed={guesses} />
            </motion.div>
            <CustomConfetti />
        </div>
    );
}
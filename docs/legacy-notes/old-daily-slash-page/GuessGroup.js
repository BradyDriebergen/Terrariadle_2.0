import GuessWeaponList from "./GuessWeaponList";
import weapons from "./item.json";
import { AnimatePresence, motion } from "framer-motion";
import styles from "./GuessGroup.module.css"

export default function GuessGroup(guesses, randomWeapon) {
    const result = guesses.map(name => weapons.find(weapon => weapon.name === name)).filter(Boolean);
    return (
        <div className={styles.wrapper}>
            <div className={styles.container}>
                <div className={styles.categories}>
                    <div className={styles.category} title="Image of Weapon">Weapon</div>
                    <div className={styles.category} title="Melee, Ranged, Magic, Summon, Throwing">Damage Type</div>
                    <div className={styles.category} title="Weapon Damage">Damage</div>
                    <div className={styles.category} title="Snail, Extremely Slow, Very Slow, Slow, Average, 
                                                    Fast, Very Fast, Insanely Fast">Use Time</div>
                    <div className={styles.category} title="White, Blue, Green, Orange, Light_Red, Pink, 
                                                    Light_Purple, Lime, Yellow, Cyan, Red">Rarity</div>
                    <div className={styles.category} title="Auto or Manual">Operation</div>
                    <div className={styles.category} title="Yes or No">Material</div>
                    <div className={styles.category} title="Crafting, Chest, Buy, Drop, Fishing, Background Object">Obtained</div>
                </div>
                <ul>
                    <AnimatePresence initial={false}>
                    {result.map((guess, index) => (
                        index === 0 ? (
                            <motion.li
                                key={guess.name}
                                initial={{ opacity: 0, scale: 0.9 }}
                                animate={{ 
                                    opacity: 0.9, 
                                    scale: 1,
                                    x: [0, -5, 10, -10, 5, 0]
                                }}
                                transition={{ 
                                    duration: 1,
                                    type: "spring",
                                    stiffness: 260,
                                    damping: 40
                                }}
                            >
                                {GuessWeaponList(guess, randomWeapon)}
                            </motion.li>
                        ) : (
                            <li key={guess.name}>
                                {GuessWeaponList(guess, randomWeapon)}
                            </li>
                        )
                    ))}
                    </AnimatePresence>
                </ul>
            </div>
        </div>
    );
}
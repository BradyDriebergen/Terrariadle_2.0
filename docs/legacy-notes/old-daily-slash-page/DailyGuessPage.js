import {useEffect, useState, useRef} from 'react';
import weapons from './item.json';
import GuessGroup from './GuessGroup';
import IntroScreen from './IntroScreen';
import axios from 'axios';
import WinningComponent from './DWWinningComponent';
import styles from "./DailyGuessPage.module.css"
import SEO from '../SEO'

const { v4: uuidv4 } = require('uuid');
const ip = process.env.REACT_APP_LOCAL_IP || '';
// const swLink = process.env.REACT_APP_WS_LINK;

export default function DailyGuessPage() {
    const [options, setOptions] = useState([]);
    const [input, setInput] = useState("");
    const [guesses, setGuesses] = useState([]);
    const [randomWeapon, setRandomWeapon] = useState(null);
    const [prevRandomWeapon, setPrevRandomWeapon] = useState(null);
    const [hintOne, setHintOne] = useState("");
    const [hintTwo, setHintTwo] = useState("");
    const [hintThree, setHintThree] = useState("");
    const [hintOneRevealed, setHintOneRevealed] = useState(false);
    const [hintTwoRevealed, setHintTwoRevealed] = useState(false);
    const [hintThreeRevealed, setHintThreeRevealed] = useState(false);
    const [dropdownVisible, setDropdownVisible] = useState(false);
    const [selectedIndex, setSelectedIndex] = useState(-1);
    const [hasWon, setHasWon] = useState(false);
    const [guessTotalCount, setTotalGuessCount] = useState(0);
    const [remainingTime, setRemainingTime] = useState(0);
    const [loading, setLoading] = useState(true);
    const itemRefs = useRef([]);
    const inputRef = useRef(null);

    // Initial Fetch for webpage
    useEffect(() => {
        getUserID();
        getGuesses();
        fetchRandomWeapon();
        getTotalGuesses();
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    const getUserID = async () => {
        let userId = getCookie('userId');
    
        if (!userId) {
            userId = uuidv4(); // Generate a new user ID if not found
        }
    
        // Always refresh expiration time on every visit
        setCookie('userId', userId); 
    
        return userId;
    };

    const getCookie = (name) => {
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return parts.pop().split(';').shift();
    }

    const setCookie = (name, value) => {
        const expires = new Date(Date.now() + 86400000).toUTCString();
        document.cookie = `${name}=${value}; expires=${expires}; path=/`;
    }

    const getGuesses = async () => {
        const response = await axios.post(ip + '/api/init-user/gtw', { 
            userId: getCookie('userId')
        });
        setHasWon(response.data['success']);
        setGuesses(response.data['guesses']);
    };

    const checkGuess = async (guess) => {
        const response = await axios.post(ip + '/api/guess/gtw', { 
            userId: getCookie('userId'),
            guess: guess
        });
        setHasWon(response.data['success']);
        setGuesses(response.data['guesses']);
    }

    const fetchRandomWeapon = async () => {
        try {
            const response = await axios.get(ip + '/api/guess-data/gtw');
            setRandomWeapon(response.data[0]);
            setPrevRandomWeapon(response.data[1]);
            setLoading(false);
        } catch (error) {
            console.error('Error fetching random weapon:', error);
        }
    }

    const getTotalGuesses = async () => {
        try {
            const response = await axios.get(ip + '/api/players-guessed/gtw');
            setTotalGuessCount(response.data[0]);
        } catch (error) {
            console.error('Error fetching players guessed:', error);
        }
    }

    // useEffect(() => {
    //     const ws = new WebSocket(swLink);
    //     ws.onerror = (e) => console.error("WebSocket error", e);

    //     ws.onopen = () => {
    //         ws.send(JSON.stringify({ mode: 'gtw' }));
    //     };

    //     ws.onmessage = (event) => {
    //         try {
    //             const { value } = JSON.parse(event.data);
    //             setTotalGuessCount(value);
    //         } catch (err) {
    //             console.error('Error parsing WebSocket message:', err);
    //         }
    //     };

    //     return () => {
    //         ws.close();
    //     };
    // }, []);

    useEffect(() => {
        if (guesses.length > 0) {
            if (!hintOneRevealed)
                setHintOne("Mode Obtained" + ((guesses.length >= 4) ? ": Click to reveal" : " in " + (4 - guesses.length) + " Tries"));
            if (!hintTwoRevealed)
                setHintTwo("Weapon Type" + ((guesses.length >= 7) ? ": Click to reveal" : " in " + (7 - guesses.length) + " Tries"));
            if (!hintThreeRevealed)
                setHintThree("Image Clue" + ((guesses.length >= 12) ? ": Click to reveal" : " in " + (12 - guesses.length) + " Tries"));
        } else {
            setHintOne('Mode Obtained');
            setHintTwo('Weapon Type');
            setHintThree('Image Clue');
        }
    }, [guesses.length, hintOneRevealed, hintThreeRevealed, hintTwoRevealed])

    useEffect(() => {
        if (input === '') {
            setOptions([]);
            setDropdownVisible(false);
        } else {
            const filtered = weapons
                .filter(option => !guesses.some(guess => guess === option.name))
                .filter(option => 
                    option.name.toLowerCase().includes(input.toLowerCase())
                );
            
            setOptions(filtered);
        }
    }, [input, guesses]);

    useEffect(() => {
        if (selectedIndex >= 0 && selectedIndex < options.length) {
            itemRefs.current[selectedIndex]?.scrollIntoView({
                behavior: 'smooth',
                block: 'nearest'
            });
        }
    }, [selectedIndex, options]);

    // Handles input change to list values
    const handleInputChange = (e) => {
        setInput(e.target.value);
        setDropdownVisible(true);
        setSelectedIndex(-1);
    }

    // Handles Tab functionality to autofill input
    const handleKeyDown = (e) => {
        if (e.key === 'Tab' && options.length > 0) {
            e.preventDefault();
            setInput(options[0].name);
        } else if (e.key === 'Enter') {
            if (selectedIndex >= 0 && selectedIndex < options.length) {
                setInput(options[selectedIndex].name);
                setDropdownVisible(false);
            }  
            if (options.find(option => option.name === input)) {
                handleGuess();
            }
        } else if (e.key === 'ArrowDown') {
            e.preventDefault();
            setSelectedIndex((prevIndex) => (prevIndex + 1) % options.length);
        } else if (e.key === 'ArrowUp') {
            e.preventDefault();
            setSelectedIndex((prevIndex) => (prevIndex - 1 + options.length) % options.length);
        }
    }

    const handleOptionClick = (option) => {
        setInput(option.name);
        setDropdownVisible(false);
        inputRef.current.focus();
    };

    const handleGuess = async () => {
        try {
            checkGuess(input);
            setInput("");
        } catch (error) {
            console.error('Error alerting:', error);
        }
    }

    const hasWonRef = useRef(false);

    useEffect(() => {
        if (hasWon && !hasWonRef.current && randomWeapon) {
            hasWonRef.current = true;
            setHintOneRevealed(true);
            setHintOne(randomWeapon['mode-obtained']);
            setHintTwoRevealed(true);
            setHintTwo(randomWeapon['weapon-type']);
            setHintThreeRevealed(true);
            setHintThree(<img src={randomWeapon.info['image-path'].replace(".", "assets/")} alt='' />);
        }
    }, [hasWon, randomWeapon]);

    useEffect(() => {
        const fetchAndSetRemainingTime = async () => {
            await fetchRemainingTime();
        };
    
        fetchAndSetRemainingTime();
    
        const interval = setInterval(() => {
            setRemainingTime(prevTime => {
                if (prevTime < 1) {
                    clearInterval(interval);
                    if (!hasWon) {
                        alert("Daily weapon changed! Refresh the page to start guessing.")
                    }
                    return 0;
                }
                return prevTime - 1;
            });
        }, 1000);
    
        return () => clearInterval(interval);
    }, [hasWon]);

    const fetchRemainingTime = async () => {
        try {
            const response = await axios.get(ip + '/api/remaining-time');
            console.log("remaining-time api called");
            setRemainingTime(response.data);
        } catch (error) {
            console.error('Error fetching remaining time:', error)
        }
    }

    if (loading) {
        return <div>Loading...</div>
    }

    return (
        <div>
            <SEO />
            <div className={styles.container}>
                <div className={styles.gtw_message}>Guess today's weapon!</div>
                <div className={styles.hint_buttons}>
                    <button 
                        className={styles.hint_button}
                        onClick={() => {
                            setHintOneRevealed(!hintOneRevealed);
                            setHintOne(randomWeapon['mode-obtained']);
                        }}
                        disabled={guesses.length < 4 || hasWon}
                        >{hintOne}</button>
                    <button 
                        className={styles.hint_button}
                        onClick={() => {
                            setHintTwoRevealed(!hintTwoRevealed);
                            setHintTwo(randomWeapon['weapon-type']);
                        }}
                        disabled={guesses.length < 7 || hasWon}
                        >{hintTwo}</button>
                    <button 
                        className={styles.hint_button}
                        onClick={() => {
                            setHintThreeRevealed(!hintThreeRevealed)
                            setHintThree(<img src={randomWeapon.info['image-path'].replace(".", "assets/")} alt='' />);
                        }}
                        disabled={guesses.length < 12 || hasWon}
                        >{hintThree}</button>
                </div>
                {(!hasWon && remainingTime > 1) &&
                    <div className={styles.input_section}>
                        <input 
                            type="text" 
                            value={input} 
                            onChange={handleInputChange} 
                            onKeyDown={handleKeyDown}
                            list="weapon-options" 
                            placeholder="Type any weapon to guess..."
                            ref={inputRef}
                        />
                        <button onClick={() => {
                            if (options.find(option => option.name === input)) {
                                handleGuess();
                            }
                        }}>
                            <img src='SaveButton.png' alt='Enter' />
                        </button>
                    </div>
                }
            </div>
            {dropdownVisible && (
                <div className={styles.dropdown}>
                    <ul>
                        {options.length === 0 && 
                            <li>No Results</li>
                        }
                        {options.map((option, index) => (
                            <li 
                                key={index} 
                                onClick={() => handleOptionClick(option)}
                                className={index === selectedIndex ? styles.option_highlight : ''}
                                ref={el => itemRefs.current[index] = el}
                            >
                                <img src={option.info['image-path'].replace(".", "assets/")} alt='' />
                                {option.name}
                            </li>
                        ))}
                    </ul>
                </div>
            )}
            {(hasWon && randomWeapon) && 
                <WinningComponent 
                    guesses={guessTotalCount} 
                    randomWeapon={randomWeapon} 
                    remainingTime={remainingTime} 
                    user={getCookie('userId')}
                />
            }
            {(guesses.length !== 0 && randomWeapon) ? (
               GuessGroup(guesses, randomWeapon)
            ) : (
                prevRandomWeapon ? IntroScreen(guessTotalCount, prevRandomWeapon) : <div>Loading...</div>
            )}
        </div>
    );
}

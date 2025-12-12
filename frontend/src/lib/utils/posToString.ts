export function ConvertPositionToString(pos: number): string {
    const absPos = Math.abs(pos);

    const lastTwoDigits = absPos % 100;
    if (lastTwoDigits >= 11 && lastTwoDigits <= 13) {
        return `${pos}th`;
    }

    switch (absPos % 10) {
        case 1:
            return `${pos}st`;
        case 2:
            return `${pos}nd`;
        case 3:
            return `${pos}rd`;
        default:
            return `${pos}th`;
    }
}
// Find the smallest positive integer that does not occur in a given sequence.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    const checkArr = [];
    for (let n of A) {
        if (n > 0) {
            checkArr[n - 1] = 1;
        }
    }
    for (let i = 0; i < checkArr.length; i++) {
        if (checkArr[i] === undefined) {
            return i + 1;
        }
    }
    return checkArr.length < 1 ? 1 : checkArr.length + 1;
}
// Compute number of distinct values in an array.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    const map = {};
    let count = 0;
    for (let n of A) {
        map[n] = 0;
    }
    for (let i in map) {
        count++;
    }
    return count;
}
// Find longest sequence of zeros in binary representation of an integer.

function solution(N) {
    // write your code in JavaScript (Node.js 8.9.4)
    const bin = N.toString(2);
    const arr = bin.split('1');
    let maxGap = 0;
    for (let i = 1; i < arr.length - 1; i++) {
        maxGap = Math.max(maxGap, arr[i].length);
    }
    return maxGap;
}
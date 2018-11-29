// Given a log of stock prices compute the maximum possible earning.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    let currMin = A[0], currMax = 0;
    for (let i = 1; i < A.length; i++) {
        currMin = Math.min(A[i], currMin);
        currMax = Math.max(currMax, A[i] - currMin);
    }
    return currMax;
}
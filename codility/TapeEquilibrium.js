// Minimize the value |(A[0] + ... + A[P-1]) - (A[P] + ... + A[N-1])|.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    const cumulativeSum = [];
    A.reduce((acc, curr, i) => {
        return cumulativeSum[i] = curr + acc;
    }, 0);
    const totalSum = cumulativeSum[A.length - 1];
    let minDiff = Number.POSITIVE_INFINITY;
    let movingDiff;
    for (let i = 0; i < A.length - 1; i++) {
        minDiff = Math.min(minDiff, Math.abs(totalSum - 2 * cumulativeSum[i]));
    }
    return minDiff;
}

// Find a maximum sum of a compact subsequence of array elements.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    let movingSum = Number.NEGATIVE_INFINITY;
    let maxSum = Number.NEGATIVE_INFINITY;
    for (let n of A) {
        movingSum = Math.max(n, n + movingSum);
        maxSum = Math.max(maxSum, movingSum);
    }
    return maxSum;
}
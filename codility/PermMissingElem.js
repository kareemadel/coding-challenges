// Find the missing element in a given permutation.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    const N = A.length + 1;
    const completeSum = N * (N + 1) / 2;
    const arrSum = A.reduce((acc, curr) => acc + curr, 0);
    return completeSum - arrSum;
}
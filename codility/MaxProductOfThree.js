// Maximize A[P] * A[Q] * A[R] for any triplet (P, Q, R).

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    const sorted = [...A].sort((a, b) => a - b);
    const prod1 = sorted[0] * sorted[1] * sorted[A.length - 1];
    const prod2 = sorted[A.length - 1] * sorted[A.length - 2] * sorted[A.length - 3];
    return Math.max(prod1, prod2);
}
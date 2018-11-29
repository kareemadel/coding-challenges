// Find the maximal sum of any double slice.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    const sumToI = Array(...Array(A.length)).map(()=> 0);
    const sumFromI = Array(...Array(A.length)).map(()=> 0);
    for (let i = 1; i < A.length - 1; i++) {
        sumToI[i] = Math.max(sumToI[i - 1] + A[i], A[i], 0);
    }
    for (let i = A.length - 2; i > 0; i--) {
        sumFromI[i] = Math.max(sumFromI[i + 1] + A[i], A[i], 0);
    }
    let max = Number.NEGATIVE_INFINITY;
    for (let i = 1; i < A.length - 1; i++) {
        max = Math.max(max, sumToI[i - 1] + sumFromI[i + 1]);
    }
    return max;
}
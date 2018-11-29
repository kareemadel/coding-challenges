// Rotate an array to the right by a given number of steps.

function solution(A, K) {
    // write your code in JavaScript (Node.js 8.9.4)
    const shift = K % A.length;
    if (shift === 0) {
        return A;
    }
    const rem = A.length - shift;
    const result = [];
    for (let i = rem; i < A.length; i++) {
        result.push(A[i]);
    }
    for (let i = 0; i < rem; i++) {
        result.push(A[i]);
    }
    return result;
}
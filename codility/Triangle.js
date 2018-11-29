// Determine whether a triangle can be built from a given set of edges.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    if (A.length < 3) {
        return 0;
    }
    const sorted = [...A].sort((a, b) => a - b);
    for (let i = 0; i < A.length - 2; i++) {
        if (sorted[i] + sorted[i + 1] > sorted[i + 2]) {
            return 1;
        }
    }
    return 0;
}
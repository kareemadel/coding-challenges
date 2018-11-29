// Count minimal number of jumps from position X to Y.

function solution(X, Y, D) {
    // write your code in JavaScript (Node.js 8.9.4)
    return Math.ceil((Y - X) / D);
}
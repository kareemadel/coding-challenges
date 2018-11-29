// Find value that occurs in odd number of elements.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    return A.reduce((acc, curr) => acc ^ curr, 0);
}

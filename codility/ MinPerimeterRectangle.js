// Find the minimal perimeter of any rectangle whose area equals N.

function solution(N) {
    // write your code in JavaScript (Node.js 8.9.4)
    const root = Math.trunc(Math.sqrt(N));
    for (let i = Math.trunc(root); i > 0; i--) {
        if (N % i === 0) {
            return 2 * (i + N / i);
        }
    }
}
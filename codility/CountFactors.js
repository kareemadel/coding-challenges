// Count factors of given number n.

function solution(N) {
    // write your code in JavaScript (Node.js 8.9.4)
    const root = Math.sqrt(N);
    let counter = 0;
    for (let i = 1; i <= root; i++) {
        if ( N % i === 0) {
            if (N / i === i) counter++;
            else counter += 2;
        }
    }
    return counter;
}
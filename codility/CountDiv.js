// Compute number of integers divisible by k in range [a..b].

function solution(A, B, K) {
    // write your code in JavaScript (Node.js 8.9.4)
    const startRem = A % K;
    const start = startRem === 0? A : A + (K - startRem);
    const end = B - B % K;
    if (end < start) {
        return 0;
    }
    return (end - start) / K + 1;
    
}

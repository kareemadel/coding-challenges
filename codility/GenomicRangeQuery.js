// Find the minimal nucleotide from a range of sequence DNA.

function solution(S, P, Q) {
    // write your code in JavaScript (Node.js 8.9.4)
    const counters = [];
    let from, to;
    const result = [];
    const movingCounter = { A: 0, C: 0, G: 0, T: 0 };
    counters.push({...movingCounter});
    for (let char of S) {
        movingCounter[char]++;
        counters.push({...movingCounter});
    }
    for (let i = 0; i < P.length; i++) {
        from = counters[P[i]];
        to = counters[Q[i] + 1];
        if (from.A !== to.A) {
            result.push(1);
        } else if (from.C !== to.C) {
            result.push(2);
        } else if (from.G !== to.G) {
            result.push(3);
        } else if (from.T !== to.T) {
            result.push(4);
        }
    }
    return result;
}
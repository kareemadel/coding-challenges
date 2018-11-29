// Compute the number of intersections in a sequence of discs.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    const discStart = Array(...Array(A.length)).map(() => 0);
    const discEnd = Array(...Array(A.length)).map(() => 0);
    for (let i = 0, l = A.length; i < A.length; i++) {
        const start = i > A[i] ? i - A[i] : 0;
        const end = i + A[i] < l ? i + A[i] : l;
        discStart[start]++;
        discEnd[end]++;
    }
    let activeDiscs = 0;
    let result = 0;
    for (let i = 0; i < A.length; i++) {
        if (discStart[i] > 0) {
            result += activeDiscs * discStart[i];
            result += (discStart[i] - 1) * discStart[i] / 2;
            if (result > 10000000) return -1; 
            activeDiscs += discStart[i];
        }
        activeDiscs -= discEnd[i];
    }
    return result;
}

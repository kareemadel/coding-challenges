// Find an index of an array such that its value occurs at more than half of indices in the array.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    const counters = {};
    for (let i = 0; i < A.length; i++) {
        if (counters[A[i]] === undefined) {
            counters[A[i]] = {count: 1, firstIndex: i};
        } else {
            counters[A[i]].count++;
        }
    }
    for (let n in counters) {
        if (counters[n].count > A.length / 2) {
            return counters[n].firstIndex;
        }
    }
    return -1;
}
// Find the earliest time when a frog can jump to the other side of a river.

function solution(X, A) {
    // write your code in JavaScript (Node.js 8.9.4)
    const stepMap = {};
    let steps = X;
    for (let i =  0; i < A.length; i++) {
        if (A[i] <= X && stepMap[A[i]] === undefined){
            stepMap[A[i]] = i;
            steps--;
        }
        if (steps === 0) {
            return i
        }
    }
    return -1;
}
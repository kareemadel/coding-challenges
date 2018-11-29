// Count the number of passing cars on the road.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    let count = 0;
    let increment = 0;
    for (let i = 0; i < A.length; i++) {
        if (A[i] === 0) {
            increment++;
        } else {
            count += increment;
        }
        if (count > 1000000000) {
            return -1;
        }
    }
    return count;
}
// Check whether array A is a permutation.

function solution(A) {
    const len = A.length;
    const countArr = Array.apply(null, Array(len)).map(Number.prototype.valueOf, 0);
    for (let n of A) {
        if (n > len) {
            return 0;
        }
        countArr[n - 1]++;
    }
    for (let count of countArr) {
        if (count !== 1) {
            return 0;
        }
    }
    return 1;
}
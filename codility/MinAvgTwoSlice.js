// Find the minimal average of any slice containing at least two elements.

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    let minAvg = Number.POSITIVE_INFINITY;
    let minIndex, doubleAvg, tripleAvg, movingMin;
    for (let i = 0; i < A.length - 1; i++) {
        doubleAvg = (A[i] + A[i + 1]) / 2;
        tripleAvg = i < A.length - 2 ? (A[i] + A[i + 1] + A[i + 2]) / 3 : Number.POSITIVE_INFINITY;
        movingMin = Math.min(doubleAvg, tripleAvg);
        if (movingMin < minAvg) {
            minIndex = i;
            minAvg = movingMin;
        }
    }
    return minIndex;
}
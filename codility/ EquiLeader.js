// Find the index S such that the leaders of the sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N - 1] are the same.

function solution(A) {
    const counters = {};
    for (let n of A) {
        counters[n] = counters[n] ? counters[n] + 1 : 1;
    }
    let leader = null, total = null;
    for (let n in counters) {
        if (counters[n] > A.length / 2) {
            leader = +n;
            total = counters[n];
            break;
        }
    }
    if (leader === null) return 0;
    let rightCount = total, leftCount = 0, equiCount = 0;
    for (let i = 0; i < A.length; i++) {
        if (A[i] === leader) {
            rightCount--;
            leftCount++;
        }
        if (leftCount > (i + 1) / 2 && rightCount > (A.length - i - 1) / 2) {
            equiCount++;
        }
    }
    return equiCount;
}
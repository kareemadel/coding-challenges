// Divide an array into the maximum number of same-sized blocks, each of which should contain an index P such that A[P - 1] < A[P] > A[P + 1].

function solution(A) {
    // write your code in JavaScript (Node.js 8.9.4)
    const peaks = [];
    for (let i = 1; i < A.length - 1; i++) {
        if (A[i] > A[i - 1] && A[i] > A[i + 1]) {
            peaks.push(i);
        }
    }
    if (peaks.length === 0) return 0;
    let partitionSize, partitions, partitionCount;
    for (let i = peaks.length; i > 0; i--) {
        if (A.length % i === 0) {
            partitionSize = A.length / i;
            partitions = Array(...Array(i)).map(() => false);
            partitionCount = 0;
            for (let peak of peaks) {
                if (!partitions[Math.trunc(peak / partitionSize)]) {
                    partitions[Math.trunc(peak / partitionSize)] = true;
                    partitionCount++;
                }
            }
            if (partitionCount === i) {
                return i;
            }
        }
    }
    return 0;
}

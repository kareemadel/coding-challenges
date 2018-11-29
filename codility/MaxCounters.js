// Calculate the values of counters after applying all alternating operations: increase counter by 1; set value of all counters to current maximum.

function solution(N, A) {
    const counters = Array(...Array(N)).map(() => 0);
    let maxCounter = 0;
    let minCounter = 0;
    for (let op of A) {
        if (op <= N) {
            counters[op - 1] = counters[op - 1] > minCounter ? counters[op - 1] + 1 : minCounter + 1;
            maxCounter = Math.max(maxCounter, counters[op - 1]);
        } else {
            minCounter = maxCounter;
        }
    }
    
    counters.forEach((n, i) => {
        if (n < minCounter) {
            counters[i] = minCounter;
        }
    });
    return counters;
}
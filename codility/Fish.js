// N voracious fish are moving along a river. Calculate how many fish are alive.

function solution(A, B) {
    let counter = 0;
    const stack = [];
    for (let i = 0; i < A.length; i++) {
        if (B[i] === 1) {
            stack.push(A[i]);
        } else {
            while (stack.length !== 0) {
                if (stack[stack.length - 1] > A[i]) {
                    break;
                } else {
                    stack.pop();
                }
            }
            if (stack.length === 0) counter++;
        }
    }
    counter += stack.length;
    return counter
}
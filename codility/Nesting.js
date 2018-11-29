// Determine whether a given string of parentheses (single type) is properly nested.

function solution(S) {
    const stack = [];
    for (const b of S) {
        if (b === '(') {
            stack.push(b);
        } else {
            if ( '(' !== stack.pop()) return 0;
        }
    }
    if ( stack.length !== 0) return 0;
    return 1;
}
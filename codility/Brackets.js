// Determine whether a given string of parentheses (multiple types) is properly nested.

function solution(S) {
    const brackets = {'}': '{', ')': '(', ']': '['};
    const stack = [];
    for (const b of S) {
        if (b === '(' || b === '{' || b === '[') {
            stack.push(b);
        } else {
            if ( brackets[b] !== stack.pop()) return 0;
        }
    }
    if ( stack.length !== 0) return 0;
    return 1;
}
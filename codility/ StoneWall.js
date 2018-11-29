// Cover "Manhattan skyline" using the minimum number of rectangles.

function solution(H) {
    const stack = [H[0]];
    let count = 1, top;
    for (let i = 1; i < H.length; i++) {
        top = stack[stack.length - 1];
        while(stack.length !== 0 && H[i] < top) {
            stack.pop();
            top = stack[stack.length - 1];
        }
        if (H[i] === top) continue;
        else {
            stack.push(H[i]);
            count++;
        }
    }
    return count;
}
def factor(x):
    if x == 1:
        return [1]
    cur = 2
    arr = []
    while x > 1:
        if x % cur == 0:
            arr.append(cur)
            x = x / cur
        else:
            cur = cur + 1
    return arr

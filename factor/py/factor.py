def factor(x):
    cur = 2
    arr = []
    while x > 1:
        while x % cur == 0:
            arr.append(cur)
            x = x / cur
        cur = cur + 1
    return arr

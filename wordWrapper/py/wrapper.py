def wrapper(string, size):
    result = []
    while string != "":
        line = string[:size]
        ll = line.rfind(" ")
        if ll > -1:
            line = line[:ll]
        string = string[len(line):].strip()
        result.append(line.strip())

    return result

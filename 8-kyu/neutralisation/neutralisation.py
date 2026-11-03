def neutralise(s1: str, s2 : str) -> str:
    res: str = ""
    for i in range(0, len(s1)) :
        if s1[i] == '+' and s2[i] == '+' :
            res += "+"
        elif s1[i] == '-' and s2[i] == '-' :
            res += "-"
        elif s1[i] == '+' and s2[i] == '-' or s1[i] == '-' and s2[i] == '+':
            res += "0"
        
    return res
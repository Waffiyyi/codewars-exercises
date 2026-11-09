def elimination(arr: list[int]) -> int:
    freq: list[int] = []
    
    for n in arr:
        if n in freq:
            return n
        else:
            freq.append(n)
    return None
    74: "melon",
    75: "pineapple",
    76: "melon",
    77: "pineapple",
    78: "cucumber",
    79: "pineapple",
    80: "cucumber",
    81: "apple",
    82: "grape",
    83: "orange",
    84: "grape",
    85: "cherry",
    86: "grape",
    87: "cherry",
    88: "pear",
    89: "cherry",
    90: "apple",
    91: "kiwi",
    92: "banana",
    93: "kiwi",
    94: "banana",
    95: "melon",
    96: "banana",
    97: "melon",
    98: "pineapple",
    99: "apple",
    100: "pineapple"}
    
    newN:int = number
    
    digSum: int = 0
    strNum: str = str(number)
    
    for ch in strNum :
        digSum += int(ch)
    
    newN = number - digSum
    
    if newN in fruits:
        return fruits[newN]
    
    return subtract_sum(newN)
# import numpy as np

def unique_values(a,b):
    # Return the unique values of a and b
    return sorted(set(a+b))

a = ["a","b","c"]
b = ["b","c","d"]

print(unique_values(a,b))

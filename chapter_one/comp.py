from random import randint
import time

ELEMENTS = 1000000
RANDOM_RANGE = 1000000


x = [randint(1, RANDOM_RANGE) for i in range(ELEMENTS)]
now = time.monotonic()
x.sort()
print("Time: ", time.monotonic() - now, len(x))


# even in C implementation of sorting, go is faster than python in 3 times

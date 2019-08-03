#!/usr/bin/env python3

import asyncio
import time
import random

# ANSI colors
c = (
        "\033[0m",  # End of color
        "\033[36m",  # Cyan
        "\033[91m",  # Red
        "\033[35m",  # Magenta
    )


async def make_random(idx: int, threshold: int= 6) -> int:
    print(c[idx + 1] + f"Initiated make_random({idx}).")
    i = random.randint(0, 10)

    while i <= threshold:
        print(c[idx + 1] + f"make_random({idx}) == {i} tow low; retrying.")
        await asyncio.sleep(idx + 1)
        i = random.randint(0, 10)

    print(c[idx + 1] + f"---> Finished: make_random({idx}) == {i}" + c[0])
    return i

async def count():
    print("One")
    await asyncio.sleep(1)
    print("Two")


async def main():
    # await asyncio.gather(count(), count(), count())
    res = await asyncio.gather(*(make_random(i, 10-i-1) for i in range(3)))
    return res


if __name__ == "__main__":
    # s = time.perf_counter()
    random.seed(444)

    # python3.6
    #r1, r2, r3 = asyncio.get_event_loop().run_until_complete(main())

    # if python3.7
    r1, r2, r3 = asyncio.run(main())

    # elapsed = time.perf_counter() - s
    # print(f"{__file__} executed in {elapsed: 0.2f} seconds.")
    print()
    print(f"r1: {r1}, r2: {r2}, r3: {r3}")

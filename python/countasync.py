#!/usr/bin/env python3

import asyncio
import time

async def count():
    print("One")
    await asyncio.sleep(1)
    print("Two")


async def main():
    await asyncio.gather(count(), count(), count())


if __name__ == "__main__":
    s = time.perf_counter()
    asyncio.get_event_loop().run_until_complete(main())
    elapsed = time.perf_counter() - s
    print(f"{__file__} executed in {elapsed: 0.2f} seconds.")

import asyncio
import threading
import time
import httpx
import requests


async def fetchall(session: httpx.AsyncClient, now: float):
    resp = await session.get('https://ifconfig.me')
    print(resp.status_code, time.monotonic() - now)


async def main(now: float):
    async with httpx.AsyncClient() as session:
        tasks = [fetchall(session, now) for _ in range(20)]
        await asyncio.gather(*tasks)


now = time.monotonic()
asyncio.run(main(now))
print("asyncio", time.monotonic() - now)


def sync_fetchall(now: float):
    resp = requests.get('https://ifconfig.me')
    print(resp.status_code, time.monotonic() - now)


def thread_fetchall(now: float):
    threads = [threading.Thread(target=sync_fetchall, args=(now, )) for i in range(20)]
    for thread in threads:
        thread.start()
    for thread in threads:
        thread.join()


now = time.monotonic()
thread_fetchall(now)
print("threaded", time.monotonic() - now)


now = time.monotonic()
for _ in range(20):
    sync_fetchall(now)
print("sync", time.monotonic() - now)



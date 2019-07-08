import concurrent.futures
import requests
import threading
import time
import aiohttp
import asyncio

sites = [
    "https://www.bilibili.com",
    "https://www.douban.com",
] * 80

class Synchronous:
    def __init__(self, sites):
        self.sites = sites

    @staticmethod
    def download_site(url, session):
        with session.get(url) as response:
            print(f"Read {len(response.content)} from {url}")


    def download_all_sites(self, sites):
        with requests.Session() as session:
            for url in sites:
                self.download_site(url, session)


    def run(self):
        start_time = time.time()
        self.download_all_sites(self.sites)
        duration = time.time() - start_time
        print(f"Downloaded {len(sites)} in {duration} seconds")


class MultiThreading:
    def __init__(self, sites):
        self.sites = sites
        self.thread_local = threading.local()

    def get_session(self):
        if not hasattr(self.thread_local, "session"):
            self.thread_local.session = requests.Session()
        return self.thread_local.session

    def download_site(self, url):
        session = self.get_session()
        with session.get(url) as response:
            print(f'Read {len(response.content)} from {url}')

    def download_all_sites(self):
        with concurrent.futures.ThreadPoolExecutor(max_workers=7) as executor:
            executor.map(self.download_site, self.sites)

    def run(self):
        start_time = time.time()
        self.download_all_sites()
        duration = time.time() - start_time
        print(f'Downloaded {len(sites)} in {duration} seconds')


class Asynchronous:
    def __init__(self, sites):
        self.sites = sites

    @staticmethod
    async def download_site(session, url):
        async with session.get(url) as response:
            print('Read {0} from {1}'.format(len(await response.text()), url))

    async def download_all_sites(self, sites):
        async with aiohttp.ClientSession() as session:
            tasks = []
            for url in sites:
                task = asyncio.ensure_future(self.download_site(session, url))
                tasks.append(task)

            await asyncio.gather(*tasks, return_exceptions=True)
            
    def async_run(self):
        start_time = time.time()
        asyncio.get_event_loop().run_until_complete(self.download_all_sites(self.sites))
        duration = time.time() - start_time
        print(f'Downloaded {len(sites)} sites in {duration} seconds')


def main():
    print('Synchronous')
    syc = Synchronous(sites)
    syc.run()

    print('MultiThreading')
    mt = MultiThreading(sites)
    mt.run()

    print('Asynchronous')
    asyc = Asynchronous(sites)
    asyc.async_run()
    
if __name__ == "__main__":
    main()


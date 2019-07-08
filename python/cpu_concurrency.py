import time
import multiprocessing

class Sync:
    def __init__(self):
        self.nu = [5_000_000 + x for x in range(20)]

    @staticmethod
    def cpu_bound(number):
        return sum(i*i for i in range(number))

    def find_sums(self, numbers):
        for number in numbers:
            self.cpu_bound(number)

    def run(self):
        start_time = time.time()
        self.find_sums(self.nu)
        duration = time.time() - start_time
        print(f"Duration {duration} seconds")


class MultiProcessing:
    def __init__(self):
        self.nu = [5_000_000 + x for x in range(20)]

    @staticmethod
    def cpu_bound(number):
        return sum(i * i  for i in range(number))

    def find_sums(self, numbers):
        with multiprocessing.Pool() as pool:
            pool.map(self.cpu_bound, numbers)

    def run(self):
        start_time = time.time()
        self.find_sums(self.nu)
        duration = time.time() - start_time
        print(f"Duration {duration} seconds")

def main():
    print('Sync 1-thread 1-process:')
    sy = Sync()
    sy.run()
     
    print('Multi Process:')
    mp = MultiProcessing()
    mp.run()


if __name__ == "__main__":
    main()


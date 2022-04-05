## Script to check for all prime numbers in a given range.
from datetime import datetime
from time import time


def check_prime(num:int) -> bool:
    flag_var:bool = True
    upper_limit = num//2
    i = 2
    while i <= upper_limit:
        if num%i == 0:
            flag_var = False
        i += 1
    
    return flag_var

def claculate_time(total_seconds:int) -> str:
    total_minutes = total_seconds//60
    seconds = total_seconds - (total_minutes*60)
    hours = total_minutes//60
    minutes = total_minutes - (hours*60)

    return f"\nThis process took {hours} hour(s), {minutes} minute(s) and {seconds} second(s)."

def main():
    limit = int(input("Enter the upper limit of the range: "))

    i: int = 3
    list_of_primes = []

    t1 = time()
    while i<=limit:
        print(f"\nChecking: {i}")
        if check_prime(i):
            list_of_primes.append(i)
        i += 1
    t2 = time()

    print(f"\nThe list of Prime Numbers until {limit} is: \n{list_of_primes}")
    print(claculate_time(t2-t1))

if __name__ == "__main__":
    main()
    
    


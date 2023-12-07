import time

start_time = time.time()  # Capture the start time
for i in range(1000000000):
    print(i)
end_time = time.time()  # Capture the end time

duration = end_time - start_time  # Calculate the duration
print(f"Time taken: {duration} seconds")

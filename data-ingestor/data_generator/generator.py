import json, random
from helpers.get_time import get_current_time 

# Lets setup the data array which we will be writing to our JSON file
def generate():
    data = []
    currentTimestamp = get_current_time()

    # For the previous 10 minutes create some a data point at every second
    currentPreviousTime = 0
    while currentPreviousTime < (60 * 10):
        dataTimeStamp = currentTimestamp - currentPreviousTime
        data.append({
            'timestamp': dataTimeStamp,
            'cpu_load': random.uniform(0, 100),
            'concurrency': random.uniform(0, 500000)
        })
        currentPreviousTime += 1

    with open('data.json', 'w') as dataFile:
        json.dump(data, dataFile)
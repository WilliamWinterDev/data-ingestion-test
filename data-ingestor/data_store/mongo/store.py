import pymongo, json
from helpers.get_time import get_current_time 

def store():
    print('Starting to store data in MongoDB')
    
    mongo_client = pymongo.MongoClient("mongodb://localhost:6002/")
    database = mongo_client["metrics"]
    collection = database["cpu_data"]
    collection.create_index(
        [("timestamp", pymongo.DESCENDING)],
        unique=True
    )

    with open('data.json') as metrics_file:
        metrics_data = json.load(metrics_file)

    current_time = get_current_time()
    seconds_to_ingest = 60 * 5
    clean_data = []

    for i in range(len(metrics_data)):
        if metrics_data[i]["timestamp"] >= (current_time - seconds_to_ingest):
            clean_data.append(metrics_data[i])

    print("Inserting " + str(len(clean_data)) + " items into the store")
    
    try:
        result = collection.insert_many(clean_data)
    except pymongo.errors.BulkWriteError as errors:
        # We can ignore error code 11000 as this just means we are trying to insert a duplicate key since the timestamp should be unique, 
        # this would mainly be an issue if we have mock data
        unexpected_errors = list(filter(lambda x: x['code'] != 11000, errors.details['writeErrors']))
        if(len(unexpected_errors) > 0):
            print("Something has gone wrong when trying to insert the data: {unexpected_errors}")
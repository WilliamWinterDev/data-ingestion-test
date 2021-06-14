# Data Ingestion

A simple set of scripts using Python and Golang to ingest a series of data via a Python script and then output it using a Golang API

## Requirements

**Docker** - Docker is used for our MongoDB container and the Golang API.

**Python 3.9** - Used for the data generator and data ingestor.

## Data-Ingestor

The data-ingestor is written in Python 3.9 and uses PyMongo to store a JSON file in a Mongo database

### Usage

`docker compose up -d` inside the root folder so we have the Docker MongoDB instance

`pip install -r requirements.txt` - Install the requirements for the data ingestor.

`python ingestor.py` this will generate a file called __data.json__ which will include 10 minutes of mock metrics data. The script will then look at __data.json__ and then loop through the data and make sure the data is within the past 5 minutes and if it is then import the data into the MongoDB container.

## Data-API

The data api is written in Golang and uses Mux to create a basic API to read the metric data out of the docker MongoDB container

### Usage

`docker compose build && docker compose up -d` inside the root folder we should now have the read only API and the MongoDB container running, to access the API you can use [https://localhost:6001/](http://localhost:6001/).

__startTimestamp__ & __endTimestamp__ are both epoch timestamps for the data segment you wish to view.

### Routes
[http://localhost:6001/api/metrics](http://localhost:6001/api/metrics) - This route will return all metrics

[http://localhost:6001/api/metrics/{startTimestamp}/{endTimestamp}](http://localhost:6001/api/metrics/{startTimestamp}/{endTimestamp}) - This route will let you filter the data between 2 timestamps

_Note:_ If you are using Golang directly to start the API the port will be __6003__.

## Todo with more time

Given I had more time I would implement the following:

- Combine the Python and Golang applications to use single environment variables file so they are not duplication and easier to manage.
- I would add unit tests to both projects.
- I would rework filters so they're easier to create instead of being hard coded to timestamp.
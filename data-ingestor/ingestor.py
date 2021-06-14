from data_generator.generator import generate
import data_store.mongo.store

# Lets use our mock data generator to generate a json file 
generate()

data_store.mongo.store.store()
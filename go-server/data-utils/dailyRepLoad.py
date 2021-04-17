import pandas as pd
import numpy as np
import requests
import json
import csv
import yaml
import boto3
import os
from io import StringIO

s3 = boto3.resource(
    's3',
    aws_access_key_id = os.getenv('AWS_ACCESS_KEY_ID'),
    aws_secret_access_key = os.getenv('AWS_SECRET_ACCESS_KEY')
)  

def saveData(URL, headers, outputFile):

    outputFilePath = "representatives/" + outputFile
    bucket = 'gov-connect'

    response = requests.get(url = URL, headers = headers)
    df = json.loads(response.text)
    outputDF = pd.DataFrame(df['results'][0]['members'])

    csv_buffer = StringIO()
    outputDF.to_csv(csv_buffer)
    s3.Object(bucket, outputFilePath).put(Body=csv_buffer.getvalue())


#load in configurations
with open("../keys/keys.yml") as file:
    config = yaml.load(file, Loader=yaml.FullLoader)

# save house members list
URL = 'https://api.propublica.org/congress/v1/116/house/members.json'
headers = {'x-api-key': config['keys.propublica']}
saveData(URL=URL, headers=headers, outputFile='house_members.csv')

# save senate members list
URL = 'https://api.propublica.org/congress/v1/116/senate/members.json'
saveData(URL=URL, headers=headers, outputFile='senate_members.csv')
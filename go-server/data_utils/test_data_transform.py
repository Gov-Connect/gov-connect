import numpy as np
import pandas as pd
import os

## Create file load in function
def AddPhotoURL(fileName):
    ## load in test csv files as dataframes
    # fileName = "top_house_reps.csv"
    fileContent = pd.read_csv(TargetFolder + "/" + fileName)

    ## add photoURL based on https://theunitedstates.io/images/congress/450x550/O000172.jpg
    for i, row in fileContent.iterrows():
        photo_url = "https://theunitedstates.io/images/congress/450x550/" + row.id + ".jpg"
        fileContent.loc[i,"photo_url"] = photo_url

    ## save out file 
    filePath = "../test_data/" + fileName
    fileContent.to_csv(filePath, index=False)
    print("Adding Photo URL to " + file)

## pull all files in the 'test_data' dir and execute function
TargetFolder = "../test_data"
for file in os.listdir(TargetFolder):
    if file != "user_favorite_reps.csv":
        AddPhotoURL(file)
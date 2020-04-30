import logging
import boto3
import urllib.request as urllib
import os
import time
from botocore.exceptions import ClientError

pathToNewLogs = './new/'
pathToOldLogs = './old/'


def internet_on():
    try:
        urllib.urlopen('http:/172.217.18.238', timeout=1)
        return True
    except urllib.URLError as err:
        logging.error(err) 
        return False

def upload_file(file_name, bucket, object_name=None):
    """Upload a file to an S3 bucket

    :param file_name: File to upload
    :param bucket: Bucket to upload to
    :param object_name: S3 object name. If not specified then file_name is used
    :return: True if file was uploaded, else False
    """

    # If S3 object_name was not specified, use file_name
    if object_name is None:
        object_name = file_name

    # Upload the file
    s3_client = boto3.client('s3')
    try:
        response = s3_client.upload_file(file_name, bucket, object_name)
        #print(response)
    except ClientError as e:
        logging.error(e)
        print(e)
        return False
    return True

def main():
    while True:
        newLogsDirectory = os.fsencode(pathToNewLogs)
        oldLogsDirectory = os.fsencode(pathToOldLogs)

        if(internet_on):
            for file in os.listdir(newLogsDirectory):
                filename = os.fsdecode(file)
                print('Uploading ' + filename)
                upload_file(pathToNewLogs+ filename, 'rowa', 'logs/{}'.format(filename))
                os.rename(pathToNewLogs + filename, pathToOldLogs + filename)
            time.sleep(3)
if __name__ == "__main__":
    main()

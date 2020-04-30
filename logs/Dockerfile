FROM python:3

ADD uploadLogs.py /

RUN pip install boto3

CMD [ "python", "./uploadLogs.py" ]
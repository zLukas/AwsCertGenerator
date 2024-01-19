import boto3
import os


client = boto3.client('dynamodb')
TABLE_NAME=os.getenv("TABLE_NAME")


def get_database_record(cert_name: str):
    return client.get_item(TableName=TABLE_NAME,
        Key={"Name": {"S": cert_name}})


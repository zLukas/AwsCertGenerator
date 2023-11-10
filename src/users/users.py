import boto3 

client = boto3.client('iam')
IAM_EXCEPTIONS =(client.exceptions.LimitExceededException,
                client.exceptions.NoSuchEntityException,
                client.exceptions.DeleteConflictException,
                client.exceptions.ConcurrentModificationException,
                client.exceptions.ServiceFailureException,
                client.exceptions.EntityAlreadyExistsException,
                client.exceptions.InvalidInputException,)



def list_users():
    response = client.list_users()
    users = map(lambda u: (u["UserName"], u["Path"]), response["Users"])
    return list(users)

def delete_users(users: list):
    return_val = None
    response = None
    for name in users:
        try: 
            keys_meta = client.list_access_keys(UserName = name)
            keys_ids = list(map(lambda k: k["AccessKeyId"], keys_ids["AccessKeyMetadata"]))
            map(lambda id : client.delete_access_keys(UserName=name, AccessKeyId=id), keys_ids)
            response = client.delete_user(UserName=name)

        except IAM_EXCEPTIONS as e:
            reponse = {"Error": e}
        finally:
            return_val[name] = response
    return return_val

def create_users(users: list):
    return_val = None 
    key = None
    for name in users:
        try: 
            response = client.create_user(
                Path='/certClient/',
                UserName=name,
                Tags=[
                {
                    'Key': 'Creator',
                    'Value': 'Lambda'
                },
                ]
            )
            key = client.create_access_key(UserName=name)
        except IAM_EXCEPTIONS as e:
            reponse = {"Error": e}
            key = None
        finally:
            return_val[name] = {"user": response,
                                "key": key}
    return return_val

RUN = {
    "create": create_users,
    "delete": delete_users,
    "list": list_users
}

def handler(event, context): 
    option = event["option"]
    RUN[option](event["args"])

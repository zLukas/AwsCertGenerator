import boto3 
client = boto3.client('iam')
IAM_EXCEPTIONS =(client.exceptions.LimitExceededException,
                client.exceptions.NoSuchEntityException,
                client.exceptions.DeleteConflictException,
                client.exceptions.ConcurrentModificationException,
                client.exceptions.ServiceFailureException,
                client.exceptions.EntityAlreadyExistsException,
                client.exceptions.InvalidInputException,)



def list_users(*args):
    response = client.list_users() 
    users = list(map(lambda u: (u["UserName"], u["Path"]), response["Users"]))
    users = list(filter(lambda u: "cert" in  u[1], users))
    return users

def delete_users(users: list):
    return_val = {}
    response = {}
    for name in users:
        try: 
            keys_meta = client.list_access_keys(UserName = name)
            keys_ids = list(map(lambda k: k["AccessKeyId"], keys_meta["AccessKeyMetadata"]))
            for key in keys_ids:
                client.delete_access_key(
                    UserName=name, 
                    AccessKeyId=key)
            response = client.delete_user(UserName=name)

        except IAM_EXCEPTIONS as e:
            reponse = {"Error": e}
        finally:
            return_val[name] = response
    return return_val

def create_users(users: list):
    return_val = {} 
    key = {}
    response = {}
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
            response= {"Error": e}

        finally:
            return_val[name] = {"user": response,
                                "key": key}
    return return_val

RUN = {
    "create": create_users,
    "delete": delete_users,
    "list": list_users
}

def lambda_handler(event, context): 
    option = event["option"]
    print(RUN[option](event["args"]))
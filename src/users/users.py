import boto3 

client = boto3.client('iam')

def list_users():
    response = client.list_users()
    users = map(lambda u: (u["UserName"], u["Path"]), response["Users"])
    return list(users)

def delete_users(users: list):
    delete_exceptions = (client.exceptions.LimitExceededException,
                client.exceptions.NoSuchEntityException,
                client.exceptions.DeleteConflictException,
                client.exceptions.ConcurrentModificationException,
                client.exceptions.ServiceFailureException,)
    return_val = None
    for name in user:
        try: 
            response = client.delete_user(UserName=user_name)

            except delete_exceptions as e:
                reponse = {"Error": e}
            finally:
                return_val[name] = response
    return return_val

def create_users(users: list):
    create_exceptions = (client.exceptions.LimitExceededException,
        client.exceptions.EntityAlreadyExistsException,
        client.exceptions.NoSuchEntityException,
        client.exceptions.InvalidInputException,
        client.exceptions.ConcurrentModificationException,
        client.exceptions.ServiceFailureException,)
    return_val = None 
    for name in users:
        try: 
            response = client.create_user(
                Path='/certClient/',
                UserName=user_name,
                Tags=[
                {
                    'Key': 'Creator',
                    'Value': 'Lambda'
                },
                ]
            )
        except create_exceptions as e:
            reponse = {"Error": e}
        finally:
            return_val[name] = response
    return return_val

RUN = {
    "create": create_users,
    "delete": delete_users,
    "list": list_users
}

def handler_name(event, context): 
    option = event["option"]
    RUN[option](event["args"])

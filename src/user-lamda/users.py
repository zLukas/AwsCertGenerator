import boto3 
import os

client = boto3.client('cognito-idp')
COGNITO_POOL_ID=os.getenv("COGNITO_POOL_ID")

def request(func):
    def try_catch(*args):
        try:
            response = func(*args)
        except Exception as e:
            response = str(e)
        finally:
            return response
    return try_catch

@request
def create_user(*users: list):
    for user in users:
        response = client.admin_create_user(
                UserPoolId=COGNITO_POOL_ID,
                Username=user["Username"],
                TemporaryPassword='Password0$',
                #MessageAction='RESEND',
                DesiredDeliveryMediums=['EMAIL'],
                ClientMetadata={
                    'creator': 'Admin Lambda',
                },
                UserAttributes=[
                {
                    'Name': 'email',
                    'Value': user["email"]
                },])
        response = client.admin_add_user_to_group(
                    UserPoolId=COGNITO_POOL_ID,
                    Username=user["Username"],
                    GroupName="default-group")
    print(f"{user} created succefully")
    return response

@request
def list_users(*args):
    EMAIL_IDX = 1
    response = client.list_users(
            UserPoolId=COGNITO_POOL_ID)
    return [(i["Username"], i["Attributes"][EMAIL_IDX]["Value"])
                                    for i in response["Users"]]

@request
def delete_users(*users: str):
    response = None
    for user in users:
        response = client.admin_delete_user(
            UserPoolId=COGNITO_POOL_ID,
            Username=user)
    return response


RUN = {
    "create": create_user,
    "delete": delete_users,
    "list": list_users
}

def lambda_handler(event, context): 
    option = event["option"]
    print(RUN[option](event["args"]))

if __name__ == "__main__":
    create_user({
            "Username": "test_user",
            "email": "not@yourbussiness.com"
        })
    print(list_users())
    delete_users("test_user")
    print(list_users())
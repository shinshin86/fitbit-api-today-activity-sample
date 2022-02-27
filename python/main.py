from datetime import datetime
import requests


def main():
    today = datetime.now().strftime('%Y-%m-%d')
    url = "https://api.fitbit.com/1/user/-/activities/date/" + today + ".json"

    # Your token
    token  = ""

    headers = {
        "Authorization": "Bearer {}".format(token)
    }

    r = requests.get(url, headers=headers)
    print(r.json())


if __name__ == "__main__":
    main()

from selenium import webdriver
import time
import os
import dotenv
import requests

dotenv.load_dotenv()

website = os.getenv('WEBSITE', 'http://localhost:5173')
if website[-1] == '/':
    website = website[:-1]
    
path = os.getenv('WEBSITE_PATH', '/readreport?id=0')

# Set up Chrome driver
opt = webdriver.ChromeOptions()
opt.add_argument('--headless')
opt.add_argument('--no-sandbox')
driver = webdriver.Chrome(options=opt)

# Set the admin cookie
driver.get(website)
domain = website.split('//')[1]
driver.add_cookie({'name': 'admin', 'value': 'Us3_m3_d4ddy', 'path': '/'})

while True:
    # Go to the report page
    driver.get(website+path)

    # Delay for 2 seconds
    print('Waiting for 2 seconds...')
    time.sleep(2)
    
    print("Checking if report exists...")
    exists = False
    url = driver.current_url
    if url == website+path:
        if 'Report #' in driver.page_source:
            exists = True
            print('Deleting report...')
            # Delete the report
            delete_button = driver.find_element(value='qsHJGDJQHSGDJHQGSDJFQSGSD')
            delete_button.click()
    else:
        exists = True
        # Probably got redirected by XSS
        # Delete the report via API
        requests.delete(website+'/api/report/0', cookies={'admin': 'Us3_m3_d4ddy'})
        
    print('Report exists =', exists)

    if not exists:
        print("Waiting for 30 seconds...")
        time.sleep(30)

    print('Reloading cookie...')
    driver.add_cookie({'name': 'admin', 'value': 'Us3_m3_d4ddy', 'path': '/'})
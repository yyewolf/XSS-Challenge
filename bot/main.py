from selenium import webdriver
from selenium.webdriver.common.keys import Keys
import time
import os
import dotenv

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
driver.add_cookie({'name': 'admin', 'value': 'ckjsdmldMLKJSQ456', 'path': '/'})

while True:
    # Go to the report page
    driver.get(website+path)
    
    print('Waiting for report to be generated...')
    
    exists = False

    # Wait for the report or not found message to appear
    while True:
        if 'Report #' in driver.page_source or 'Report not found' in driver.page_source:
            if 'Report #' in driver.page_source:
                exists = True
            break
        time.sleep(0.1)
        
    print('Report exists =', exists)

    # Delay for 2 seconds
    print('Waiting for 2 seconds...')
    time.sleep(2)

    if exists:
        print('Deleting report...')
        # Delete the report
        delete_button = driver.find_element(value='qsHJGDJQHSGDJHQGSDJFQSGSD')
        delete_button.click()
    else:
        print("Waiting for 30 seconds...")
        time.sleep(30)

    print('Reloading cookie...')
    driver.add_cookie({'name': 'admin', 'value': 'ckjsdmldMLKJSQ456', 'path': '/'})
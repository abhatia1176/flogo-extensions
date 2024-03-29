# Custom Functions
Set of Custom Functions for Flogo Enterprise and Flogo on TCI.

## How to install Custom Functions: 
The steps below uses array function group as an example, but steps remain the same for any category.

1. To install array functions, there are two options:

   - Use the following github url - github.com/abhatia1176/flogo-extensions/function/array
   - Clone Repo and zip up the directory "array" under "flogo-extensions/function" and upload it to the extensions in Flogo.

   We will use github url for uploading the extension. For other categories, install URL are listed below:
   - Flogo: github.com/abhatia1176/flogo-extensions/function/flogo
   - URL: github.com/abhatia1176/flogo-extensions/function/url
   - DateTime: github.com/abhatia1176/flogo-extensions/function/datetime
   - JSON - github.com/abhatia1176/flogo-extensions/function/json
   - Array - github.com/abhatia1176/flogo-extensions/function/array
   
2. Go to "Environment & Tools" tab on TCI or TIBCO Flogo Enterprise, select "Extensions" under "Connector Management & Extensions".  On the next page, clikc on Upload. The screenshot below shows the use of github url. Once you click import, the function(s) will be imported as new extension. Please note that this will add a new category called custom_array under extensions. This can be changed by updating the descriptor.json file for that category.
![image](https://user-images.githubusercontent.com/4227956/73557110-42dd7380-4416-11ea-98d8-7d7747b90717.png)

3. Once the import is successful, click Done, as shown in the screenshot below.
![image](https://user-images.githubusercontent.com/4227956/73557493-e595f200-4416-11ea-8ff4-da9bc7cb1bc0.png)

4. A new category called "Custom_array" will be available, with all functions from that category listed on the right hand side, as shown below.
![image](https://user-images.githubusercontent.com/4227956/73557873-a916c600-4417-11ea-9c06-c27e306c0dfe.png)

Follow same approach as above for other category of functions. At present, there is only one category.

## How to Use Custom Functions in Flows:

1. Below screenshot shows an example of using custom array function `sum` in a mapper to get sum of all elements in a number array. The mapper is defined with the following schema:
     `{"sum": 123}`
     
     ![image](https://user-images.githubusercontent.com/4227956/73559003-e0867200-4419-11ea-809e-0bdf56f6c0a9.png)

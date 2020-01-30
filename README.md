# BainyQuoteApi
---
Get random quote from [BrainyQuote](https://www.brainyquote.com/) using this rest api...

### Base URL = `https://brainyquote.greenh.dev/`
---
## GET `/`
get random quote from random topic

## GET `/topic/{topic_name}`
get random quote from a topic

Parameters: 
`topic_name` = brainyquote topic name
- read `topic_list.txt` to get current topic avaiable in this api..

## Get `/findbyid?id={id}`
get quote by id

Parameters: 
`id` = quote id 

Quote ID == `{topic_name}{section}_{index}`
* example = `age1_12`
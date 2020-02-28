# BainyQuoteApi
---
Get random quote from [BrainyQuote](https://www.brainyquote.com/) using this rest api...

### Base URL = `https://brainyquote.greenh.dev/`
---
## GET `/`
get random quote from random topic

Query Parameter:
  `total`(optional) = return total number of quotes. max = 10, default = 0



## GET `/topic/{topic_name}`
get random quote from a topic

Parameter: 
`topic_name` = brainyquote topic name
- read `topic_list.txt` to get current topic avaiable in this api..

## Get `/findbyid?id={id}`
get quote by id

Query Parameter: 
`id` = quote id 

Quote ID == `{topic_name}{section}_{index}`
* example = `age1_12`

## Get `/daily`
get quote of the day
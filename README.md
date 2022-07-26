# URL shortener

Service is an API which allows user to create short links. <br/>

There are two methods <br/>
1) POST - creates a new record in service for current long link
2) GET - returns  long link for current short

Example of usage:<br/>
## POST
Input the following json in the POST request body (http://localhost:4000/create)
```json
{
    "longLink":"https://yandex.ru/"
}
```

Service will generate a short token and return next json:
```json
{"shortLink":"a4VqvURlab","longLink":"https://yandex.ru/"}
```
"shortLink" is the key to get our long link "https://yandex.ru/"
<br/>

## GET
Now we are going to get out long link.
Input the following json in the GET request body (http://localhost:4000/getLongLink)
```json
{
    "shortLink":"a4VqvURlab"
}
```

The answer will be:
```json
{"shortLink":"a4VqvURlab","longLink":"https://yandex.ru/"}
```

## Non-existent links
If you try to get non-existent link then service will just return `there is no long link for current short`
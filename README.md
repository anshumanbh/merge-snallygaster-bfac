# merge-snallygaster-bfac

Merging and Sorting the URLs/backupfiles obtained from [Snallygaster](https://github.com/hannob/snallygaster) and [BFAC](https://github.com/mazen160/bfac).

## Sample Snallygaster File:
```
[{"cause": "dotenv", "url": "http://991049ef.ngrok.io/.env", "misc": ""}, {"cause": "git_dir", "url": "http://a1191310.ngrok.io/.git/config", "misc": ""}, {"cause": "dotenv", "url": "http://a1191310.ngrok.io/.env", "misc": ""},{"cause": "dotenv", "url": "http://991049ef.ngrok.io/.env", "misc": ""}]
```

## Sample BFAC File:
```
[{"url": "http://a1191310.ngrok.io/.", "status_code": 200, "content_length": 11266}, {"url": "http://a1191310.ngrok.io/.git/config", "status_code": 200, "content_length": 7}, {"url": "https://991049ef.ngrok.io/.", "status_code": 200, "content_length": 11266}, {"url": "http://991049ef.ngrok.io/.", "status_code": 200, "content_length": 11266}]
```

## Merged & Sorted Results File:
```
http://991049ef.ngrok.io/.
http://991049ef.ngrok.io/.env
http://a1191310.ngrok.io/.
http://a1191310.ngrok.io/.env
http://a1191310.ngrok.io/.git/config
https://991049ef.ngrok.io/.
```

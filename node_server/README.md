# 再現手順
## 環境
- Ubuntu20.04
- node ver. v16.14.0
## 手順
1. 以下のコマンドを実行
```
mkdir node_server
cd node_server
npm install website-scraper
```

2. `node_server`ディレクトリ直下に`full_download.js`を作成して，以下のコードを記述．
```
import scrape from 'website-scraper'; // only as ESM, no CommonJS
const options = {
  urls: ['https://rakuten.co.jp/'],
  directory: './rakuten'
};

// with async/await
const result = await scrape(options);

// with promise
scrape(options).then((result) => {});
```

3. `package.json`に以下のコードを追加．
```
"type": "module"
```

4. 以下のコマンドを実行
```
node full_download.js
```

5. 作成された`rakuten`ディレクトリ直下に`server.js`を作成し，以下のコードを記述．
```
import { createServer } from 'http';
import { readFile } from 'fs';

var server = createServer(
    (request,response)=>{
        readFile('./index.html','UTF-8',(error,data)=>{
            response.writeHead(200, {'Content-Type':'text/html'});
            response.write(data);
            response.end();
        })
       
    }
);
server.listen(3000);
```
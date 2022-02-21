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
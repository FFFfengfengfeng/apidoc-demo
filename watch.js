const fs = require('fs');

fs.watch(`${__dirname}/dest/api_data.json`, (event, filename) => {
    console.log(`事件类型是: ${event}`);
    fs.readFile(`${__dirname}/dest/api_data.json`, 'utf-8', (err, data) => {
        if (err) {
            console.log(err);
        } else {
            let result = JSON.parse(data);
            console.log(result[0].group);
        }
    });
});
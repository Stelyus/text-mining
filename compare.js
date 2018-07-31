fs = require('fs')

// First I want to read the file
fs.readFile('./ref', function read(err, data) {
    if (err) {
        throw err;
    }

    // Invoke the next step here however you like

    var content = JSON.parse(data)

    fs.readFile('./me', function read(err, data1) {
         if (err) throw err;

         var me_content = JSON.parse(data1)

        if (content.length != me_content.length) {
            console.log(`Not same length: ref: ${content.length}, me: ${me_content.length}`);
            return;
        }
        
         for (let i = 0; i < me_content.length; i++) {
            if (me_content[i].word != content[i].word || me_content[i].freq != content[i].freq || me_content[i].distance != content[i].distance)
            {
                console.log("Not the same on " + i + " th element")
                return;
            }
         }
    })
});

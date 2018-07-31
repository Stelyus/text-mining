var fs = require('fs')
var exec = require("child_process").exec;


function execute(command, done) {
    path_binary_ref = process.argv[2]
    path_dct_ref = process.argv[3]


    path_binary_me = process.argv[4]
    path_dct_me = process.argv[5]

    exec(`echo "${command}" | ${path_binary_ref} ${path_dct_ref}`, function (err, stdout, stderr) {
        ref = JSON.parse(stdout)

        exec(`echo "${command}" | ${path_binary_me} ${path_dct_me}`, function (err, stdout1, stderr) {
            me = JSON.parse(stdout1)
             if (ref.length != me.length) {
                console.log('\x1b[31m%s\x1b[0m', "FAILED", command)
                console.log(`Not same length: ref: ${ref.length}, me: ${me.length}`);
            }
            
         for (let i = 0; i < me.length; i++) {
            if (me[i].word != ref[i].word || me[i].freq != ref[i].freq || me[i].distance != ref[i].distance)
            {
                if (ref.length == me.length) {
                    console.log('\x1b[31m%s\x1b[0m', "FAILED", command)
                }

                console.log("Not the same on " + i + " th element")
                console.log(ref[i])
                console.log(me[i])

                if (ref.length != me.length) {
                    w_ref = ref.map(w => w.word)
                    w_me = me.map(w=>w.word)

                    console.log("Print the 5 first elements")
                    w_ref.filter(w => w_me.indexOf(w) == -1).forEach((w, i) => {
                        if (i < 5) console.log('\x1b[33m%s\x1b[0m', "MISSING", w)
                    })
                }

                console.log()

                return done()
            }
         }

         console.log('\x1b[32m%s\x1b[0m', "PASSED", command)
         done()
        })
    })
}

function testing() {
    execute("approx 0 test", () => {
        execute("approx 1 test", () => {
            execute("approx 2 test", () => {
                execute("approx 2 mylovekis", () => {
                    execute("approx 4 mxrtnw4vrto", () => {
                        execute("approx 2 mxcomp", () => {
                            execute("approx 1 myzo", () => {
                                execute("approx 1 nabilo", () => {
                                    execute("approx 4 nelidetours", () => {
                                        execute("approx 2 nabilo", () => {})
                                    })
                                })
                            })
                        })
                    })
                })
            })
        })
    })
    
    
   
}

testing()

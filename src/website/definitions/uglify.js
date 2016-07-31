var uglify = require('uglify-js');

F.onCompileScript = function (filename, content) {
    if (filename === '') {
        //don't uglify inline file
        return content;
    }
    return uglify.minify(content, {fromString: true, mangle: false}).code;
};
var ObjectId = require('mongodb').ObjectID;

exports.install = function () {
    F.route('/', home);
};

function home() {
    this.view('/home/index');
}
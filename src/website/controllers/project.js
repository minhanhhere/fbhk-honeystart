var ObjectId = require('mongodb').ObjectID;

exports.install = function () {
    F.route('/project/new', createProject);
    F.route('/project/{id}', getProject);
    F.route('/payment/new/{projectId}', showPayment);
    F.route('/payment/success', showPaymentSuccess);
};

function getProject(id) {
    id = id || '';

    var self = this;
    var projects = DATABASE('projects');

    projects.find({_id: new ObjectId(id)}).limit(1).next(function (err, doc) {
        if (!err) {
            doc.id = doc._id;
            self.view('/project/detail', doc);
        }
    });
}

function createProject() {
    var self = this;
    self.view('/project/new-project', {});
}

function showPayment(projectId) {

    projectId = projectId || '';

    var self = this;
    var projects = DATABASE('projects');

    projects.find({_id: new ObjectId(projectId)}).limit(1).next(function (err, doc) {
        if (!err) {
            doc.id = doc._id;
            self.view('/payment/new', {
                project: doc,
                packageId: self.query.package_id
            });
        }
    });
}

function showPaymentSuccess() {
    var self = this;
    self.view('/payment/success', self.query);
}
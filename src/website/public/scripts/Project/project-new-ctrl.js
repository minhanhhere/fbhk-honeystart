'use strict';

(function () {

    angular.module('app.project').controller('ProjectNewCtrl', ProjectNewCtrl);

    ProjectNewCtrl.$inject = ['$rootScope', '$window', '$api', '$location', '$timeout', '$uibModal'];

    function ProjectNewCtrl($rootScope, $window, $api, $location, $timeout, $uibModal) {

        var vm = this;
        var ApiProject = $api.project();

        vm.prices = [
            {
                name: 'Value',
                cost: '1,000,000',
                features: [
                    'Couples Page',
                    '-',
                    '-',
                    '-',
                ]
            },
            {
                name: 'Silver',
                cost: '2,000,000',
                features: [
                    'Couples Page',
                    'Thank-you gift',
                    '-',
                    '-',
                ]
            },
            {
                name: 'Gold',
                cost: '3,000,000',
                features: [
                    'Couples Page',
                    'Thank-you gift',
                    'Copywriter',
                    '-',
                ]
            },
            {
                name: 'Plantinum',
                cost: '4,000,000',
                features: [
                    'Couples Page',
                    'Thank-you gift',
                    'Copywriter',
                    'Wedding Planner',
                ]
            },
        ];

        vm.loadProjectAsync = function (id) {
            ApiProject.get({id: id}, function (res) {
                vm.project = res;
                vm.loadFacebookPlugin();
            });
        };

        vm.submit = function () {
            ApiProject.createProject({}, function (res) {
                console.log(res);
                window.location.href = '/project/' + res.id;
            }, function (res) {
                //error
            })
        };

        vm.getFullName = function (arr) {
            return arr.join(' ');
        };

        vm.getFullNameLimit = function (first, last, limit) {
            var fSplit = first.split(' ');
            var lSplit = last.split(' ');
            var split = fSplit.concat(lSplit);
            return split.slice(0, limit).join(' ');
        };
    }

})();

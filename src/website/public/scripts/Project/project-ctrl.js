'use strict';

(function () {

    angular.module('app.project', ['app.rest', 'app.modal']).controller('ProjectCtrl', ProjectCtrl);

    ProjectCtrl.$inject = ['$rootScope', '$window', '$api', '$location', '$timeout', '$uibModal', '$injectedModel'];

    function ProjectCtrl($rootScope, $window, $api, $location, $timeout, $uibModal, $injectedModel) {

        var vm = this;
        var ApiProject = $api.project();
        var tab = {
            STORY: 0,
            UPDATES: 1,
            COMMENTS: 2,
            BACKERS: 3
        };

        vm.showless = true;

        vm.template = {
            left: '/scripts/Project/project-left.html?v=' + $window.app_version,
            packages: '/scripts/Project/project-packages.html?v=' + $window.app_version,
            updates: '/scripts/Project/project-updates.html?v=' + $window.app_version,
            backers: '/scripts/Project/project-backers.html?v=' + $window.app_version
        };

        vm.project = $injectedModel;

        console.log(vm.project);

        vm.activeTab = 0;

        vm.setActiveTab = function (index) {
            vm.activeTab = index % 4;

            switch (index) {
                case tab.UPDATES:
                    ApiProject.getUpdates({id: vm.project.id}, function (res) {
                        vm.project.updates = res;
                    });
                    break;
                case tab.BACKERS:
                    vm.getBackers();
                    break;
            }
        };

        vm.getBackers = function () {
            ApiProject.getBackers({id: vm.project.id}, function (res) {
                vm.project.backers = res;
            });
        };

        vm.formatEstimation = function (ts) {
            return moment.unix(ts / 1000).format('MMM, YYYY');
        };

        vm.fromNow = function (ts) {
            return moment(ts).fromNow();
        };

        vm.loadProjectAsync = function (id) {
            ApiProject.get({id: id}, function (res) {
                vm.project = res;
                vm.loadFacebookPlugin();
            });
        };

        vm.toggleShowless = function () {
            vm.showless = !vm.showless;
        };

        vm.loadFacebookPlugin = function () {
            $("meta[property='og\\:url']").attr("content", $location.absUrl());
            $("meta[property='og\\:title']").attr("content", vm.project.name + ' - HoneyStart');
            $("meta[property='og\\:image']").attr("content", vm.project.avatar);
            $timeout(function () {
                var wrapper = $('.fb-comments-wrapper');
                wrapper.empty();
                var fbdiv = $('<div>');
                fbdiv.addClass('fb-comments');
                fbdiv.attr('data-href', $location.absUrl());
                fbdiv.attr('data-numposts', 10);
                fbdiv.attr('data-width', '100%');
                fbdiv.attr('data-order-by', 'reverse_time');
                wrapper.append(fbdiv);
                wrapper.append('<script>FB.XFBML.parse();</script>');
            }, 200);
        };

        vm.showPackges = function() {
            var modal = $uibModal.open({
                animation: true,
                size: 'md',
                templateUrl: '/scripts/Project/picker-modal.html?v=' + window.app_version,
                controller: 'PackagePickerModalCtrl',
                controllerAs: 'vm',
                resolve: {
                    projectId: function () {
                        return vm.project.id;
                    },
                    packages: function () {
                        return vm.project.packages;
                    },
                }
            });
        };

        vm.getBackers();

        vm.getFullName = function (arr) {
            return arr.join(' ');
        };

        vm.getFullNameLimit = function (first, last, limit) {
            var fSplit = first.split(' ');
            var lSplit = last.split(' ');
            var split = fSplit.concat(lSplit);
            return split.slice(0, limit).join(' ');
        }
    }

})();

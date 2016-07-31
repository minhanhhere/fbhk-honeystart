'use strict';

(function () {

    angular.module('app.rest', ['ngResource']).service('$api', Api);

    Api.$inject = ['$rootScope', '$resource', '$http'];

    function Api($rootScope, $resource, $http) {

        var rootUrl = 'http://localhost:5000';
        var apiUrl = rootUrl + '/api/v1';

        this.welcome = function () {
            return $resource(apiUrl);
        };

        this.user = function () {
            return $resource(apiUrl + '/users/:id', {id: '@id'}, {
                login: {method: 'POST', url: apiUrl + '/users/login'},
                loginFacebook: {method: 'POST', url: apiUrl + '/users/loginFacebook'},
                register: {method: 'POST', url: apiUrl + '/users'},
            });
        };

        this.project = function () {
            return $resource(apiUrl + '/projects/:id', {id: '@id'}, {
                createProject: {method: 'POST', url: apiUrl + '/projects'},
                getHome: {method: 'GET', url: apiUrl + '/projects/home', isArray: true},
                getPackageById: {method: 'GET', url: apiUrl + '/projects/:id/packages/:pid', params: {pid: '@pid'}},
                getUpdates: {method: 'GET', url: apiUrl + '/projects/:id/updates', isArray: true},
                getBackers: {method: 'GET', url: apiUrl + '/projects/:id/backers', isArray: true}
            });
        };

        this.transaction = function () {
            return $resource(apiUrl + '/transactions/:id', {id: '@id'}, {
                prepare: {method: 'POST', url: apiUrl + '/transactions/prepare'},
                verify: {method: 'POST', url: apiUrl + '/transactions/verify', timeout: 25000},
            });
        };

        this.disqus = function () {
            return $http({method: 'GET', url: apiUrl + '/disqus'});
        };
    }

})();

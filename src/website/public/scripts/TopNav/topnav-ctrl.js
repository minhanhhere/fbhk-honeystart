'use strict';

(function () {

    angular.module("HoneyStart").controller('TopNavCtrl', TopNavCtrl);

    TopNavCtrl.$inject = ['$scope', '$auth', '$api'];

    function TopNavCtrl($scope, $auth, $api) {

        var vm = this;
        var ApiUser = $api.user();

        vm.auth = $auth;

        vm.changeTheme = function (setTheme) {

            $('<link>')
                .appendTo('head')
                .attr({type: 'text/css', rel: 'stylesheet'})
                .attr('href', 'styles/app-' + setTheme + '.css?v=' + window.app_version);
        };

        vm.showAuthModal = function (forRegister) {
            $auth.showModal(forRegister, function (user) {
            });
        };

        vm.autoLoginFB = function () {
            FB.api('/me', {fields: 'last_name,first_name,email'}, function (me) {
                vm.email = me.email;
                vm.firstName = me.first_name;
                vm.lastName = me.last_name;
                vm.facebookId = me.id;
                ApiUser.loginFacebook({
                    email: vm.email,
                    firstName: vm.firstName,
                    lastName: vm.lastName,
                    facebookId: vm.facebookId
                }, function (res) {
                    vm.loading = false;
                    $auth.setUser(res);
                }, function (res) {
                    vm.error = "Login failed: " + res.data;
                    vm.loading = false;
                });
            });
        };
    }

})();

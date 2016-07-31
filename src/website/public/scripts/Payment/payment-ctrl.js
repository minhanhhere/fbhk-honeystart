'use strict';

(function () {

    angular.module('app.payment', ['app.rest', 'app.auth', 'timer']).controller('PaymentCtrl', PaymentCtrl);

    PaymentCtrl.$inject = ['$auth', '$window', '$location', '$api', '$injectedModel'];

    function PaymentCtrl($auth, $window, $location, $api, $injectedModel) {

        var vm = this;
        var ApiProject = $api.project();
        var ApiTransaction = $api.transaction();

        vm.auth = $auth;

        vm.template = {
            invoice: '/scripts/Payment/payment-invoice.html?v=' + $window.app_version,
            info: '/scripts/Payment/payment-info.html?v=' + $window.app_version,
        };

        vm.data = {
            owner: {},
            address: {},
            contribution: {
                type: 'full',
                customName: '',
                amount: 500000
            },
            project: $injectedModel.project,
            package: _.find($injectedModel.project.packages, function (it) {
                return it.id == $injectedModel.packageId;
            })
        };

        console.log($injectedModel);

        vm.amounts = [];

        function init() {
            for (var i = 2; i <= 10; i++) {
                vm.amounts.push({
                    value: i * 100000
                });
            }
            vm.amounts.push({
                value: 1500000
            });
            vm.amounts.push({
                value: 2000000
            });
            vm.amounts.push({
                key: 'OTHER',
                value: 0
            });
            var user = $auth.getUser();
            if (user) {
                vm.data.owner = {
                    firstName: user.firstName,
                    lastName: user.lastName,
                    email: user.email,
                    phone: user.phone,
                };
                vm.data.address = {
                    street: user.address ? user.address.street : '',
                    block: user.address ? user.address.block : '',
                    district: user.address ? user.address.district : '',
                    city: user.address ? user.address.city : '',
                    country: user.address ? user.address.country : '',
                };
            }
        }

        init();

        vm.showCustom = false;

        vm.setAmount = function (key) {
            if (key == 'OTHER') {
                vm.showCustom = true;
                vm.data.contribution.amount = 0;
            } else {
                vm.showCustom = false;
            }
        };

        vm.setCustomAmount = function () {
            vm.showCustom = false;
        };

        vm.onOtherContributionTypeClick = function () {
            vm.data.contribution.type = 'other';
            $('#txtContributionCustomName').focus();
        };

        vm.showAuthModal = function (forRegister) {
            $auth.showModal(forRegister, function (user) {
                vm.data.owner.email = user.email;
                vm.data.owner.firstName = user.firstName;
                vm.data.owner.lastName = user.lastName;
                vm.data.owner.phone = user.phone;
                vm.data.address = user.address;
            });
        };

        vm.getGroomName = function () {
            if (vm.data.project.groom) {
                return vm.getName([vm.data.project.groom.lastName, vm.data.project.groom.firstName]);
            }
            return '';
        };

        vm.getBrideName = function () {
            if (vm.data.project.bride) {
                return vm.getName([vm.data.project.bride.lastName, vm.data.project.bride.firstName]);
            }
            return '';
        };

        vm.getName = function (arr) {
            return arr.join(' ');
        };

        vm.prepareTransaction = function () {
            //todo validate input
            //todo show loading
            var postData = {
                owner: vm.data.owner,
                address: vm.data.address,
                projectId: vm.data.project.id,
                packageId: vm.data.package.id,
                totalAmount: vm.data.contribution.amount,
                type: vm.data.contribution.type,
                cancelUrl: $location.absUrl()
            };
            ApiTransaction.prepare({}, postData, function (res) {
                //todo hide loading
                window.location = res.redirect;
            }, function (res) {
                //todo hide loading
            });
        }
    }

})();

'use strict';

(function () {

    angular.module("app.project").controller('PackagePickerModalCtrl', PackagePickerModalCtrl);

    PackagePickerModalCtrl.$inject = ['$uibModalInstance', 'projectId', 'packages'];

    function PackagePickerModalCtrl($uibModalInstance, projectId, packages) {

        var vm = this;

        vm.projectId = projectId;
        vm.packages = packages;

        vm.check1 = true;
        vm.check2 = true;
        vm.check3 = true;

        vm.dismiss = function () {
            $uibModalInstance.dismiss('cancel');
        };

        vm.continue = function() {
			window.location.href = '/payment/new/' + vm.projectId + '?package_id=' + vm.packages[0].id;
        };
    }

})();
